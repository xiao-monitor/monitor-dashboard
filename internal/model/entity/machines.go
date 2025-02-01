// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Machines is the golang structure for table machines.
type Machines struct {
	MachineUuid        string      `json:"machine_uuid"        orm:"machine_uuid"        ` // 机器主键
	MachineName        string      `json:"machine_name"        orm:"machine_name"        ` // 机器备注名字
	MachineDescription string      `json:"machine_description" orm:"machine_description" ` // 机器描述
	MachineTags        string      `json:"machine_tags"        orm:"machine_tags"        ` // 机器标签
	MachineRegion      string      `json:"machine_region"      orm:"machine_region"      ` //
	HostSystem         string      `json:"host_system"         orm:"host_system"         ` // 机器系统(Linux,Windows,MacOS)
	HostPlatform       string      `json:"host_platform"       orm:"host_platform"       ` // 机器平台(Ubuntu,Centos)
	HostCpu            string      `json:"host_cpu"            orm:"host_cpu"            ` // 机器 CPU
	HostMem            uint64      `json:"host_mem"            orm:"host_mem"            ` // 机器运存大小
	HostSwap           uint64      `json:"host_swap"           orm:"host_swap"           ` // 机器虚拟内存大小
	HostVirtualization string      `json:"host_virtualization" orm:"host_virtualization" ` // 机器虚拟化
	HostBoot           *gtime.Time `json:"host_boot"           orm:"host_boot"           ` // 机器启动时间
	IsRegister         int         `json:"is_register"         orm:"is_register"         ` // 是否注册
	CreatedAt          *gtime.Time `json:"created_at"          orm:"created_at"          ` // 创建时间
	UpdatedAt          *gtime.Time `json:"updated_at"          orm:"updated_at"          ` // 修改时间
}
