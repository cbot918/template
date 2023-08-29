package health

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"users/internal/logic/health"
	"users/internal/svc"
	"users/internal/types"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := health.NewHealthLogic(r.Context(), svcCtx)
		resp, err := l.Health(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
