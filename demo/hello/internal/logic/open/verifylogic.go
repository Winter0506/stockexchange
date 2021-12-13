package logic

import (
	"context"

	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyLogic struct {
	ctx context.Context
	logx.Logger
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyLogic {
	return VerifyLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyLogic) Verify(req types.VerifyReq) (*types.VerifyResp, error) {
	return &types.VerifyResp{}, nil
}
