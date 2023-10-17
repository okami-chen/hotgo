// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AffiliateDao is the data access object for table ue_user_affiliate.
type AffiliateDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AffiliateColumns // columns contains all the column names of Table for convenient usage.
}

// AffiliateColumns defines and stores column names for table ue_user_affiliate.
type AffiliateColumns struct {
	Id               string // 自动编号
	Uid              string // 用户
	IsAffiliate      string // 是否加盟
	AffiliateAudit   string // 加盟审核
	AuditAdmin       string // 审核人员
	AffiliateScale   string // 分成
	AffiliateEmail   string // 邮箱
	AffiliatePayType string // 结算方式
	AffiliateKey     string // 自定义key
	AffiliateTime    string // 加盟时间
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
}

// affiliateColumns holds the columns for table ue_user_affiliate.
var affiliateColumns = AffiliateColumns{
	Id:               "id",
	Uid:              "uid",
	IsAffiliate:      "is_affiliate",
	AffiliateAudit:   "affiliate_audit",
	AuditAdmin:       "audit_admin",
	AffiliateScale:   "affiliate_scale",
	AffiliateEmail:   "affiliate_email",
	AffiliatePayType: "affiliate_pay_type",
	AffiliateKey:     "affiliate_key",
	AffiliateTime:    "affiliate_time",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewAffiliateDao creates and returns a new DAO object for table data access.
func NewAffiliateDao() *AffiliateDao {
	return &AffiliateDao{
		group:   "bee",
		table:   "ue_user_affiliate",
		columns: affiliateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AffiliateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AffiliateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AffiliateDao) Columns() AffiliateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AffiliateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AffiliateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AffiliateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
