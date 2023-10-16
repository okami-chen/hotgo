// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UeCfDomainDao is the data access object for table ue_cf_domain.
type UeCfDomainDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UeCfDomainColumns // columns contains all the column names of Table for convenient usage.
}

// UeCfDomainColumns defines and stores column names for table ue_cf_domain.
type UeCfDomainColumns struct {
	Id        string // 自动编号
	AccountId string // 关联账号
	DomainId  string // 域名编号
	Domain    string // 域名名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// ueCfDomainColumns holds the columns for table ue_cf_domain.
var ueCfDomainColumns = UeCfDomainColumns{
	Id:        "id",
	AccountId: "account_id",
	DomainId:  "domain_id",
	Domain:    "domain",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUeCfDomainDao creates and returns a new DAO object for table data access.
func NewUeCfDomainDao() *UeCfDomainDao {
	return &UeCfDomainDao{
		group:   "sync",
		table:   "ue_cf_domain",
		columns: ueCfDomainColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UeCfDomainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UeCfDomainDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UeCfDomainDao) Columns() UeCfDomainColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UeCfDomainDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UeCfDomainDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UeCfDomainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
