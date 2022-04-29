package svc

import (
	"go-zero/vendor/github.com/zeromicro/go-zero/rest"
	"shorturl/api/internal/config"
	"shorturl/api/internal/middleware"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Transformer    transformer.Transformer //增加数据类型
	AuthMiddleware rest.Middleware
}

//传递服务依赖
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Transformer:    transformer.NewTransformer(zrpc.MustNewClient(c.Transformer)), //服务传递 允许访问
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
