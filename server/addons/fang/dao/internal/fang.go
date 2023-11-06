// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FangDao is the data access object for table ue_fang.
type FangDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns FangColumns // columns contains all the column names of Table for convenient usage.
}

// FangColumns defines and stores column names for table ue_fang.
type FangColumns struct {
	Id          string // 编号
	Sid         string // 编号
	Village     string // 小区
	Title       string // 标题
	Price       string // 价格
	HouseType   string // 户型
	Area        string // 面积
	ToWard      string // 朝向
	Lift        string // 电梯
	Water       string // 用水
	Electricity string // 用电
	Tenancy     string // 租期
	Verify      string // 核验
	CheckIn     string // 入住
	Flag        string // 旗帜
	Url         string // 网址
	Remark      string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// fangColumns holds the columns for table ue_fang.
var fangColumns = FangColumns{
	Id:          "id",
	Sid:         "sid",
	Village:     "village",
	Title:       "title",
	Price:       "price",
	HouseType:   "house_type",
	Area:        "area",
	ToWard:      "to_ward",
	Lift:        "lift",
	Water:       "water",
	Electricity: "electricity",
	Tenancy:     "tenancy",
	Verify:      "verify",
	CheckIn:     "check_in",
	Flag:        "flag",
	Url:         "url",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewFangDao creates and returns a new DAO object for table data access.
func NewFangDao() *FangDao {
	return &FangDao{
		group:   "wh",
		table:   "ue_fang",
		columns: fangColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FangDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FangDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FangDao) Columns() FangColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FangDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FangDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FangDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
