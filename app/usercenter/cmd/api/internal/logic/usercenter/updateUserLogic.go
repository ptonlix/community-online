package usercenter

import (
	"context"

	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) error {

	_, err := l.svcCtx.UsercenterRpc.UpdateUser(l.ctx, &usercenter.UpdateUserReq{
		Id:       req.Id,
		Nickname: req.Nickname,
		Sex:      req.Sex,
		Avatar:   req.Avatar,
		Info:     req.Info,
	})
	if err != nil {
		return err
	}

	return nil
}
