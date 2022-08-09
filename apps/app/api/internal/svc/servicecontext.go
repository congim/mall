package svc

import (
	"github.com/congim/mall/apps/app/api/internal/config"
	"github.com/congim/mall/apps/banner/rpc/banner"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	BannerRPC banner.Banner
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BannerRPC: banner.NewBanner(zrpc.MustNewClient(c.BannerRPC)),
	}
}
