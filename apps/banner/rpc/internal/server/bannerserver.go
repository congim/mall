// Code generated by goctl. DO NOT EDIT!
// Source: banner.proto

package server

import (
	"context"

	"github.com/congim/mall/apps/banner/rpc/banner"
	"github.com/congim/mall/apps/banner/rpc/internal/logic"
	"github.com/congim/mall/apps/banner/rpc/internal/svc"
)

type BannerServer struct {
	svcCtx *svc.ServiceContext
	banner.UnimplementedBannerServer
}

func NewBannerServer(svcCtx *svc.ServiceContext) *BannerServer {
	return &BannerServer{
		svcCtx: svcCtx,
	}
}

func (s *BannerServer) HomeBanner(ctx context.Context, in *banner.BannerReq) (*banner.BannerRes, error) {
	l := logic.NewHomeBannerLogic(ctx, s.svcCtx)
	return l.HomeBanner(in)
}