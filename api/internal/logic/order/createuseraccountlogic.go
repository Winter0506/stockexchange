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

type CreateUserAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateUserAccountLogic {
	return CreateUserAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserAccountLogic) CreateUserAccount(req types.ReqUserAccountCreate) (*types.RespUserAccountDetail, error) {
	// 简单实现 创建用户账户
	rsp, err := l.svcCtx.Order.CreateUserAccount(l.ctx, &order.CreateUserAccountInfo{
		UserId:        int64(req.UserId),
		Account:       req.Account,
		MarketValue:   0,
		Available:     0,
		ProfitAndLoss: 0,
	})
	if err != nil {
		logx.Errorf("创建用户账户失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "创建用户账户失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespUserAccountDetail{
			DetailMeta: detailStatus,
		}, errors.New("创建用户账户失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "创建用户账户成功",
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
