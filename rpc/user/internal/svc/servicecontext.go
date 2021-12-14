package svc

import (
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	//"gorm.io/gorm/schema"
	//"github.com/go-sql-driver/mysql"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"stockexchange/rpc/model"
	"stockexchange/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel // 手动添加
}

// TODO 到后面业务再改用gorm 现在先用sqlx
/*func GormServiceContext(c config.Config) *ServiceContext {
	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,    // 使用单数表名，启用该选项，此时，表名是 `user` 而不是`users`
		},
	})
	//如果出错就panic
	if err != nil {
		panic(err)
	}
	// 自动同步更新表结构 这是建表语句
	// db.AutoMigrate(&models.User{})

	return &ServiceContext{Config: c, Model: db}
}*/

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//           func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel
		Model: model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动添加
		// TODO 怎样改成gorm
		// Model:  model.NewUserModel(GormServiceContext(c), c.Cache),
	}
}
