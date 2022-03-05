package logic

import (
	"context"

	"home-bank/something/internal/svc"
	"home-bank/something/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SomethingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSomethingLogic(ctx context.Context, svcCtx *svc.ServiceContext) SomethingLogic {
	return SomethingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SomethingLogic) Something(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
