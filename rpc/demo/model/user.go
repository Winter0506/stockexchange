package model

/*type BaseModel struct {
	ID        int32     `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
	IsDeleted bool      `gorm:"column:isDeleted;default:0;type:int comment '是否删除, 0否1是'"`
}*/

type User struct {
	BaseModel
	UserName string `gorm:"column:username;unique,type:varchar(20)"`
	Password string `gorm:"column:password;type:varchar(100);not null"`
	Email    string `gorm:"column:email;unique;type:varchar(50);not null"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"`
	Role     int    `gorm:"column:role;default:1;type:int comment '1表示普通用户, 2表示管理员'"`
}
