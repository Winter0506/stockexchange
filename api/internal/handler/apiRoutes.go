package handler

import (
	"net/http"
	"stockexchange/api/internal/handler/inventory"
	"stockexchange/api/internal/handler/operation"
	"stockexchange/api/internal/handler/order"
	"stockexchange/api/internal/handler/user"

	stock "stockexchange/api/internal/handler/stock"
	"stockexchange/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// stock路由
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

	// user路由
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

	// operation路由
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/operation/detail",
				Handler: operation.FavDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/operation/add",
				Handler: operation.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/v1/operation/delete",
				Handler: operation.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
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
					Method:  http.MethodGet,
					Path:    "/api/v1/operation/stockfav",
					Handler: operation.StockFavHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	// Inventory
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Admin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/inventory/setinv",
					Handler: inventory.SetInvHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/inventory/invdetail",
					Handler: inventory.InvDetailHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
	// Order
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/order/account/:id",
				Handler: order.UserAccountDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/order/account/create",
				Handler: order.CreateUserAccountHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/api/v1/order/account/update",
				Handler: order.UpdateUserAccountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/order/hold/:id",
				Handler: order.HoldPositionListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/order/trust/createtrust",
				Handler: order.CreateTrustHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
