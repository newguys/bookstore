package logic

import (
	"context"

	"bookstore/cmd/bookstorerpc/bookstore"
	"bookstore/cmd/bookstorerpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *bookstore.CheckReq) (*bookstore.CheckResp, error) {
	// todo: add your logic here and delete this line
	if in.Token == "Evildoer" {
		return &bookstore.CheckResp{Ok: true},nil
	}
	return &bookstore.CheckResp{Ok: false}, nil
}
