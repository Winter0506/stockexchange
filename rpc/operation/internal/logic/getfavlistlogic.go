package logic

import (
	"context"
	"stockexchange/rpc/operation/model"

	"stockexchange/rpc/operation/internal/svc"
	"stockexchange/rpc/operation/operation"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetFavListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavListLogic {
	return &GetFavListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavListLogic) GetFavList(in *operation.UserFavRequest) (*operation.UserFavListResponse, error) {
	var rsp *operation.UserFavListResponse
	var userFavs []model.UserFav
	var userFavList []*operation.UserFavResponse

	rsp = &operation.UserFavListResponse{}
	// 查询用户的收藏记录  有用户id 无股票id
	// 查询某件商品被哪些用户收藏了 无用户id 有股票id
	// 只要一个跑通另外一个必然能够跑通
	result := l.svcCtx.DbEngine.Where(&model.UserFav{User: in.UserId, Stock: in.StockId}).Find(&userFavs)
	rsp.Total = int32(result.RowsAffected)

	for _, userFav := range userFavs {
		userFavList = append(userFavList, &operation.UserFavResponse{
			UserId:  userFav.User,
			StockId: userFav.Stock,
		})
	}

	rsp.Data = userFavList

	return rsp, nil
}
