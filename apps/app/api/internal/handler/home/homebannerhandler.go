package home

import (
	"net/http"

	"github.com/congim/mall/apps/app/api/internal/logic/home"
	"github.com/congim/mall/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
