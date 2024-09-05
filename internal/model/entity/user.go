// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Uid      int         `json:"uid"      orm:"uid"       ` // 用户ID
	Username string      `json:"username" orm:"username"  ` // 用户名
	Address  string      `json:"address"  orm:"address"	  ` // 账户
	Balance  int64       `json:"balance"  orm:"balance"   ` // 金币
	Nickname string      `json:"nickname" orm:"nickname"  ` // 昵称
	Avatar   string      `json:"avatar"   orm:"avatar"    ` // 头像
	Pwd      string      `json:"pwd"      orm:"pwd"       ` // 密码
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` // 创建时间
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` // 修改时间
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` // 删除时间
}
