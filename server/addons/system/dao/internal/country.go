// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CountryDao is the data access object for table ue_system_country.
type CountryDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CountryColumns // columns contains all the column names of Table for convenient usage.
}

// CountryColumns defines and stores column names for table ue_system_country.
type CountryColumns struct {
	Id           string // 自动编号
	Country      string // 国家缩写
	CountryName  string // 国家名称全程
	DiallingCode string // 国家区号
	Img          string // 国旗图片
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// countryColumns holds the columns for table ue_system_country.
var countryColumns = CountryColumns{
	Id:           "id",
	Country:      "country",
	CountryName:  "country_name",
	DiallingCode: "dialling_code",
	Img:          "img",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewCountryDao creates and returns a new DAO object for table data access.
func NewCountryDao() *CountryDao {
	return &CountryDao{
		group:   "bee",
		table:   "ue_system_country",
		columns: countryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CountryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CountryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CountryDao) Columns() CountryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CountryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CountryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CountryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
