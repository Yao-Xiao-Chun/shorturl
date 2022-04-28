// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

package server

import (
	"context"

	"shorturl/rpc/transform/internal/logic"
	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/transform"
)

type TransformerServer struct {
	svcCtx *svc.ServiceContext
	transform.UnimplementedTransformerServer
}

func NewTransformerServer(svcCtx *svc.ServiceContext) *TransformerServer {
	return &TransformerServer{
		svcCtx: svcCtx,
	}
}

func (s *TransformerServer) Expand(ctx context.Context, in *transform.ExpandReq) (*transform.ExpandResp, error) {
	l := logic.NewExpandLogic(ctx, s.svcCtx)
	return l.Expand(in)
}

func (s *TransformerServer) Shorten(ctx context.Context, in *transform.ShortenReq) (*transform.ShortenResp, error) {
	l := logic.NewShortenLogic(ctx, s.svcCtx)
	return l.Shorten(in)
}

func (s *TransformerServer) Appversion(ctx context.Context, in *transform.CheckversionReq) (*transform.CheckversionResp, error) {
	l := logic.NewAppversionLogic(ctx, s.svcCtx)
	return l.Appversion(in)
}
