package handler

import (
	"net/http"
	"stockexchange/api/internal/handler/user"

	stock "stockexchange/api/internal/handler/stock"
	"stockexchange/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/stock/:id",
				Handler: stock.DetailHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Admin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/stock/create",
					Handler: stock.CreateHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
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
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/user/detail/:id",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/api/v1/user/update",
				Handler: user.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/v1/user/delete/:id",
				Handler: user.DeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Admin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/user",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/user",
					Handler: user.UpdateAdminHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
