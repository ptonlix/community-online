package usercenter

import (
	"context"

	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePwdLogic) ChangePwd(req *types.ChangePwdReq) error {

	_, err := l.svcCtx.UsercenterRpc.ChanegPwd(l.ctx, &usercenter.ChangePwdReq{
		Mobile:      req.Mobile,
		Msgcode:     req.Msgcode,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return err
	}

	return nil
}
