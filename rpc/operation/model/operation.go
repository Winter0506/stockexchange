package model

type UserFav struct {
	BaseModel
	User  int32 `gorm:"type:int;index:idx_user_stock,unique"`
	Stock int32 `gorm:"type:int;index:idx_user_stock,unique"`
}

func (UserFav) TableName() string {
	return "userfav"
}
