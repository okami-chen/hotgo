// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Currency is the golang structure of table ue_system_currency for DAO operations like Where/Data.
type Currency struct {
	g.Meta    `orm:"table:ue_system_currency, do:true"`
	Id        interface{} // 自动编号
	Status    interface{} // 状态
	Code      interface{} // 代码
	Name      interface{} // 名称
	Symbol    interface{} // 符号
	Rate      interface{} // 汇率
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
