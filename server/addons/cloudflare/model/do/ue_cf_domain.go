// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UeCfDomain is the golang structure of table ue_cf_domain for DAO operations like Where/Data.
type UeCfDomain struct {
	g.Meta    `orm:"table:ue_cf_domain, do:true"`
	Id        interface{} // 自动编号
	AccountId interface{} // 关联账号
	DomainId  interface{} // 域名编号
	Domain    interface{} // 域名名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
