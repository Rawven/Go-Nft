package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllDcLogic {
	return &GetAllDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllDcLogic) GetAllDc(in *nft.Empty) (*nft.DcPageVOList, error) {
	mysql := db.GetMysql()
	//查找所有DcInfo 按照id排序
	var dcInfos []model.DcInfo
	mysql.Find(&model.DcInfo{}).Order("id").Find(&dcInfos)
	dcPageVOList := GetAllDc(dcInfos)
	return &nft.DcPageVOList{
		DcPageVO: dcPageVOList,
	}, nil
}