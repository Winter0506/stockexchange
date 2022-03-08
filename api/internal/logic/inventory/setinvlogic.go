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

type SetInvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetInvLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetInvLogic {
	return SetInvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetInvLogic) SetInv(req types.ReqSetInv) (*types.RespSetInv, error) {
	// 简单实现 设置库存
	_, err := l.svcCtx.Inventory.SetInv(l.ctx, &inventory.StockInfo{
		StockId: req.StockId,
		Num:     req.Num,
	})
	if err != nil {
		logx.Errorf("设置股票库存失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "设置股票库存失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespSetInv{
			DetailMeta: detailStatus,
		}, errors.New("设置股票库存失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "设置股票库存成功",
		Status: http.StatusOK,
	}
	return &types.RespSetInv{DetailMeta: detailStatus}, nil
}
