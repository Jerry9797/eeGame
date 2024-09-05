// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Uid      interface{} // 用户ID
	Username interface{} // 用户名
	Balance  interface{} // 金币
	Nickname interface{} // 昵称
	Avatar   interface{} // 头像
	Pwd      interface{} // 密码
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 修改时间
	DeleteAt *gtime.Time // 删除时间
}
