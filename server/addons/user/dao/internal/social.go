// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SocialDao is the data access object for table ue_user_social.
type SocialDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SocialColumns // columns contains all the column names of Table for convenient usage.
}

// SocialColumns defines and stores column names for table ue_user_social.
type SocialColumns struct {
	Id        string // 自动编号
	Uid       string // 用户
	Skey      string // 键名
	Svalue    string // 键值
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// socialColumns holds the columns for table ue_user_social.
var socialColumns = SocialColumns{
	Id:        "id",
	Uid:       "uid",
	Skey:      "skey",
	Svalue:    "svalue",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewSocialDao creates and returns a new DAO object for table data access.
func NewSocialDao() *SocialDao {
	return &SocialDao{
		group:   "bee",
		table:   "ue_user_social",
		columns: socialColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SocialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SocialDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SocialDao) Columns() SocialColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SocialDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SocialDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SocialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
