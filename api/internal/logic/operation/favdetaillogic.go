package operation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"stockexchange/rpc/operation/operation"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FavDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) FavDetailLogic {
	return FavDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavDetailLogic) FavDetail(req types.ReqUserFav) (*types.RespFavDetail, error) {
	// 首先应该判断用户id 和 股票id是否存在 这里省略
	// 使用token信息判断是否是用户本人
	userId := l.ctx.Value("userId")
	userIdInt, err := json.Number(fmt.Sprintf("%v", userId)).Int64()
	if err != nil {
		detailStatus := types.DetailMeta{
			Msg:    "鉴权失败,请稍后重试",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespFavDetail{
			DetailMeta: detailStatus,
		}, errors.New("鉴权失败,请稍后重试")
	}
	if int32(userIdInt) != req.UserId {
		detailStatus := types.DetailMeta{
			Msg:    "无权使用他人信息,请重试",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespFavDetail{
			DetailMeta: detailStatus,
		}, errors.New("无权使用他人信息,请重试")
	}
	_, err = l.svcCtx.Operation.GetUserFavDetail(l.ctx, &operation.UserFavRequest{
		UserId:  req.UserId,
		StockId: req.StockId,
	})
	if err != nil {
		e, _ := status.FromError(err)
		if e.Code() == codes.NotFound {
			detailStatus := types.DetailMeta{
				Msg:    "收藏股票记录不存在",
				Status: http.StatusOK,
			}
			return &types.RespFavDetail{
				DetailMeta: detailStatus,
			}, nil
		}
		logx.Errorf("收藏股票详情查询失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "收藏股票详情查询失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespFavDetail{
			DetailMeta: detailStatus,
		}, errors.New("收藏股票详情查询失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "收藏股票记录存在",
		Status: http.StatusOK,
	}
	return &types.RespFavDetail{DetailMeta: detailStatus}, nil
}
