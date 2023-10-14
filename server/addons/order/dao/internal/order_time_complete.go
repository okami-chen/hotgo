// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderTimeCompleteDao is the data access object for table ue_order_time_complete.
type OrderTimeCompleteDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns OrderTimeCompleteColumns // columns contains all the column names of Table for convenient usage.
}

// OrderTimeCompleteColumns defines and stores column names for table ue_order_time_complete.
type OrderTimeCompleteColumns struct {
	Id        string // 自动编号
	OrderId   string // 订单编号
	AdminId   string // 操作人
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// orderTimeCompleteColumns holds the columns for table ue_order_time_complete.
var orderTimeCompleteColumns = OrderTimeCompleteColumns{
	Id:        "id",
	OrderId:   "order_id",
	AdminId:   "admin_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewOrderTimeCompleteDao creates and returns a new DAO object for table data access.
func NewOrderTimeCompleteDao() *OrderTimeCompleteDao {
	return &OrderTimeCompleteDao{
		group:   "bee",
		table:   "ue_order_time_complete",
		columns: orderTimeCompleteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderTimeCompleteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderTimeCompleteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderTimeCompleteDao) Columns() OrderTimeCompleteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderTimeCompleteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderTimeCompleteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderTimeCompleteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
