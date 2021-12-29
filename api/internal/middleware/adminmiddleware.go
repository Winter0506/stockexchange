package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

var (
	errorUserInfo = errors.New("用户信息错误")
	authDeny      = errors.New("用户无管理员权限")
)

const (
	key = `authorization`
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// jwtCode := r.Header.Get("authorization")
		role := r.Context().Value("role")
		roleInt, err := json.Number(fmt.Sprintf("%v", role)).Int64()
		if err != nil {
			httpx.Error(w, errorUserInfo)
			return
		}
		if roleInt != 1 {
			httpx.Error(w, authDeny)
		} else {
			next(w, r)
		}
	}
}

/*
func (m *UserCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get(userKey)
		jwtUserId := r.Context().Value("userId")

		userInt, err := json.Number(userId).Int64()
		if err != nil {
			httpx.Error(w, errorUserInfo)
			return
		}

		jwtInt, err := json.Number(fmt.Sprintf("%v", jwtUserId)).Int64()
		if err != nil {
			httpx.Error(w, errorUserInfo)
			return
		}

		if jwtInt != userInt {
			httpx.Error(w, authDeny)
			return
		}

		next(w, r)
	}
}
*/
