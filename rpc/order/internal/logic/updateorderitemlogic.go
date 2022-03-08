package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderItemLogic {
	return &UpdateOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderItemLogic) UpdateOrderItem(in *order.UpdateOrderRequest) (*order.OrderInfoResponse, error) {
	// todo: add your logic here and delete this line
	// 订单目前不允许修改
	return &order.OrderInfoResponse{}, nil
}
