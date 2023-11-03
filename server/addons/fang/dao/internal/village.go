// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VillageDao is the data access object for table ue_village.
type VillageDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns VillageColumns // columns contains all the column names of Table for convenient usage.
}

// VillageColumns defines and stores column names for table ue_village.
type VillageColumns struct {
	Id        string // 编号
	Name      string // 名称
	Data1     string // 数据_1
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// villageColumns holds the columns for table ue_village.
var villageColumns = VillageColumns{
	Id:        "id",
	Name:      "name",
	Data1:     "data_1",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewVillageDao creates and returns a new DAO object for table data access.
func NewVillageDao() *VillageDao {
	return &VillageDao{
		group:   "wh",
		table:   "ue_village",
		columns: villageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VillageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VillageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VillageDao) Columns() VillageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VillageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VillageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VillageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
