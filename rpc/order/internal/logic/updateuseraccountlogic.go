package logic

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/order/model"
	"time"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAccountLogic {
	return &UpdateUserAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserAccountLogic) UpdateUserAccount(in *order.UpdateUserAccountInfo) (*order.UserAccountResponse, error) {
	// 先查找数据库 如果账户数据库没有用户的信息就不能更新
	hasUserAccountInfo, err := l.svcCtx.UserAccountModel.FindOne(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "用户账户信息不存在")
	}
	err = l.svcCtx.UserAccountModel.Update(&model.Useraccount{
		Userid:        in.UserId,
		Account:       float64(in.Account),
		MarketValue:   hasUserAccountInfo.MarketValue,
		Available:     float64(in.Account) - hasUserAccountInfo.MarketValue.Float64,
		ProfitAndLoss: hasUserAccountInfo.ProfitAndLoss,
		CreatedAt:     hasUserAccountInfo.CreatedAt,
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
		IsDeleted:     0,
	})
	if err != nil {
		return nil, errors.New("更新用户账户错误")
	}

	hasUserAccountInfo, err = l.svcCtx.UserAccountModel.FindOne(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "用户账户信息查询错误,请重试")
	}
	marketValue := hasUserAccountInfo.MarketValue.Float64
	return &order.UserAccountResponse{
		UserId:        hasUserAccountInfo.Userid,
		Account:       float32(hasUserAccountInfo.Account),
		MarketValue:   float32(marketValue),
		Available:     float32(hasUserAccountInfo.Available),
		ProfitAndLoss: float32(hasUserAccountInfo.ProfitAndLoss),
	}, nil
}
