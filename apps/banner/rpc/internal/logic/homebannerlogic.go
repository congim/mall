package logic

import (
	"context"

	"github.com/congim/mall/apps/banner/rpc/banner"
	"github.com/congim/mall/apps/banner/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	BANNER_STATE_TYPE_USE   int32 = 1
	BANNER_STATE_TYPE_UNUSE int32 = 2
)

type HomeBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerLogic {
	return &HomeBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBannerLogic) HomeBanner(in *banner.BannerReq) (*banner.BannerRes, error) {
	dbBanners, err := l.svcCtx.BannerModel.AllBanners(l.ctx, BANNER_STATE_TYPE_USE)
	if err != nil {
		logx.Error("all banners fail", logx.Field("error", err))
		return nil, err
	}
	var banners []*banner.BannerInfo
	for _, b := range dbBanners {
		banners = append(banners, &banner.BannerInfo{
			Id:    b.Id,
			Name:  b.Name,
			Image: b.Image,
			Url:   b.Url,
		})
	}

	return &banner.BannerRes{
		Banners: banners,
	}, nil
}
