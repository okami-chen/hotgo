// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAffiliate is the golang structure for table user_affiliate.
type UserAffiliate struct {
	Id               uint64      `json:"id"               description:"自动编号"`
	Uid              uint64      `json:"uid"              description:"用户"`
	IsAffiliate      int         `json:"isAffiliate"      description:"是否加盟"`
	AffiliateAudit   int         `json:"affiliateAudit"   description:"加盟审核"`
	AuditAdmin       uint64      `json:"auditAdmin"       description:"审核人员"`
	AffiliateScale   uint        `json:"affiliateScale"   description:"分成"`
	AffiliateEmail   string      `json:"affiliateEmail"   description:"邮箱"`
	AffiliatePayType uint        `json:"affiliatePayType" description:"结算方式"`
	AffiliateKey     string      `json:"affiliateKey"     description:"自定义key"`
	AffiliateTime    *gtime.Time `json:"affiliateTime"    description:"加盟时间"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        description:"删除时间"`
}
