package inventory

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/inventory"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func SetInvHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqSetInv
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := inventory.NewSetInvLogic(r.Context(), ctx)
		resp, err := l.SetInv(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
