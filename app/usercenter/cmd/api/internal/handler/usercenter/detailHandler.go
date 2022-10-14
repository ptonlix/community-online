package usercenter

import (
	"net/http"

	"github/community-online/common/result"

	"github/community-online/app/usercenter/cmd/api/internal/logic/usercenter"
	"github/community-online/app/usercenter/cmd/api/internal/svc"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := usercenter.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
		result.HttpResult(r, w, resp, err)
	}
}
