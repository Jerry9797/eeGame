package user

import (
	"context"
	v1 "eeGame/api/user"
	"eeGame/internal/dao"
	"eeGame/internal/model/entity"
	"eeGame/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) Register(ctx context.Context, req *v1.RegisterReq) (int64, error) {
	if req.Pwd != req.ConfirmPwd {
		g.RequestFromCtx(ctx).Response.Writeln("pwd error")
	}
	pwdEncrypt, _ := gmd5.Encrypt(req.Pwd)
	user := entity.User{
		Username: req.Username,
		Pwd:      pwdEncrypt,
		Balance:  0,
		Nickname: req.Nickname,
	}
	return dao.User.Register(ctx, user)
}

func (s *sUser) SignIn(ctx context.Context, req *v1.LoginReq) (*entity.User, error) {
	user, err := dao.User.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return user, fmt.Errorf("查询用户错误： %v", err)
	}
	pwdEncrypt, _ := gmd5.Encrypt(req.Pwd)
	if pwdEncrypt != user.Pwd {
		return user, fmt.Errorf("密码错误")
	}
	return user, nil
}
