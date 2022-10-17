package logic

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github/community-online/app/sms/cmd/rpc/internal/logic/tencentsms"
	"github/community-online/app/sms/cmd/rpc/internal/svc"
	"github/community-online/app/sms/cmd/rpc/pb"
	"github/community-online/common/globalkey"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMsgCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMsgCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMsgCodeLogic {
	return &GetMsgCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMsgCodeLogic) GetMsgCode(in *pb.GetMsgCodeReq) (*pb.GetMsgCodeResp, error) {
	//_, randCode := l.selectSmsClient("tencent", in) // TEST
	smsClient, randCode := l.selectSmsClient("tencent", in)
	go smsClient.Send()

	// 将验证码和手机号保存到redis中
	l.svcCtx.RedisClient.SetexCtx(l.ctx, fmt.Sprintf(globalkey.CacheSmsPhoneKey, in.MsgType, in.PhoneNum), randCode, globalkey.CacheSmsPhoneKeyExp)

	return &pb.GetMsgCodeResp{MsgCode: randCode}, nil
}

func (l *GetMsgCodeLogic) selectSmsClient(option string, in *pb.GetMsgCodeReq) (SmsClient, string) {

	var randCode string = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	var client SmsClient
	switch option {
	case "tencent":
		templateSet := []string{randCode, strconv.Itoa(globalkey.CacheSmsPhoneKeyExp / 60)}
		smsRequest := tencentsms.NewSmsRequest(&l.svcCtx.Config.SmsConfig, tencentsms.WithPhoneNumberSet([]string{in.PhoneNum}),
			tencentsms.WithTemplateIdSet(l.svcCtx.Config.SmsConfig.TemplateIds[in.MsgType].(string)), tencentsms.WithTemplateParamSet(templateSet))
		client = tencentsms.NewSmsClient(tencentsms.WithRequest(*smsRequest), tencentsms.WithCredential(l.svcCtx.Config.SmsConfig))
	case "ali":
	}

	return client, randCode
}
