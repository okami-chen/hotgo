// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UeCfDomain is the golang structure for table ue_cf_domain.
type UeCfDomain struct {
	Id        uint64      `json:"id"         description:"自动编号"`
	AccountId uint64      `json:"account_id" description:"关联账号"`
	DomainId  string      `json:"domain_id"  description:"域名编号"`
	Domain    string      `json:"domain"     description:"域名名称"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`
}
