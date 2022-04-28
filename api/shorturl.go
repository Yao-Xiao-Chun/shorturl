package main

import (
	"flag"
	"fmt"

	"shorturl/api/internal/config"
	"shorturl/api/internal/handler"
	"shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)          //建立一个新的上下文
	server := rest.MustNewServer(c.RestConf) //增加新的配置服务
	defer server.Stop()

	handler.RegisterHandlers(server, ctx) //服务启动 注册handle

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
