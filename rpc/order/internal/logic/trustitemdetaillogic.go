package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type TrustItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTrustItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrustItemDetailLogic {
	return &TrustItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TrustItemDetailLogic) TrustItemDetail(in *order.TrustInfoRequest) (*order.TrustInfoResponse, error) {
	hasTrustItemDetail, err := l.svcCtx.TrustModel.FindOne(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该条持仓详情信息不存在")
	}
	if hasTrustItemDetail.TrustSn != in.TrustSn {
		return nil, status.Errorf(codes.NotFound, "该条持仓详情与委托号不匹配")
	}
	return &order.TrustInfoResponse{
		Id:         hasTrustItemDetail.Id,
		User:       hasTrustItemDetail.User,
		Stock:      hasTrustItemDetail.Stock,
		Number:     int32(hasTrustItemDetail.Number),
		Cost:       float32(hasTrustItemDetail.Cost),
		Direction:  uint32(hasTrustItemDetail.Direction),
		DealNumber: int32(hasTrustItemDetail.Dealnumber),
		DealCost:   float32(hasTrustItemDetail.Dealcost),
		Status:     hasTrustItemDetail.Status,
		TrustSn:    hasTrustItemDetail.TrustSn,
	}, nil
}
