// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CardDao is the data access object for table sec_card.
type CardDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns CardColumns // columns contains all the column names of Table for convenient usage.
}

// CardColumns defines and stores column names for table sec_card.
type CardColumns struct {
	Id        string // 自动编号
	Name      string // 持卡
	Title     string // 名称
	Bank      string // 银行
	Organize  string // 组织
	CardNo    string // 卡号
	ExpireAt  string // 过期时间
	Code      string // 识别码
	Remark    string // 备注
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// cardColumns holds the columns for table sec_card.
var cardColumns = CardColumns{
	Id:        "id",
	Name:      "name",
	Title:     "title",
	Bank:      "bank",
	Organize:  "organize",
	CardNo:    "card_no",
	ExpireAt:  "expire_at",
	Code:      "code",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCardDao creates and returns a new DAO object for table data access.
func NewCardDao() *CardDao {
	return &CardDao{
		group:   "sec",
		table:   "sec_card",
		columns: cardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CardDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CardDao) Columns() CardColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CardDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CardDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
