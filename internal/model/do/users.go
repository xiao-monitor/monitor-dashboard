// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table xf_users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:xf_users, do:true"`
	UserUuid  interface{} // 用户主键
	Username  interface{} // 用户名
	Email     interface{} // 用户邮箱
	Password  interface{} // 用户密码
	Role      interface{} // 用户所在角色
	IsEnable  interface{} // 是否启用用户
	IsBan     interface{} // 是否封禁
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
