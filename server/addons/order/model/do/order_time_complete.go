// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderTimeComplete is the golang structure of table ue_order_time_complete for DAO operations like Where/Data.
type OrderTimeComplete struct {
	g.Meta    `orm:"table:ue_order_time_complete, do:true"`
	Id        interface{} // 自动编号
	OrderId   interface{} // 订单编号
	AdminId   interface{} // 操作人
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
