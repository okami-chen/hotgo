// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table ue_user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table ue_user.
type UserColumns struct {
	Uid            string // 自动编号
	SiteId         string // 站点编号
	ShareUid       string // 分享用户
	Email          string // 电子邮箱
	Password       string // 密码
	Grade          string // 等级
	Status         string // 状态
	ConsumeMoney   string // 消费金额
	NonMoney       string // 非货币
	FreezeMoney    string // 冻结金额
	Amount         string // 账户余额
	Payment        string // 支付方式
	DefaultPayment string // 默认支付方式
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// userColumns holds the columns for table ue_user.
var userColumns = UserColumns{
	Uid:            "uid",
	SiteId:         "site_id",
	ShareUid:       "share_uid",
	Email:          "email",
	Password:       "password",
	Grade:          "grade",
	Status:         "status",
	ConsumeMoney:   "consume_money",
	NonMoney:       "non_money",
	FreezeMoney:    "freeze_money",
	Amount:         "amount",
	Payment:        "payment",
	DefaultPayment: "default_payment",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "bee",
		table:   "ue_user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
