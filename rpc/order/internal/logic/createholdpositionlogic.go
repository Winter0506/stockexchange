package logic

import (
	"context"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateHoldPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHoldPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHoldPositionLogic {
	return &CreateHoldPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateHoldPositionLogic) CreateHoldPosition(in *order.CreateHoldPositionRequest) (*order.HoldPositionResponse, error) {
	// todo: add your logic here and delete this line
	return &order.HoldPositionResponse{}, nil
}
