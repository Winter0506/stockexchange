// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"stockexchange/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/login",
				Handler: LoginHandler(serverCtx),
			},
		},
	)
}
