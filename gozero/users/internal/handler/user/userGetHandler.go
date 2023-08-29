package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"users/internal/logic/user"
	"users/internal/svc"
	"users/internal/types"
)

func UserGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserGetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserGetLogic(r.Context(), svcCtx)
		resp, err := l.UserGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
