package stock

import (
	"fmt"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"stockexchange/api/internal/logic/stock"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
)

func CreateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqStockCreate

		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(req)
			httpx.Error(w, err)
			return
		}

		l := stock.NewCreateLogic(r.Context(), ctx)
		resp, err := l.Create(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
