package user

import (
	"context"
	v1 "eeGame/api/user"
	"eeGame/internal/service"
)

type User struct {
}

func New() *User {
	return &User{}
}

func (c *User) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	uid, err := service.User().Register(ctx, req)
	if err != nil {
		return res, err
	}
	res = new(v1.RegisterRes)
	res.Uid = uid
	return res, nil
}
