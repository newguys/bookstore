package logic

import (
	"context"

	"bookstore/cmd/bookstoreapi/internal/svc"
	"bookstore/cmd/bookstoreapi/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	// todo: add your logic here and delete this line
	if req.Token == "Evildoer" {
		return &types.CheckResp{Ok: true},nil
	}
	return &types.CheckResp{Ok: false}, nil
}
