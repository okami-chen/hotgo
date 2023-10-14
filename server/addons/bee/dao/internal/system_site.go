// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemSiteDao is the data access object for table ue_system_site.
type SystemSiteDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SystemSiteColumns // columns contains all the column names of Table for convenient usage.
}

// SystemSiteColumns defines and stores column names for table ue_system_site.
type SystemSiteColumns struct {
	Id          string // 自动编号
	MessageType string // 客服类型
	Sort        string // 排序
	SiteName    string // 网站名称
	SiteUrl     string // 网站网址
	SiteLogo    string // 网站logo
	Currencys   string // 货币，多个逗号分隔
	Languages   string // 语言，多个逗号分隔
	SiteRemark  string // 网站备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// systemSiteColumns holds the columns for table ue_system_site.
var systemSiteColumns = SystemSiteColumns{
	Id:          "id",
	MessageType: "message_type",
	Sort:        "sort",
	SiteName:    "site_name",
	SiteUrl:     "site_url",
	SiteLogo:    "site_logo",
	Currencys:   "currencys",
	Languages:   "languages",
	SiteRemark:  "site_remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSystemSiteDao creates and returns a new DAO object for table data access.
func NewSystemSiteDao() *SystemSiteDao {
	return &SystemSiteDao{
		group:   "bee",
		table:   "ue_system_site",
		columns: systemSiteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemSiteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemSiteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemSiteDao) Columns() SystemSiteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemSiteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemSiteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemSiteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
