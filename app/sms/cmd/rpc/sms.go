package main

import (
	"flag"
	"fmt"
	"os"

	"github/community-online/app/sms/cmd/rpc/internal/config"
	"github/community-online/app/sms/cmd/rpc/internal/server"
	"github/community-online/app/sms/cmd/rpc/internal/svc"
	"github/community-online/app/sms/cmd/rpc/pb"
	"github/community-online/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 获取腾讯云密钥

	if c.SmsConfig.Secret["SecretId"] == "" {
		c.SmsConfig.Secret["SecretId"] = os.Getenv("TENCENTCLOUD_SECRET_ID")
	}
	if c.SmsConfig.Secret["SecretKey"] == "" {
		c.SmsConfig.Secret["SecretKey"] = os.Getenv("TENCENTCLOUD_SECRET_KEY")
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSmsServer(grpcServer, server.NewSmsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
