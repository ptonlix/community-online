package usercenter

import (
	"context"

	"github/community-online/app/sms/cmd/rpc/sms"
	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessagecodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessagecodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessagecodeLogic {
	return &MessagecodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessagecodeLogic) Messagecode(req *types.MsgCodeReq) (err error) {
	//调用RPC
	_, err = l.svcCtx.SmsRpc.GetMsgCode(l.ctx, &sms.GetMsgCodeReq{
		MsgType:  req.MsgType,
		PhoneNum: req.Mobile,
	})
	if err != nil {
		return err
	}
	return nil
}
