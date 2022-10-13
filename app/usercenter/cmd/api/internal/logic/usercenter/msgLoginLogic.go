package usercenter

import (
	"context"

	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"
	"github/community-online/app/usercenter/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgLoginLogic {
	return &MsgLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgLoginLogic) MsgLogin(req *types.MsgLoginReq) (*types.MsgLoginResp, error) {
	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		AuthType: model.UserAuthTypeMessage,
		AuthKey:  req.Mobile,
		Password: req.Msgcode,
	})
	if err != nil {
		return nil, err
	}

	var resp types.MsgLoginResp
	_ = copier.Copy(&resp, loginResp)

	return &resp, nil
}
