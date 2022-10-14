package usercenter

import (
	"net/http"

	"github/community-online/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/community-online/app/usercenter/cmd/api/internal/logic/usercenter"
	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
)

func VerfyCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerfyCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := usercenter.NewVerfyCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.VerfyCaptcha(&req)
		result.HttpResult(r, w, resp, err)
	}
}
