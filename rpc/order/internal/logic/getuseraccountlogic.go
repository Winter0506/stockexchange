package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"
	"stockexchange/rpc/stock/stock"
)

type GetUserAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAccountLogic {
	return &GetUserAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  账户相关
func (l *GetUserAccountLogic) GetUserAccount(in *order.IdRequest) (*order.UserAccountResponse, error) {
	hasUserAccountInfo, err := l.svcCtx.UserAccountModel.FindOne(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "用户账户信息不存在")
	}
	// 1.1 查询用户持仓
	holdPositionList, err := l.svcCtx.HoldPositionModel.FindHoldPositionByUser(in.Id)
	// 1.2 查询用户所有持仓
	var oldMarketValue float64
	var nowMarketValue float64
	if holdPositionList != nil {
		for _, holdPositionItem := range *holdPositionList {
			oldMarketValue += holdPositionItem.Cost * float64(holdPositionItem.Number)
			rspStock, _ := l.svcCtx.Stock.GetStockById(l.ctx, &stock.IdRequest{
				Id: holdPositionItem.Stock,
			})
			nowMarketValue += float64(rspStock.BaseInfo.PresentPrice) * float64(holdPositionItem.Number)
		}
	}
	profitAndLoss := nowMarketValue - oldMarketValue
	// 查询当前股价
	return &order.UserAccountResponse{
		UserId:        hasUserAccountInfo.Userid,
		Account:       float32(hasUserAccountInfo.Account),
		MarketValue:   float32(nowMarketValue),
		Available:     float32(hasUserAccountInfo.Available),
		ProfitAndLoss: float32(profitAndLoss),
	}, nil
}
