// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CfAccountDao is the data access object for table ue_cf_account.
type CfAccountDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns CfAccountColumns // columns contains all the column names of Table for convenient usage.
}

// CfAccountColumns defines and stores column names for table ue_cf_account.
type CfAccountColumns struct {
	Id        string // 自动编号
	Email     string // email
	Token     string // 令牌
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// cfAccountColumns holds the columns for table ue_cf_account.
var cfAccountColumns = CfAccountColumns{
	Id:        "id",
	Email:     "email",
	Token:     "token",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCfAccountDao creates and returns a new DAO object for table data access.
func NewCfAccountDao() *CfAccountDao {
	return &CfAccountDao{
		group:   "sync",
		table:   "ue_cf_account",
		columns: cfAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CfAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CfAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CfAccountDao) Columns() CfAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CfAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CfAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CfAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
