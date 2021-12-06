package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"stockexchange/rpc/model"
	"stockexchange/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
