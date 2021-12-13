package svc

import (
	"hello/internal/config"
	"hello/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tech_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})

	return &ServiceContext{Config: c, DbEngin: db}
}
