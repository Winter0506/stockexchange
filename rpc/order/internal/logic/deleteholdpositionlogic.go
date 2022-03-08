package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteHoldPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHoldPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHoldPositionLogic {
	return &DeleteHoldPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  卖出的时候 正好一点不留 就用delete
func (l *DeleteHoldPositionLogic) DeleteHoldPosition(in *order.DeleteHoldPositionRequest) (*order.DeleteHoldPositionResponse, error) {
	// todo: add your logic here and delete this line

	return &order.DeleteHoldPositionResponse{}, nil
}
