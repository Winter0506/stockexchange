package order

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"stockexchange/rpc/order/order"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HoldPositionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHoldPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HoldPositionListLogic {
	return HoldPositionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HoldPositionListLogic) HoldPositionList(req types.ReqUserAccountId) (*types.RespHoldPositionList, error) {
	// 简单实现 查询用户持仓列表
	resp, err := l.svcCtx.Order.GetHoldPositionList(l.ctx, &order.UserHoldPositionRequest{
		User: int64(req.Id),
	})
	if err != nil {
		logx.Errorf("用户持仓列表查询失败: ", err.Error())
		detailStatus := types.HoldPositionListMeta{
			Msg:    "用户持仓列表查询失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespHoldPositionList{
			HoldPositionListMeta: detailStatus,
		}, errors.New("用户持仓列表查询失败")
	}
	detailStatus := types.HoldPositionListMeta{
		Msg:    "用户持仓列表查询成功",
		Status: http.StatusOK,
	}

	holdPositionList := make([]string, 0)
	for _, holdPosition := range resp.Data {
		tempHoldPosition, _ := json.Marshal(holdPosition)
		holdPositionList = append(holdPositionList, string(tempHoldPosition))
	}
	return &types.RespHoldPositionList{
		HoldPositionList:     holdPositionList,
		HoldPositionListMeta: detailStatus,
	}, nil
}
