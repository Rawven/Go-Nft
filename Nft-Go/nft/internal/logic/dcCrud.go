package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/model"
)

func GetDcPageVOList(dcInfos []*model.DcInfo) []*nft.DcPageVO {
	var dcPageVOList []*nft.DcPageVO
	for _, dcInfo := range dcInfos {
		dcPageVOList = append(dcPageVOList, &nft.DcPageVO{
			Url:         dcInfo.Cid,
			Name:        dcInfo.Name,
			DcId:        dcInfo.Id,
			Hash:        dcInfo.Hash,
			CreatorName: dcInfo.CreatorName,
			Price:       dcInfo.Price,
		})
	}
	return dcPageVOList
}
