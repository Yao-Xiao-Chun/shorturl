package logic

import (
	"context"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppversionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppversionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppversionLogic {
	return &AppversionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppversionLogic) Appversion(in *transform.CheckversionReq) (*transform.CheckversionResp, error) {
	// todo: add your logic here and delete this line

	return &transform.CheckversionResp{Res: true}, nil
}
