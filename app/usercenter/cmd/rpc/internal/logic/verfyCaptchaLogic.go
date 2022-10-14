package logic

import (
	"context"
	"encoding/json"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/common/xerr"

	"github.com/pkg/errors"
	"github.com/wenlng/go-captcha/captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

var CaptchaVerfyError = xerr.NewErrCode(xerr.USER_LOGIN_VERIFY_ERROR)

type VerfyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerfyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerfyCaptchaLogic {
	return &VerfyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerfyCaptchaLogic) VerfyCaptcha(in *pb.VerfyCaptchaReq) (*pb.VerfyCaptchaResp, error) {
	jsonDots, err := l.svcCtx.RedisClient.GetCtx(l.ctx, in.CaptchaKey)
	if err != nil {
		return nil, errors.Wrapf(CaptchaVerfyError, "err : %v ", err)
	}
	var dots map[int]captcha.CharDot
	err = json.Unmarshal([]byte(jsonDots), &dots)
	if err != nil {
		return nil, errors.Wrapf(CaptchaVerfyError, "err : %v ", err)
	}
	// 判断是否正确

	for k, v := range dots {
		if !captcha.CheckPointDist(in.Dots[k*2], in.Dots[k*2+1], int64(v.Dx), int64(v.Dy), int64(v.Width), int64(v.Height)) {
			logx.Infof("Verfy Captcha Code Failed! err: %v", CaptchaVerfyError)
			return &pb.VerfyCaptchaResp{Result: false}, nil
		}
	}
	return &pb.VerfyCaptchaResp{Result: true}, nil
}
