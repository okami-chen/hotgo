// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Affiliate is the golang structure for table affiliate.
type Affiliate struct {
	Id               uint64      `json:"id"                 description:"自动编号"`
	Uid              uint64      `json:"uid"                description:"用户"`
	IsAffiliate      int         `json:"is_affiliate"       description:"是否加盟"`
	AffiliateAudit   int         `json:"affiliate_audit"    description:"加盟审核"`
	AuditAdmin       uint64      `json:"audit_admin"        description:"审核人员"`
	AffiliateScale   uint        `json:"affiliate_scale"    description:"分成"`
	AffiliateEmail   string      `json:"affiliate_email"    description:"邮箱"`
	AffiliatePayType uint        `json:"affiliate_pay_type" description:"结算方式"`
	AffiliateKey     string      `json:"affiliate_key"      description:"自定义key"`
	AffiliateTime    *gtime.Time `json:"affiliate_time"     description:"加盟时间"`
	CreatedAt        *gtime.Time `json:"created_at"         description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updated_at"         description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deleted_at"         description:"删除时间"`
}
