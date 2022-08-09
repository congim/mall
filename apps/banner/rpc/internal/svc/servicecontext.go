package svc

import (
	"github.com/congim/mall/apps/banner/rpc/internal/config"
	"github.com/congim/mall/apps/banner/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	BannerModel model.BannerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DBSource)
	return &ServiceContext{
		Config:      c,
		BannerModel: model.NewBannerModel(dbConn, c.Cache),
	}
}
