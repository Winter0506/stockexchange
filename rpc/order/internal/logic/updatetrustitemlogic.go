package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateTrustItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTrustItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTrustItemLogic {
	return &UpdateTrustItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTrustItemLogic) UpdateTrustItem(in *order.UpdateTrustRequest) (*order.TrustInfoResponse, error) {
	// todo: add your logic here and delete this line
	// 委托只要已提交也就已经确定 不允许修改
	// 程序这里做的是即时成交 也不允许撤单
	return &order.TrustInfoResponse{}, nil
}
