package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	global2 "Nft-Go/common/util"
	"Nft-Go/user/internal/model"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.Response, error) {
	//链上注册
	dubbo := api.GetBlcDubbo()
	result, err := dubbo.SignUp(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	//本地注册
	sha256 := cryptor.Sha256(in.GetPassword())
	mod := model.User{
		Username:   in.GetUsername(),
		Password:   sha256,
		Email:      in.GetEmail(),
		PrivateKey: result.GetPrivateKey(),
		Address:    result.GetAddress(),
		Avatar:     in.GetAvatar(),
	}
	mysql := db.GetMysql()
	rds := db.GetRedis()
	tx := mysql.Create(&mod)
	if tx.Error != nil {
		logger.Error("插入用户失败", tx.Error.Error())
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	role := model.UserRole{
		ID:     0,
		UserID: mod.ID,
		RoleID: 2,
	}

	tx = mysql.Create(&role)
	if tx.Error != nil {
		logger.Error("插入role失败", tx.Error.Error())
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	key := viper.Get("key")
	info := global2.UserInfo{
		UserId:     int32(mod.ID),
		UserName:   mod.Username,
		Address:    mod.Address,
		Balance:    0,
		Avatar:     mod.Avatar,
		PrivateKey: mod.PrivateKey,
	}
	token, err := global2.GetJwt(key.(string), info)
	logger.Info("生成token", token)
	if err != nil {
		logger.Error("生成token失败", err.Error())
		return &user.Response{Message: err.Error()}, nil
	}
	logger.Info("token生成成功", token)
	set := rds.Set(l.ctx, token, mod.ID, 0)
	if set.Err() != nil {
		//This 问题
		return &user.Response{Message: set.Err().Error()}, nil
	}
	logger.Info("token存储成功", token)
	return &user.Response{
		Message: "注册成功",
		Code:    200,
		Data:    token,
	}, nil
}
