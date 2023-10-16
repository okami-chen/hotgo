// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LanguageDao is the data access object for table ue_system_language.
type LanguageDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LanguageColumns // columns contains all the column names of Table for convenient usage.
}

// LanguageColumns defines and stores column names for table ue_system_language.
type LanguageColumns struct {
	Id         string // 自动编号
	Language   string // 语言全称-英文
	LanguageZh string // 语言全称-中文
	Code       string // 语言简称
	Image      string // 国旗图片
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// languageColumns holds the columns for table ue_system_language.
var languageColumns = LanguageColumns{
	Id:         "id",
	Language:   "language",
	LanguageZh: "language_zh",
	Code:       "code",
	Image:      "image",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewLanguageDao creates and returns a new DAO object for table data access.
func NewLanguageDao() *LanguageDao {
	return &LanguageDao{
		group:   "bee",
		table:   "ue_system_language",
		columns: languageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LanguageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LanguageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LanguageDao) Columns() LanguageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LanguageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LanguageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LanguageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
