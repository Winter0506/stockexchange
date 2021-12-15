package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"stockexchange/rpc/demo/global"
	"time"
)

func InitDB() {
	// 配置和日志后面再加
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "127.0.0.1", 3306, "stockexchange")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
