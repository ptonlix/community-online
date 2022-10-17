package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/common/globalkey"
	"github/community-online/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnlineUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnlineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlineUserLogic {
	return &GetOnlineUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnlineUserLogic) GetOnlineUser(in *pb.GetOnlineUserReq) (*pb.GetOnlineUserResp, error) {
	// 统计在线用户数
	userList, err := l.svcCtx.RedisClient.KeysCtx(l.ctx, fmt.Sprintf(globalkey.CacheUserOnlineKey, "*"))
	logx.Error(userList)
	length := len(userList)
	if length == 0 {
		return &pb.GetOnlineUserResp{}, nil
	}
	onlineUserList := make([]int64, length)
	for k, v := range userList {
		onlineUserList[k], _ = strconv.ParseInt(strings.Split(v, ":")[1], 10, 0)
	}
	logx.Error(onlineUserList)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Set Online user failed, err:%v", err)
	}

	return &pb.GetOnlineUserResp{OnlineUser: onlineUserList}, nil
}
