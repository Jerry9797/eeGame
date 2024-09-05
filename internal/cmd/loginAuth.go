package cmd

import (
	"context"
	v1 "eeGame/api/user"
	"eeGame/internal/consts"
	"eeGame/internal/dao"
	"eeGame/internal/model/entity"
	"eeGame/internal/service"
	"eeGame/utility/response"
	"encoding/json"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:       consts.CacheModeRedis,
		ServerName:      consts.BackendServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFunc,
		LoginAfterFunc:  loginAfterFunc,
		LogoutPath:      "/user/logout",
		AuthPaths:       g.SliceStr{"/api"},
		//AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		AuthAfterFunc: authAfterFunc,
		MultiLogin:    consts.MultiLogin,
	}
	err = gfAdminToken.Start()
	return
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	pwd := r.Get("pwd").String()
	ctx := context.TODO()

	if username == "" || pwd == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}

	//验证账号密码是否正确
	req := v1.LoginReq{
		Username: username,
		Pwd:      pwd,
	}
	user, err := service.User().SignIn(ctx, &req)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
	}
	return consts.GTokenAdminPrefix + username, user
}

// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		userKey := respData.GetString("userKey")
		username := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		user, err := dao.User.GetUserByUsername(r.Context(), username)
		if err != nil {
			panic(err)
		}
		data := &v1.LoginRes{
			Type:  consts.TokenType,
			Token: respData.GetString("token"),
			//todo 通过角色查询权限
		}
		data.Uid = user.Uid
		data.Username = user.Username
		data.Address = user.Address
		response.JsonExit(r, 0, "登陆成功", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.User
	getString := respData.Data
	if getString == "" {
		response.Json(r, 401, "用户未登录", nil)
		return
	}
	marshal, err := json.Marshal(getString)
	if err != nil {
		return
	}
	var result map[string]interface{}
	// 将 JSON 字符串解析到 map 中
	if err := json.Unmarshal(marshal, &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if err := gconv.Struct(result["data"], &userInfo); err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxUserId, userInfo.Uid)
	r.SetCtxVar(consts.CtxUserName, userInfo.Username)
	r.SetCtxVar(consts.CtxAddress, userInfo.Address)
	r.Middleware.Next()
}
