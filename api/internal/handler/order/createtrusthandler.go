package order

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/order"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func CreateTrustHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqCreateTrustItem
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewCreateTrustLogic(r.Context(), ctx)
		resp, err := l.CreateTrust(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
