// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserSocialDao is the data access object for table ue_user_social.
type UserSocialDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UserSocialColumns // columns contains all the column names of Table for convenient usage.
}

// UserSocialColumns defines and stores column names for table ue_user_social.
type UserSocialColumns struct {
	Id        string // 自动编号
	Uid       string // 用户
	Skey      string // 键名
	Svalue    string // 键值
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// userSocialColumns holds the columns for table ue_user_social.
var userSocialColumns = UserSocialColumns{
	Id:        "id",
	Uid:       "uid",
	Skey:      "skey",
	Svalue:    "svalue",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewUserSocialDao creates and returns a new DAO object for table data access.
func NewUserSocialDao() *UserSocialDao {
	return &UserSocialDao{
		group:   "bee",
		table:   "ue_user_social",
		columns: userSocialColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserSocialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserSocialDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserSocialDao) Columns() UserSocialColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserSocialDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserSocialDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserSocialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
