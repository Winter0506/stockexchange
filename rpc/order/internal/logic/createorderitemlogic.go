package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderItemLogic {
	return &CreateOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  订单相关
func (l *CreateOrderItemLogic) CreateOrderItem(in *order.OrderItemRequest) (*order.OrderInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &order.OrderInfoResponse{}, nil
}
