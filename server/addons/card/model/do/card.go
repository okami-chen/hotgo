// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Card is the golang structure of table sec_card for DAO operations like Where/Data.
type Card struct {
	g.Meta    `orm:"table:sec_card, do:true"`
	Id        interface{} // 自动编号
	Name      interface{} // 持卡
	Title     interface{} // 名称
	Bank      interface{} // 银行
	Organize  interface{} // 组织
	CardNo    interface{} // 卡号
	ExpireAt  *gtime.Time // 过期时间
	Code      interface{} // 识别码
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
