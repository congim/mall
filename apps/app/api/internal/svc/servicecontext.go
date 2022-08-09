package svc

import (
	"github.com/congim/mall/apps/app/api/internal/config"
	"github.com/congim/mall/apps/home/rpc/home"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	HomeRPC home.Home
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		HomeRPC: home.NewHome(zrpc.MustNewClient(c.HomeRPC)),
	}
}
