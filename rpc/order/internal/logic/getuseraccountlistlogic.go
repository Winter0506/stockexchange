package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserAccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAccountListLogic {
	return &GetUserAccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAccountListLogic) GetUserAccountList(in *order.PageInfo) (*order.UserAccountListResponse, error) {
	all, err := l.svcCtx.UserAccountModel.FindAll()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询全部用户账户出错")
	}

	rsp := &order.UserAccountListResponse{
		Total: 0,
		Data:  nil,
	}

	rsp.Total = int32(len(*all))
	allValue := *all
	page, pageSize := in.Pn, in.PSize
	tmpAll := allValue[(page-1)*pageSize : (page-1)*pageSize+pageSize]

	for _, eveRet := range tmpAll {
		userAccountRsp := &order.UserAccountResponse{
			UserId:        eveRet.Userid,
			Account:       float32(eveRet.Account),
			MarketValue:   float32(eveRet.MarketValue.Float64),
			Available:     float32(eveRet.Available),
			ProfitAndLoss: float32(eveRet.ProfitAndLoss),
		}
		rsp.Data = append(rsp.Data, userAccountRsp)
	}
	return rsp, nil
}
