package stock

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/stock"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func DetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqStockId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := stock.NewDetailLogic(r.Context(), ctx)
		resp, err := l.Detail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
