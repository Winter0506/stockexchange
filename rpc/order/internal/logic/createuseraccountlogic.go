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

type CreateUserAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserAccountLogic {
	return &CreateUserAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserAccountLogic) CreateUserAccount(in *order.CreateUserAccountInfo) (*order.UserAccountResponse, error) {
	// 先查找数据库 如果账户数据库里面没有用户的信息 再去新建用户账户信息
	hasUserAccountInfo, err := l.svcCtx.UserAccountModel.FindOne(in.UserId)
	if hasUserAccountInfo != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户账户信息已存在")
	}
	rsp, err := l.svcCtx.UserAccountModel.Insert(&model.Useraccount{
		Userid:  in.UserId,
		Account: float64(in.Account),
		MarketValue: sql.NullFloat64{
			Float64: 0,
			Valid:   true,
		},
		Available:     float64(in.Account),
		ProfitAndLoss: 0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
		IsDeleted:     0,
	})
	if err != nil {
		return nil, errors.New("创建用户账户错误")
	}

	insertId, _ := rsp.LastInsertId()
	hasUserAccountInfo, err = l.svcCtx.UserAccountModel.FindOne(insertId)
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
