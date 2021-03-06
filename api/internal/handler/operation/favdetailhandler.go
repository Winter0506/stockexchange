package operation

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/operation"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func FavDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUserFav
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := operation.NewFavDetailLogic(r.Context(), ctx)
		resp, err := l.FavDetail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
