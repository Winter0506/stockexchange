package logic

import (
	"context"
	"database/sql"
	"stockexchange/rpc/stock/model"
	"time"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockLogic) UpdateStock(in *stock.UpdateStockInfo) (*stock.Empty, error) {
	// 这个接口也不是一般人能够调用的
	rsp, _ := l.svcCtx.Model.FindOne(in.Id)
	createdTime := rsp.CreatedAt
	updatedTime := rsp.UpdatedAt
	// 因为传进来的都为空
	if in.IsDeleted != 0 {
		err := l.svcCtx.Model.Update(&model.Stock{
			Id:        int32(in.Id),
			Stockname: in.StockName,
			Stockcode: in.StockName,
			CreatedAt: createdTime,
			UpdatedAt: updatedTime,
			DeletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			IsDeleted: 1,
		})
		if err != nil {
			return nil, err
		}
	} else {
		err := l.svcCtx.Model.Update(&model.Stock{
			Id:        int32(in.Id),
			Stockname: in.StockName,
			Stockcode: in.StockName,
			CreatedAt: createdTime,
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return nil, err
		}
	}

	return &stock.Empty{}, nil
}
