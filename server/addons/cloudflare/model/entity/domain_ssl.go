// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DomainSsl is the golang structure for table domain_ssl.
type DomainSsl struct {
	Id        uint64      `json:"id"          description:"自动编号"`
	AccountId uint64      `json:"account_id"  description:"关联账号"`
	Domain    string      `json:"domain"      description:"域名"`
	Status    string      `json:"status"      description:"状态"`
	DomainId  string      `json:"domain_id"   description:"证书编号"`
	CertId    string      `json:"cert_id"     description:"证书编号"`
	CertSubId string      `json:"cert_sub_id" description:""`
	Type      string      `json:"type"        description:"类型"`
	Issuer    string      `json:"issuer"      description:"签发组织"`
	Authority string      `json:"authority"   description:"授权组织"`
	Signature string      `json:"signature"   description:"签名方式"`
	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at"  description:"更新时间"`
	ExpireAt  *gtime.Time `json:"expire_at"   description:"过期时间"`
}
