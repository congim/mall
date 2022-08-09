package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BannerModel = (*customBannerModel)(nil)

var (
	cacheBannerList = "cache:banner:list"
)

type (
	// BannerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBannerModel.
	BannerModel interface {
		bannerModel
		AllBanners(ctx context.Context, state int32) ([]*Banner, error)
	}

	customBannerModel struct {
		*defaultBannerModel
	}
)

// NewBannerModel returns a model for the database table.
func NewBannerModel(conn sqlx.SqlConn, c cache.CacheConf) BannerModel {
	return &customBannerModel{
		defaultBannerModel: newBannerModel(conn, c),
	}
}
func (c *customBannerModel) AllBanners(ctx context.Context, state int32) ([]*Banner, error) {
	var resp []*Banner
	err := c.QueryRowCtx(ctx, &resp, cacheBannerList, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `state` = ?", bannerRows, c.table)
		return conn.QueryRowsCtx(ctx, v, query, state)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
