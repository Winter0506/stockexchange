package handler

import (
	"net/http"
	"stockexchange/api/internal/handler/user"

	stock "stockexchange/api/internal/handler/stock"
	"stockexchange/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/stock/:id",
				Handler: stock.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/stock/create",
				Handler: stock.CreateHandler(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/user/captcha",
				Handler: user.CaptchaHandler(serverCtx),
			},
		},
	)
}
