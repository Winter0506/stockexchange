package order

import (
	"context"
	"errors"
	"net/http"
	"stockexchange/rpc/order/order"
	"stockexchange/rpc/stock/stock"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateTrustLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTrustLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateTrustLogic {
	return CreateTrustLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTrustLogic) CreateTrust(req types.ReqCreateTrustItem) (*types.RespTrustInfoResponse, error) {
	// 简单实现 创建委托
	rspStock, err := l.svcCtx.Stock.GetStockById(l.ctx, &stock.IdRequest{
		Id: int64(req.StockId),
	})
	// 这里应该有判断当前股价 和 委托价格的逻辑  为了简单实现 省略了
	rsp, err := l.svcCtx.Order.CreateTrustItem(l.ctx, &order.TrustItemRequest{
		User:      int64(req.UserId),
		Stock:     int64(req.StockId),
		Number:    req.Num,
		Cost:      rspStock.BaseInfo.PresentPrice,
		Direction: uint32(req.Direction),
	})
	if err != nil {
		logx.Errorf("创建用户委托失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "创建用户委托失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespTrustInfoResponse{
			DetailMeta: detailStatus,
		}, errors.New("创建用户委托失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "创建用户委托成功",
		Status: http.StatusOK,
	}
	return &types.RespTrustInfoResponse{
		TrustInfoResponse: types.TrustInfoResponse{
			Id:         int32(rsp.Id),
			User:       int32(rsp.User),
			Stock:      int32(rsp.Stock),
			Num:        rsp.Number,
			Cost:       rsp.Cost,
			Direction:  uint(rsp.Direction),
			DealNumber: rsp.DealNumber,
			DealCost:   rsp.DealCost,
			Status:     rsp.Status,
			TrustSn:    rsp.TrustSn,
		},
		DetailMeta: detailStatus,
	}, nil
}
