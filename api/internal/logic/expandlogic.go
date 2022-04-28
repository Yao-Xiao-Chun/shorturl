package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//短连接拓展方法
func (l *ExpandLogic) Expand(req *types.ExpandReq) (types.ExpandResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transform.ExpandReq{
		Shorten: req.Shorten,
	})

	if err != nil {

		return types.ExpandResp{}, err
	}

	return types.ExpandResp{Url: resp.Url}, nil
}
