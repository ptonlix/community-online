package middleware

import (
	"github/community-online/app/usercenter/cmd/rpc/usercenter"
	"github/community-online/common/ctxdata"
	"net/http"
)

type OnlineStatusMiddleware struct {
	UsercenterRpc usercenter.Usercenter
}

func NewOnlineStatusMiddleware(uRpc usercenter.Usercenter) *OnlineStatusMiddleware {
	return &OnlineStatusMiddleware{UsercenterRpc: uRpc}
}

func (m *OnlineStatusMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 刷新用户在线状态
		userId := ctxdata.GetUidFromCtx(r.Context())
		m.UsercenterRpc.FreshUserOnlineStatus(r.Context(), &usercenter.FreshUserOnlineStatusReq{
			UserId: userId,
		})
		// Passthrough to next handler if need
		next(w, r)
	}
}
