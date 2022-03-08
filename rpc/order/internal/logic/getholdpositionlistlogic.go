package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetHoldPositionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHoldPositionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHoldPositionListLogic {
	return &GetHoldPositionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHoldPositionListLogic) GetHoldPositionList(in *order.UserHoldPositionRequest) (*order.HoldPositionListResponse, error) {
	// TODO 要验证一下这个语句是否正确
	all, err := l.svcCtx.HoldPositionModel.FindHoldPositionByUser(in.User)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询用户全部持仓出错")
	}

	rsp := &order.HoldPositionListResponse{
		Total: 0,
		Data:  nil,
	}

	rsp.Total = int32(len(*all))
	allValue := *all

	for _, eveRet := range allValue {
		holdPositionResponse := &order.HoldPositionResponse{
			Id:        eveRet.Id,
			User:      eveRet.User,
			Stock:     eveRet.Stock,
			StockName: eveRet.StockName,
			Number:    int32(eveRet.Number),
			Cost:      float32(eveRet.Cost),
		}
		rsp.Data = append(rsp.Data, holdPositionResponse)
	}
	return rsp, nil
}
