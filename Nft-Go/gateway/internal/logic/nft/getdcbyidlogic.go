package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcByIdLogic {
	return &GetDcByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcByIdLogic) GetDcById(req *types.GetDcByIdRequest) (resp *types.CommonResponse, err error) {
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	id, err := api.GetNftClient().GetDcById(ctx, &nft.GetDcByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(id)
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
