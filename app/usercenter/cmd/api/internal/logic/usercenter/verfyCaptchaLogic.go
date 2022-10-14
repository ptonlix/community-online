package usercenter

import (
	"context"

	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerfyCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerfyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerfyCaptchaLogic {
	return &VerfyCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerfyCaptchaLogic) VerfyCaptcha(req *types.VerfyCaptchaReq) (*types.VerfyCaptchaResp, error) {
	verfyCaptchaResp, err := l.svcCtx.UsercenterRpc.VerfyCaptcha(l.ctx, &usercenter.VerfyCaptchaReq{
		Dots:       req.Dots,
		CaptchaKey: req.CaptchaKey,
	})
	if err != nil {
		return nil, err
	}

	var resp types.VerfyCaptchaResp
	_ = copier.Copy(&resp, verfyCaptchaResp)
	return &resp, nil
}
