package logic

import (
	"context"
	"fmt"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/app/usercenter/model"
	"github/community-online/common/globalkey"
	"github/community-online/common/tool"
	"github/community-online/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChanegPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChanegPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChanegPwdLogic {
	return &ChanegPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChanegPwdLogic) ChanegPwd(in *pb.ChangePwdReq) (*pb.ChangePwdResp, error) {
	// 验证手机验证码是否正确
	redisCode, err := l.svcCtx.RedisClient.GetCtx(l.ctx, fmt.Sprintf(globalkey.CacheSmsPhoneKey, globalkey.SmsChangePwd, in.Mobile))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询短信验证码失败，mobile:%s,err:%v", in.Mobile, err)
	}
	if redisCode != in.Msgcode {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_MESSAGE_ERROR), "短信验证码失败，mobile:%s,redisCode:%v, inputCode:%v", in.Mobile, redisCode, in.Msgcode)
	}
	// 删除验证码
	l.svcCtx.RedisClient.DelCtx(l.ctx, fmt.Sprintf(globalkey.CacheSmsPhoneKey, globalkey.SmsLogin, in.Mobile))

	// 修改密码
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", in.Mobile, err)
	}
	user.Password = tool.Md5ByString(in.NewPassword)
	err = l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Change db user password err:%v,user:%+v", err, user)
	}

	return &pb.ChangePwdResp{}, nil
}
