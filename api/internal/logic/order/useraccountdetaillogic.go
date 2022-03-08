package order

import (
	"context"
	"errors"
	"net/http"
	"stockexchange/rpc/order/order"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAccountDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAccountDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserAccountDetailLogic {
	return UserAccountDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAccountDetailLogic) UserAccountDetail(req types.ReqUserAccountId) (*types.RespUserAccountDetail, error) {
	// 简单实现 查询用户账户
	rsp, err := l.svcCtx.Order.GetUserAccount(l.ctx, &order.IdRequest{
		Id: int64(req.Id),
	})
	if err != nil {
		logx.Errorf("查询用户账户失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "查询用户账户失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespUserAccountDetail{
			DetailMeta: detailStatus,
		}, errors.New("查询用户账户失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "查询用户账户成功",
		Status: http.StatusOK,
	}
	return &types.RespUserAccountDetail{
		AccountDetailMessage: types.AccountDetailMessage{
			UserId:         int32(rsp.UserId),
			Account:        rsp.Account,
			MarketValue:    rsp.MarketValue,
			Available:      rsp.Available,
			ProfiltAndLoss: rsp.ProfitAndLoss,
		},
		DetailMeta: detailStatus,
	}, nil
}
