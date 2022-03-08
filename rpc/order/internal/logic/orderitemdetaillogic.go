package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemDetailLogic {
	return &OrderItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemDetailLogic) OrderItemDetail(in *order.OrderInfoRequest) (*order.OrderInfoResponse, error) {
	// todo: add your logic here and delete this line
	// 订单这里只做model层面的保存 先不写获取实现  订单同时也不应该显示给用户
	// 更新 创建 删除方面都是根据用户的委托来的
	return &order.OrderInfoResponse{}, nil
}
