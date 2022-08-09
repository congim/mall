package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/congim/mall/apps/banner/rpc/banner"
	"github.com/congim/mall/apps/banner/rpc/internal/config"
	"github.com/congim/mall/apps/banner/rpc/internal/server"
	"github.com/congim/mall/apps/banner/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/banner.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewBannerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		banner.RegisterBannerServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func init() {
	logx.DisableStat()
}
