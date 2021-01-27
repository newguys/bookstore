package handler

import (
	"net/http"

	"bookstore/cmd/bookstoreapi/internal/logic"
	"bookstore/cmd/bookstoreapi/internal/svc"
	"bookstore/cmd/bookstoreapi/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HelloHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), ctx)
		resp, err := l.Hello(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
