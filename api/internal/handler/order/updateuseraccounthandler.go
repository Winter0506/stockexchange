package order

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/order"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func UpdateUserAccountHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUserAccountUpdate
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewUpdateUserAccountLogic(r.Context(), ctx)
		resp, err := l.UpdateUserAccount(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
