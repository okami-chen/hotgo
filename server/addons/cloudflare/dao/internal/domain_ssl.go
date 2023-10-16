// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DomainSslDao is the data access object for table ue_cf_domain_ssl.
type DomainSslDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns DomainSslColumns // columns contains all the column names of Table for convenient usage.
}

// DomainSslColumns defines and stores column names for table ue_cf_domain_ssl.
type DomainSslColumns struct {
	Id        string // 自动编号
	AccountId string // 关联账号
	Domain    string // 域名
	Status    string // 状态
	DomainId  string // 证书编号
	CertId    string // 证书编号
	CertSubId string //
	Type      string // 类型
	Issuer    string // 签发组织
	Authority string // 授权组织
	Signature string // 签名方式
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	ExpireAt  string // 过期时间
}

// domainSslColumns holds the columns for table ue_cf_domain_ssl.
var domainSslColumns = DomainSslColumns{
	Id:        "id",
	AccountId: "account_id",
	Domain:    "domain",
	Status:    "status",
	DomainId:  "domain_id",
	CertId:    "cert_id",
	CertSubId: "cert_sub_id",
	Type:      "type",
	Issuer:    "issuer",
	Authority: "authority",
	Signature: "signature",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ExpireAt:  "expire_at",
}

// NewDomainSslDao creates and returns a new DAO object for table data access.
func NewDomainSslDao() *DomainSslDao {
	return &DomainSslDao{
		group:   "sync",
		table:   "ue_cf_domain_ssl",
		columns: domainSslColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DomainSslDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DomainSslDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DomainSslDao) Columns() DomainSslColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DomainSslDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DomainSslDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DomainSslDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
