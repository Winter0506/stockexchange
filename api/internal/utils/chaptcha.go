package utils

import (
	"github.com/mojocn/base64Captcha"
	//"github.com/tal-tech/go-zero/core/logx"
	//"net/http"
)

var Store = base64Captcha.DefaultMemStore

func GetCaptcha() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, err := cp.Generate()
	return id, b64s, err
}
