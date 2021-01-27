package svc

import (
	"bookstore/cmd/bookstorerpc/internal/config"
	"bookstore/cmd/bookstorerpc/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewUserModel(sqlx.NewMysql(c.DataSource),c.Cache),
	}
}
