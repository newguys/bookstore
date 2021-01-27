package logic

import (
	"context"
	"fmt"

	"bookstore/cmd/bookstorerpc/bookstore"
	"bookstore/cmd/bookstorerpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HelloLogic) Hello(in *bookstore.Request) (*bookstore.Response, error) {
	// todo: add your logic here and delete this line
	msg := fmt.Sprintf("hello %s welcome to  go-zero'world",in.User)
	return &bookstore.Response{Msg: msg}, nil
}
