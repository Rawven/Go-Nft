package nft

import (
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPoolLogic {
	return &GetAllPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPoolLogic) GetAllPool() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}