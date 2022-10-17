package svc

import (
	"github/community-online/app/sms/cmd/rpc/sms"
	"github/community-online/app/usercenter/cmd/api/internal/config"
	"github/community-online/app/usercenter/cmd/api/internal/middleware"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	SmsRpc        sms.Sms
	OnlineStatus  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		SmsRpc:        sms.NewSms(zrpc.MustNewClient(c.SmsRpcConf)),
		OnlineStatus:  middleware.NewOnlineStatusMiddleware(usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf))).Handle,
	}
}
