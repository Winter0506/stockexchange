package types

type LoginMessage struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Role        int32  `json:"role"`
	AccessToken string `json:"token"`
}

type LoginMeta struct {
	Msg    string `json:"msg""`
	Status int16  `json:"status"`
}

// 每次改了 都要在types.go 重写
type RespUserLogin struct {
	LoginMessage `json:"message"`
	LoginMeta    `json:"meta"`
}

type ReqUserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6, max=10"`
}

type ReqUserRegister struct {
	UserName  string `json:"username" binding:"required,max=10"`
	Password  string `json:"password" binding:"required, min=6, max=10"`
	Email     string `json:"email" binding:"required, email"`
	Gender    string `json:"gender" binding:"required"`
	Captcha   string `json:"captcha" binding:"required,min=4,max=4"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}

type RespCaptha struct {
	Captcha   string `json:"captcha" binding:"required,min=4,max=4"`
	CaptchaId string `json:"captcha_id" binding:"required"`
	Msg       string `json:"msg""`
	Status    int16  `json:"status"`
}

type ReqUserId struct {
	Id int32 `path:"id"`
}

type ReqUserUpdate struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type RespUserDelete struct {
	LoginMeta `json:meta`
}

type ReqUserUpdateAdmin struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Role     int32  `json:"role"`
}

type ReqUserList struct {
	Pn    int `json:"pn,default=0"`
	PSize int `json:"psize,default=10"`
}

type RespUserList struct {
	UserList  []string `json:"userlist"`
	LoginMeta `json:"meta"`
}
