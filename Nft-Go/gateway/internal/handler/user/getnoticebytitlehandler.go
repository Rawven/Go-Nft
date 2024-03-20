package user

import (
	"net/http"

	"Nft-Go/gateway/internal/logic/user"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetNoticeByTitleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TitleNoticeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetNoticeByTitleLogic(r.Context(), svcCtx)
		resp, err := l.GetNoticeByTitle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
