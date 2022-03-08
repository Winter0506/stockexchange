package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemListLogic {
	return &OrderItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemListLogic) OrderItemList(in *order.UserOrderInfoRequest) (*order.OrderListResponse, error) {
	// todo: add your logic here and delete this line

	return &order.OrderListResponse{}, nil
}
