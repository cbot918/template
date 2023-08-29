package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"users/internal/logic/user"
	"users/internal/svc"
	"users/internal/types"
)

func UsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUsersLogic(r.Context(), svcCtx)
		resp, err := l.Users(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
