// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemEvent is the golang structure of table ue_system_event for DAO operations like Where/Data.
type SystemEvent struct {
	g.Meta    `orm:"table:ue_system_event, do:true"`
	Id        interface{} // 自动编号
	Status    interface{} // 状态
	Name      interface{} // 表名
	Pk        interface{} // 主键
	Event     interface{} // 事件
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
