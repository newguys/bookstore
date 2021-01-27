package logic

import (
	"context"

	"bookstore/cmd/bookstorerpc/bookstore"
	"bookstore/cmd/bookstorerpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *bookstore.AddReq) (*bookstore.AddResp, error) {
	// todo: add your logic here and delete this line

	return &bookstore.AddResp{}, nil
}
