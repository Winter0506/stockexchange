package handler

import (
	"fmt"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"stockexchange/api/internal/logic"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUserLogin
		fmt.Println("到达这里0")
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println("到达这里1")
		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
