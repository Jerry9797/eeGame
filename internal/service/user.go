package service

import (
	"context"
	v1 "eeGame/api/user"
	"eeGame/internal/model/entity"
)

type IUser interface {
	Register(ctx context.Context, req *v1.RegisterReq) (int64, error)
	SignIn(ctx context.Context, req *v1.LoginReq) (*entity.User, error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("IUser 接口未注册")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
