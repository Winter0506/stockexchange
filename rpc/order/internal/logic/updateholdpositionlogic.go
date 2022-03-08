package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateHoldPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHoldPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHoldPositionLogic {
	return &UpdateHoldPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHoldPositionLogic) UpdateHoldPosition(in *order.UpdateHoldPositionRequest) (*order.HoldPositionResponse, error) {
	// todo: add your logic here and delete this line
	return &order.HoldPositionResponse{}, nil
}
