package models

import (
	"errors"
	"hello/internal/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email  string `gorm:"index:Email;type:varchar(13)"`
	Passwd string `gorm:"type:varchar(64)"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if len(u.Passwd) < 6 {
		return errors.New("密码太简单了")
	}
	u.Passwd = utils.Password(u.Passwd)
	return nil
}
