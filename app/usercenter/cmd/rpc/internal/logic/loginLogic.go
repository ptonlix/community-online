package logic

import (
	"context"
	"fmt"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"
	"github/community-online/app/usercenter/model"
	"github/community-online/common/globalkey"
	"github/community-online/common/tool"
	"github/community-online/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	switch in.AuthType {
	case model.UserAuthTypeSystem:
		return l.loginByMobile(in.AuthKey, in.Password)
	case model.UserAuthTypeMessage:
		return l.loginByMessage(in.AuthKey, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
}

// 手机密码登录
func (l *LoginLogic) loginByMobile(mobile, password string) (*pb.LoginResp, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return nil, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user.Id)
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

// 手机验证码登录
func (l *LoginLogic) loginByMessage(mobile, code string) (*pb.LoginResp, error) {
	// 验证手机验证码是否正确
	redisCode, err := l.svcCtx.RedisClient.Get(fmt.Sprintf(globalkey.CacheSmsPhoneKey, mobile))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询短信验证码失败，mobile:%s,err:%v", mobile, err)
	}
	if redisCode != code {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "短信验证码失败，mobile:%s,redisCode:%v, inputCode:%v", mobile, redisCode, code)
	}
	// 检测是否已注册
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	nickName := fmt.Sprintf("Community%s", mobile[7:])
	if user == nil {
		// 注册新账号
		reg := NewRegisterLogic(l.ctx, l.svcCtx)
		ret, err := reg.Register(&pb.RegisterReq{
			AuthKey:  mobile,
			AuthType: model.UserAuthTypeMessage,
			Mobile:   mobile,
			Nickname: nickName,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册新用户失败，mobile:%s,err:%v", mobile, err)
		}
		return &usercenter.LoginResp{
			AccessToken:  ret.AccessToken,
			AccessExpire: ret.AccessExpire,
			RefreshAfter: ret.RefreshAfter,
		}, nil
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user.Id)
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
