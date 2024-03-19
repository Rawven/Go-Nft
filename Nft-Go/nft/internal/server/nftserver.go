// Code generated by goctl. DO NOT EDIT.
// Source: nft.proto

package server

import (
	"Nft-Go/common/api/nft"
	"context"

	"Nft-Go/nft/internal/logic"
	"Nft-Go/nft/internal/svc"
)

type NftServer struct {
	svcCtx *svc.ServiceContext
	nft.UnimplementedNftServer
}

func NewNftServer(svcCtx *svc.ServiceContext) *NftServer {
	return &NftServer{
		svcCtx: svcCtx,
	}
}

func (s *NftServer) GetMessageByHash(ctx context.Context, in *nft.GetMessageByHashRequest) (*nft.GetMessageByHashDTO, error) {
	l := logic.NewGetMessageByHashLogic(ctx, s.svcCtx)
	return l.GetMessageByHash(in)
}

func (s *NftServer) CreateActivity(ctx context.Context, in *nft.CreateActivityRequest) (*nft.CommonResult, error) {
	l := logic.NewCreateActivityLogic(ctx, s.svcCtx)
	return l.CreateActivity(in)
}

func (s *NftServer) GetDcFromActivity(ctx context.Context, in *nft.GetDcFromActivityRequest) (*nft.CommonResult, error) {
	l := logic.NewGetDcFromActivityLogic(ctx, s.svcCtx)
	return l.GetDcFromActivity(in)
}

func (s *NftServer) GetAllActivity(ctx context.Context, in *nft.NftEmpty) (*nft.ActivityPageVOList, error) {
	l := logic.NewGetAllActivityLogic(ctx, s.svcCtx)
	return l.GetAllActivity(in)
}

func (s *NftServer) GetOneActivity(ctx context.Context, in *nft.GetOneActivityRequest) (*nft.ActivityDetailsVO, error) {
	l := logic.NewGetOneActivityLogic(ctx, s.svcCtx)
	return l.GetOneActivity(in)
}

func (s *NftServer) SearchActivities(ctx context.Context, in *nft.SearchActivitiesRequest) (*nft.ActivityPageVOList, error) {
	l := logic.NewSearchActivitiesLogic(ctx, s.svcCtx)
	return l.SearchActivities(in)
}

func (s *NftServer) GiveDc(ctx context.Context, in *nft.GiveDcRequest) (*nft.CommonResult, error) {
	l := logic.NewGiveDcLogic(ctx, s.svcCtx)
	return l.GiveDc(in)
}

func (s *NftServer) GetAllDc(ctx context.Context, in *nft.NftEmpty) (*nft.DcPageVOList, error) {
	l := logic.NewGetAllDcLogic(ctx, s.svcCtx)
	return l.GetAllDc(in)
}

func (s *NftServer) SelectDc(ctx context.Context, in *nft.SelectDcRequest) (*nft.DcPageVOList, error) {
	l := logic.NewSelectDcLogic(ctx, s.svcCtx)
	return l.SelectDc(in)
}

func (s *NftServer) GetDcById(ctx context.Context, in *nft.GetDcByIdRequest) (*nft.DcDetailVO, error) {
	l := logic.NewGetDcByIdLogic(ctx, s.svcCtx)
	return l.GetDcById(in)
}

func (s *NftServer) GetMyDc(ctx context.Context, in *nft.NftEmpty) (*nft.DcPageVOList, error) {
	l := logic.NewGetMyDcLogic(ctx, s.svcCtx)
	return l.GetMyDc(in)
}

func (s *NftServer) GetDigitalCollectionHistory(ctx context.Context, in *nft.GetDigitalCollectionHistoryRequest) (*nft.CollectionMessageOnChainVO, error) {
	l := logic.NewGetDigitalCollectionHistoryLogic(ctx, s.svcCtx)
	return l.GetDigitalCollectionHistory(in)
}

func (s *NftServer) CreatePool(ctx context.Context, in *nft.CreatePoolRequest) (*nft.CommonResult, error) {
	l := logic.NewCreatePoolLogic(ctx, s.svcCtx)
	return l.CreatePool(in)
}

func (s *NftServer) BuyFromPool(ctx context.Context, in *nft.BuyFromPoolRequest) (*nft.CommonResult, error) {
	l := logic.NewBuyFromPoolLogic(ctx, s.svcCtx)
	return l.BuyFromPool(in)
}

func (s *NftServer) SelectPool(ctx context.Context, in *nft.SelectPoolRequest) (*nft.PoolPageVOList, error) {
	l := logic.NewSelectPoolLogic(ctx, s.svcCtx)
	return l.SelectPool(in)
}

func (s *NftServer) GetPoolById(ctx context.Context, in *nft.GetPoolByIdRequest) (*nft.PoolDetailsVO, error) {
	l := logic.NewGetPoolByIdLogic(ctx, s.svcCtx)
	return l.GetPoolById(in)
}

func (s *NftServer) GetAllPool(ctx context.Context, in *nft.NftEmpty) (*nft.PoolPageVOList, error) {
	l := logic.NewGetAllPoolLogic(ctx, s.svcCtx)
	return l.GetAllPool(in)
}

func (s *NftServer) GetMyPool(ctx context.Context, in *nft.NftEmpty) (*nft.PoolPageVOList, error) {
	l := logic.NewGetMyPoolLogic(ctx, s.svcCtx)
	return l.GetMyPool(in)
}
