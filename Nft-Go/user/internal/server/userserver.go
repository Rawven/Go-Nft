// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"Nft-Go/user/internal/logic"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.Response, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.Response, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Logout(ctx context.Context, in *user.Empty) (*user.Response, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *UserServer) RefreshTokens(ctx context.Context, in *user.Empty) (*user.Response, error) {
	l := logic.NewRefreshTokensLogic(ctx, s.svcCtx)
	return l.RefreshTokens(in)
}

func (s *UserServer) Upload(stream user.User_UploadServer) error {
	l := logic.NewUploadLogic(stream.Context(), s.svcCtx)
	return l.Upload(stream)
}

func (s *UserServer) GetAllNotices(ctx context.Context, in *user.Empty) (*user.NoticeList, error) {
	l := logic.NewGetAllNoticesLogic(ctx, s.svcCtx)
	return l.GetAllNotices(in)
}

func (s *UserServer) GetNoticeByTitle(ctx context.Context, in *user.TitleNoticeRequest) (*user.NoticeList, error) {
	l := logic.NewGetNoticeByTitleLogic(ctx, s.svcCtx)
	return l.GetNoticeByTitle(in)
}

func (s *UserServer) GetNoticeById(ctx context.Context, in *user.IdNoticeRequest) (*user.Notice, error) {
	l := logic.NewGetNoticeByIdLogic(ctx, s.svcCtx)
	return l.GetNoticeById(in)
}
