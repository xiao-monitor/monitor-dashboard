// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	RoleUuid  string      `json:"role_uuid"  orm:"role_uuid"  ` // 角色 UUID
	RoleName  string      `json:"role_name"  orm:"role_name"  ` // 角色名
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新时间
}
