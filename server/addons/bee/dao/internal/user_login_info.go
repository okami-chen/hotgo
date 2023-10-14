// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginInfoDao is the data access object for table ue_user_login_info.
type UserLoginInfoDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns UserLoginInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserLoginInfoColumns defines and stores column names for table ue_user_login_info.
type UserLoginInfoColumns struct {
	Id            string // 自动编号
	Uid           string // 用户
	LastLoginIp   string // 最后IP
	RegisterIp    string // 注册IP
	RegisterTime  string // 最后登录
	LastLoginTime string // 最后登录
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// userLoginInfoColumns holds the columns for table ue_user_login_info.
var userLoginInfoColumns = UserLoginInfoColumns{
	Id:            "id",
	Uid:           "uid",
	LastLoginIp:   "last_login_ip",
	RegisterIp:    "register_ip",
	RegisterTime:  "register_time",
	LastLoginTime: "last_login_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewUserLoginInfoDao creates and returns a new DAO object for table data access.
func NewUserLoginInfoDao() *UserLoginInfoDao {
	return &UserLoginInfoDao{
		group:   "bee",
		table:   "ue_user_login_info",
		columns: userLoginInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserLoginInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserLoginInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserLoginInfoDao) Columns() UserLoginInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserLoginInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserLoginInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserLoginInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
