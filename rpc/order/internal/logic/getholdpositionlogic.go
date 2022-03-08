package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetHoldPositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHoldPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHoldPositionLogic {
	return &GetHoldPositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  持仓相关
func (l *GetHoldPositionLogic) GetHoldPosition(in *order.HoldPositionRequest) (*order.HoldPositionResponse, error) {
	rsp, err := l.svcCtx.HoldPositionModel.FindOne(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询全部用户该笔持仓出错")
	}
	return &order.HoldPositionResponse{
		Id:        rsp.Id,
		User:      rsp.User,
		Stock:     rsp.Stock,
		StockName: rsp.StockName,
		Number:    int32(rsp.Number),
		Cost:      float32(rsp.Cost),
	}, nil
}
