package utils

import (
	"github.com/mojocn/base64Captcha"
	//"github.com/tal-tech/go-zero/core/logx"
	//"net/http"
)

var Store = base64Captcha.DefaultMemStore

func GetCaptcha() {
	//driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, store)
	//id, b64s, err := cp.Generate()
	//if err != nil {
	//	logx.Errorf("生成验证码错误: ", err.Error())
	// 后面再写
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"msg":"生成验证码错误",
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"captchaId": id,
	//	"picPath": b64s,
	//})
}
