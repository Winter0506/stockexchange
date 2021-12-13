package handler

import (
	"net/http"

	"hello/internal/logic/open"
	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AuthorizationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAuthorizationLogic(r.Context(), ctx)
		resp, err := l.Authorization(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
