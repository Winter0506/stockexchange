package inventory

import (
	"context"
	"errors"
	"net/http"
	"stockexchange/rpc/inventory/inventory"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type InvDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) InvDetailLogic {
	return InvDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InvDetailLogic) InvDetail(req types.ReqInvDetail) (*types.RespInvDetail, error) {
	// 简单实现 查询库存
	rsp, err := l.svcCtx.Inventory.InvDetail(l.ctx, &inventory.StockInfo{
		StockId: req.StockId,
	})
	if err != nil {
		logx.Errorf("查询股票库存失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "查询股票库存失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespInvDetail{
			DetailMeta: detailStatus,
		}, errors.New("查询股票库存失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "查询股票库存成功",
		Status: http.StatusOK,
	}
	return &types.RespInvDetail{
		ReqSetInv: types.ReqSetInv{
			StockId: rsp.StockId,
			Num:     rsp.Num,
		},
		DetailMeta: detailStatus,
	}, nil
}
