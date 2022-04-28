package svc

import (
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer //增加依赖
}

//传递服务依赖
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transformer)),
	}
}
