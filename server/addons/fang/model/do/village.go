// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Village is the golang structure of table ue_village for DAO operations like Where/Data.
type Village struct {
	g.Meta    `orm:"table:ue_village, do:true"`
	Id        interface{} // 编号
	Name      interface{} // 名称
	Data1     interface{} // 数据_1
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
