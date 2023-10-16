// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CurrencyDao is the data access object for table ue_system_currency.
type CurrencyDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CurrencyColumns // columns contains all the column names of Table for convenient usage.
}

// CurrencyColumns defines and stores column names for table ue_system_currency.
type CurrencyColumns struct {
	Id        string // 自动编号
	Status    string // 状态
	Code      string // 代码
	Name      string // 名称
	Symbol    string // 符号
	Rate      string // 汇率
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// currencyColumns holds the columns for table ue_system_currency.
var currencyColumns = CurrencyColumns{
	Id:        "id",
	Status:    "status",
	Code:      "code",
	Name:      "name",
	Symbol:    "symbol",
	Rate:      "rate",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCurrencyDao creates and returns a new DAO object for table data access.
func NewCurrencyDao() *CurrencyDao {
	return &CurrencyDao{
		group:   "bee",
		table:   "ue_system_currency",
		columns: currencyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CurrencyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CurrencyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CurrencyDao) Columns() CurrencyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CurrencyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CurrencyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CurrencyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
