// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure of table xf_roles for DAO operations like Where/Data.
type Roles struct {
	g.Meta    `orm:"table:xf_roles, do:true"`
	RoleUuid  interface{} // 角色 UUID
	RoleName  interface{} // 角色名
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
