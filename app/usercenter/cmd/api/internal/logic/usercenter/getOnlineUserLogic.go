package usercenter

import (
	"context"

	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnlineUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnlineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlineUserLogic {
	return &GetOnlineUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnlineUserLogic) GetOnlineUser() (resp *types.GetOnlineUserResp, err error) {
	list, err := l.svcCtx.UsercenterRpc.GetOnlineUser(l.ctx, &usercenter.GetOnlineUserReq{})
	if err != nil {
		return nil, err
	}

	var onlineUserList types.GetOnlineUserResp
	_ = copier.Copy(&onlineUserList, list)

	return &onlineUserList, nil
}
