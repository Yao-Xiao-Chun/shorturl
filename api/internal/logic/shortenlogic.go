package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//转换为短码
func (l *ShortenLogic) Shorten(req *types.ShortenReq) (types.ShortenResp, error) {

	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{Url: req.Url})

	if err != nil {
		return types.ShortenResp{}, err
	}

	return types.ShortenResp{Shorten: resp.Shorten}, nil
}
