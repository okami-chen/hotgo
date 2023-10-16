// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DomainSsl is the golang structure of table ue_cf_domain_ssl for DAO operations like Where/Data.
type DomainSsl struct {
	g.Meta    `orm:"table:ue_cf_domain_ssl, do:true"`
	Id        interface{} // 自动编号
	AccountId interface{} // 关联账号
	Domain    interface{} // 域名
	Status    interface{} // 状态
	DomainId  interface{} // 证书编号
	CertId    interface{} // 证书编号
	CertSubId interface{} //
	Type      interface{} // 类型
	Issuer    interface{} // 签发组织
	Authority interface{} // 授权组织
	Signature interface{} // 签名方式
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	ExpireAt  *gtime.Time // 过期时间
}
