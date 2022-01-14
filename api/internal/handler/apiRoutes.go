package handler

import (
	"net/http"
	"stockexchange/api/internal/handler/operation"
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
				Method:  http.MethodGet,
				Path:    "/api/v1/stock",
				Handler: stock.GetStockListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/stock/searchcode",
				Handler: stock.GetStockByCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/stock/searchname",
				Handler: stock.GetStockByNameHandler(serverCtx),
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

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/operation/detail",
				Handler: operation.FavDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/operation/add",
				Handler: operation.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/operation/delete",
				Handler: operation.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/operation/userfav",
				Handler: operation.UserFavHandler(serverCtx),
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
					Path:    "/api/v1/operation/stockfav",
					Handler: operation.StockFavHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
