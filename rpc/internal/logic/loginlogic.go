package logic

import (
	"context"

	"github.com/flyskea/lightalk-user-rpc/rpc/internal/svc"
	"github.com/flyskea/lightalk-user-rpc/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.UserDetail, error) {
	// todo: add your logic here and delete this line

	return &user.UserDetail{}, nil
}
