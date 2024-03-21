// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "Nft-Go/gateway/internal/handler/auth"
	nft "Nft-Go/gateway/internal/handler/nft"
	user "Nft-Go/gateway/internal/handler/user"
	"Nft-Go/gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: auth.UploadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: auth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: auth.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/nft/BuyFromPool",
				Handler: nft.BuyFromPoolHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/CreateActivity",
				Handler: nft.CreateActivityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/CreatePool",
				Handler: nft.CreatePoolHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/nft/GetAllActivity",
				Handler: nft.GetAllActivityHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/nft/GetAllDc",
				Handler: nft.GetAllDcHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/nft/GetAllPool",
				Handler: nft.GetAllPoolHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GetDcById",
				Handler: nft.GetDcByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GetDcFromActivity",
				Handler: nft.GetDcFromActivityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GetDigitalCollectionHistory",
				Handler: nft.GetDigitalCollectionHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GetMessageByHash",
				Handler: nft.GetMessageByHashHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/nft/GetMyDc",
				Handler: nft.GetMyDcHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/nft/GetMyPool",
				Handler: nft.GetMyPoolHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GetOneActivity",
				Handler: nft.GetOneActivityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GetPoolById",
				Handler: nft.GetPoolByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/GiveDc",
				Handler: nft.GiveDcHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/SearchActivities",
				Handler: nft.SearchActivitiesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/SelectDc",
				Handler: nft.SelectDcHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/nft/SelectPool",
				Handler: nft.SelectPoolHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/notice/all",
					Handler: user.GetAllNoticeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/id",
					Handler: user.GetNoticeByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/title",
					Handler: user.GetNoticeByTitleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/notice/username",
					Handler: user.GetUserInfoByUserNameHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/refreshToken",
					Handler: user.RefreshTokenHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
	)
}