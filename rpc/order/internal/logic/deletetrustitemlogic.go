package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteTrustItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTrustItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTrustItemLogic {
	return &DeleteTrustItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTrustItemLogic) DeleteTrustItem(in *order.DeleteTrustRequest) (*order.DeleteTrustResponse, error) {
	// todo: add your logic here and delete this line
	// 委托这里是即时成交 不允许撤单
	return &order.DeleteTrustResponse{}, nil
}
