// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	UserUuid  string      `json:"user_uuid"  orm:"user_uuid"  ` // 用户主键
	Username  string      `json:"username"   orm:"username"   ` // 用户名
	Email     string      `json:"email"      orm:"email"      ` // 用户邮箱
	Password  string      `json:"password"   orm:"password"   ` // 用户密码
	Role      string      `json:"role"       orm:"role"       ` // 用户所在角色
	IsEnable  int         `json:"is_enable"  orm:"is_enable"  ` // 是否启用用户
	IsBan     int         `json:"is_ban"     orm:"is_ban"     ` // 是否封禁
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新时间
}
