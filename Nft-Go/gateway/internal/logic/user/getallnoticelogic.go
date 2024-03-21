package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/util"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllNoticeLogic {
	return &GetAllNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllNoticeLogic) GetAllNotice() (resp *types.CommonResponse, err error) {
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	notices, err := api.GetUserClient().GetAllNotices(ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}
	data, err := jsonx.MarshalToString(
		notices.GetNotices())
	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Code:    200,
		Data:    data,
		Message: "",
	}, nil
}