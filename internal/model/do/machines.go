// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Machines is the golang structure of table xf_machines for DAO operations like Where/Data.
type Machines struct {
	g.Meta             `orm:"table:xf_machines, do:true"`
	MachineUuid        interface{} // 机器主键
	MachineName        interface{} // 机器备注名字
	MachineDescription interface{} // 机器描述
	MachineTags        interface{} // 机器标签
	MachineRegion      interface{} //
	HostSystem         interface{} // 机器系统(Linux,Windows,MacOS)
	HostPlatform       interface{} // 机器平台(Ubuntu,Centos)
	HostCpu            interface{} // 机器 CPU
	HostMem            interface{} // 机器运存大小
	HostSwap           interface{} // 机器虚拟内存大小
	HostVirtualization interface{} // 机器虚拟化
	HostBoot           *gtime.Time // 机器启动时间
	IsRegister         interface{} // 是否注册
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 修改时间
}
