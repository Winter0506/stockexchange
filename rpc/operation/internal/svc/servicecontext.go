package svc

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"stockexchange/rpc/operation/internal/config"
	"stockexchange/rpc/operation/model"
	"time"
)

type ServiceContext struct {
	Config   config.Config
	DbEngine *gorm.DB
	Cache    cache.CacheConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启动Gorm支持
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "operation_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: newLogger,
	})
	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	//自动同步更新表结构,不需要建表了
	db.AutoMigrate(&model.UserFav{})

	return &ServiceContext{Config: c, DbEngine: db, Cache: c.Cache}
}
