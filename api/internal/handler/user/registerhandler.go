package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/user"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUserRegister
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
