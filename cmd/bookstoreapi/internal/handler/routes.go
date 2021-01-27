// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"bookstore/cmd/bookstoreapi/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/hello",
				Handler: HelloHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/check",
				Handler: CheckHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: AddHandler(serverCtx),
			},
		},
	)
}
