package logic

import (
	"context"

	"github/community-online/app/usercenter/cmd/rpc/internal/svc"
	"github/community-online/app/usercenter/cmd/rpc/pb"
	"github/community-online/app/usercenter/cmd/rpc/usercenter"
	"github/community-online/app/usercenter/model"
	"github/community-online/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAuthKeyLogic {
	return &GetUserAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByAuthKeyLogic) GetUserAuthByAuthKey(in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err : %v , in : %+v", err, in)
	}

	var respUserAuth usercenter.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)

	return &pb.GetUserAuthByAuthKeyResp{
		UserAuth: &respUserAuth,
	}, nil
}
