package usercenter

import (
	"context"

	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (*types.GetCaptchaResp, error) {
	getCaptchaResp, err := l.svcCtx.UsercenterRpc.GetCaptcha(l.ctx, &usercenter.GetCaptchaReq{})
	if err != nil {
		return nil, err
	}

	var resp types.GetCaptchaResp
	_ = copier.Copy(&resp, getCaptchaResp)
	return &resp, nil
}
