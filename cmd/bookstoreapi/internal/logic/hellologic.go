package logic

import (
	"context"
	"fmt"

	"bookstore/cmd/bookstoreapi/internal/svc"
	"bookstore/cmd/bookstoreapi/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) HelloLogic {
	return HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req types.HelloReq) (*types.HelloResp, error) {
	// todo: add your logic here and delete this line
	msg := fmt.Sprintf("Let's roll, %s", req.User)
	return &types.HelloResp{Msg:msg}, nil
}
