package logic

import (
	"context"

	"github.com/congim/mall/apps/home/rpc/home"
	"github.com/congim/mall/apps/home/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	BANNER_STATE_TYPE_USE   int32 = 1
	BANNER_STATE_TYPE_UNUSE int32 = 2
)

type BannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerLogic {
	return &BannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BannerLogic) Banner(in *home.BannerReq) (*home.BannerRes, error) {
	dbBanners, err := l.svcCtx.BannerModel.AllBanners(l.ctx, BANNER_STATE_TYPE_USE)
	if err != nil {
		logx.Error("all banners fail", logx.Field("error", err))
		return nil, err
	}
	var banners []*home.BannerInfo
	for _, b := range dbBanners {
		banners = append(banners, &home.BannerInfo{
			Id:    b.Id,
			Name:  b.Name,
			Image: b.Image,
			Url:   b.Url,
		})
	}

	return &home.BannerRes{
		Banners: banners,
	}, nil
}
