package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type TrustItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTrustItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrustItemListLogic {
	return &TrustItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TrustItemListLogic) TrustItemList(in *order.UserTrustInfoRequest) (*order.TrustListResponse, error) {
	all, err := l.svcCtx.TrustModel.FindTrustByUser(in.User)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询用户委托列表出错")
	}

	rsp := &order.TrustListResponse{
		Total: 0,
		Data:  nil,
	}

	rsp.Total = int32(len(*all))
	allValue := *all
	page, pageSize := in.Page, in.PageSize
	tmpAll := allValue[(page-1)*pageSize : (page-1)*pageSize+pageSize]

	for _, eveRet := range tmpAll {
		userAccountRsp := &order.TrustInfoResponse{
			Id:         eveRet.Id,
			User:       eveRet.User,
			Stock:      eveRet.Stock,
			Number:     int32(eveRet.Number),
			Cost:       float32(eveRet.Cost),
			Direction:  uint32(eveRet.Direction),
			DealNumber: int32(eveRet.Dealnumber),
			DealCost:   float32(eveRet.Dealcost),
			Status:     eveRet.Status,
			TrustSn:    eveRet.TrustSn,
		}
		rsp.Data = append(rsp.Data, userAccountRsp)
	}
	return rsp, nil
}
