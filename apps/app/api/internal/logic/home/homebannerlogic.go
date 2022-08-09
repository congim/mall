package home

import (
	"context"

	"github.com/congim/mall/apps/app/api/internal/svc"
	"github.com/congim/mall/apps/app/api/internal/types"
	"github.com/congim/mall/apps/home/rpc/home"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerLogic {
	return &HomeBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBannerLogic) HomeBanner() (*types.BannerRes, error) {
	rpcBanners, err := l.svcCtx.HomeRPC.Banner(l.ctx, &home.BannerReq{})
	if err != nil {
		logx.Error("home banner fail", logx.Field("error", err))
		return nil, err
	}

	var banners []*types.Banner
	for _, b := range rpcBanners.Banners {
		banners = append(banners, &types.Banner{
			Id:    b.Id,
			Name:  b.Name,
			Image: b.Image,
			Url:   b.Url,
		})
	}
	logx.Info("home banner list", logx.Field("banners", banners))
	return &types.BannerRes{Banners: banners}, nil
}
