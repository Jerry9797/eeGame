package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta     `path:"/register" tags:"User" method:"post" summary:"注册"`
	Username   string `v:"required"`
	Pwd        string `v:"required"`
	ConfirmPwd string `v:"required"`
	Nickname   string
}
type RegisterRes struct {
	Uid int64
}

type LoginReq struct {
	g.Meta   `path:"/login" tags:"User" method:"post" summary:"登陆"`
	Username string `v:"required"`
	Pwd      string `v:"required"`
}
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	Uid      int
	Username string
	Address  string `json:"address"`
}

type LogoutReq struct {
	g.Meta   `path:"/logout" tags:"User" method:"post" summary:"登陆"`
	Username string `v:"required"`
	Pwd      string `v:"required"`
}
type LogoutRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Msg    string
}
