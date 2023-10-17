// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Affiliate is the golang structure of table ue_user_affiliate for DAO operations like Where/Data.
type Affiliate struct {
	g.Meta           `orm:"table:ue_user_affiliate, do:true"`
	Id               interface{} // 自动编号
	Uid              interface{} // 用户
	IsAffiliate      interface{} // 是否加盟
	AffiliateAudit   interface{} // 加盟审核
	AuditAdmin       interface{} // 审核人员
	AffiliateScale   interface{} // 分成
	AffiliateEmail   interface{} // 邮箱
	AffiliatePayType interface{} // 结算方式
	AffiliateKey     interface{} // 自定义key
	AffiliateTime    *gtime.Time // 加盟时间
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
}
