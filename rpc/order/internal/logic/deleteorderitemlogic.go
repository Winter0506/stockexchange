package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderItemLogic {
	return &DeleteOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderItemLogic) DeleteOrderItem(in *order.DeleteOrderRequest) (*order.DeleteOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &order.DeleteOrderResponse{}, nil
}
