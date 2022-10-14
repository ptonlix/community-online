package logic

import (
	"context"
	"encoding/json"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/common/globalkey"
	"github/community-online/common/xerr"

	"github.com/pkg/errors"
	"github.com/wenlng/go-captcha/captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

var CaptchaGenerateError = xerr.NewErrCode(xerr.CAPTCHA_GENERATE_ERROR)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *pb.GetCaptchaReq) (*pb.GetCaptchaResp, error) {
	// 生成图形校验码
	// Captcha Single Instances
	capt := captcha.GetCaptcha()

	// Generate Captcha
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return nil, errors.Wrapf(CaptchaGenerateError, "err : %v ", err)
	}
	logx.Infof("Get Captcha Key: %s, Dots: %v", key, dots)
	// 缓存数据
	jsonDots, err := json.Marshal(dots)
	if err != nil {
		return nil, errors.Wrapf(CaptchaGenerateError, "err : %v ", err)
	}
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, key, string(jsonDots), globalkey.CacheCaptchaKeyExp)
	if err != nil {
		return nil, errors.Wrapf(CaptchaGenerateError, "err : %v ", err)
	}
	return &pb.GetCaptchaResp{
		ImageBase64: b64,
		ThumbBase64: tb64,
		CaptchaKey:  key,
	}, nil
}
