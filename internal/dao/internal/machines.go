// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MachinesDao is the data access object for table xf_machines.
type MachinesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns MachinesColumns // columns contains all the column names of Table for convenient usage.
}

// MachinesColumns defines and stores column names for table xf_machines.
type MachinesColumns struct {
	MachineUuid        string // 机器主键
	MachineName        string // 机器备注名字
	MachineDescription string // 机器描述
	MachineTags        string // 机器标签
	MachineRegion      string //
	HostSystem         string // 机器系统(Linux,Windows,MacOS)
	HostPlatform       string // 机器平台(Ubuntu,Centos)
	HostCpu            string // 机器 CPU
	HostMem            string // 机器运存大小
	HostSwap           string // 机器虚拟内存大小
	HostVirtualization string // 机器虚拟化
	HostBoot           string // 机器启动时间
	IsRegister         string // 是否注册
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
}

// machinesColumns holds the columns for table xf_machines.
var machinesColumns = MachinesColumns{
	MachineUuid:        "machine_uuid",
	MachineName:        "machine_name",
	MachineDescription: "machine_description",
	MachineTags:        "machine_tags",
	MachineRegion:      "machine_region",
	HostSystem:         "host_system",
	HostPlatform:       "host_platform",
	HostCpu:            "host_cpu",
	HostMem:            "host_mem",
	HostSwap:           "host_swap",
	HostVirtualization: "host_virtualization",
	HostBoot:           "host_boot",
	IsRegister:         "is_register",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewMachinesDao creates and returns a new DAO object for table data access.
func NewMachinesDao() *MachinesDao {
	return &MachinesDao{
		group:   "default",
		table:   "xf_machines",
		columns: machinesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MachinesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MachinesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MachinesDao) Columns() MachinesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MachinesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MachinesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MachinesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
