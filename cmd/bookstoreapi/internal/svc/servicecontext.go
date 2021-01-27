package svc

import (
	"bookstore/cmd/bookstoreapi/internal/config"
	"bookstore/cmd/bookstorerpc/bookstoreclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	BookStore bookstoreclient.Bookstore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BookStore:bookstoreclient.NewBookstore(zrpc.MustNewClient(c.BookStore)),
	}
}
