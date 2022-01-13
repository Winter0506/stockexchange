package logic

import (
	"context"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockListLogic {
	return &GetStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockListLogic) GetStockList(in *stock.PageInfo) (*stock.StockListResponse, error) {
	// 管理员查询所有
	all, err := l.svcCtx.Model.FindAll()

	if err != nil {
		return nil, err
	}

	rsp := &stock.StockListResponse{
		Total: 0,
		Data:  nil,
	}

	rsp.Total = int32(len(*all))
	allValue := *all
	page, pageSize := in.Pn, in.PSize
	tmpAll := allValue[(page-1)*pageSize : (page-1)*pageSize+pageSize]

	for _, eveRet := range tmpAll {
		stockInfoRsp := &stock.StockInfoResponse{
			Id:        eveRet.Id,
			StockName: eveRet.Stockname,
			StockCode: eveRet.Stockcode,
		}
		rsp.Data = append(rsp.Data, stockInfoRsp)
	}

	return rsp, nil
}
