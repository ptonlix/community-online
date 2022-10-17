package logic

import (
	"context"
	"fmt"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/common/globalkey"
	"github/community-online/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FreshUserOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreshUserOnlineStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreshUserOnlineStatusLogic {
	return &FreshUserOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreshUserOnlineStatusLogic) FreshUserOnlineStatus(in *pb.FreshUserOnlineStatusReq) (*pb.FreshUserOnlineStatusResp, error) {
	// 统计在线用户数
	err := l.svcCtx.RedisClient.SetexCtx(l.ctx, fmt.Sprintf(globalkey.CacheUserOnlineKey, in.UserId), fmt.Sprint(in.UserId), globalkey.CacheUserOnlineKeyExp)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Set Online user failed, err:%v", err)
	}

	return &pb.FreshUserOnlineStatusResp{}, nil
}
