// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure of table ue_device for DAO operations like Where/Data.
type Device struct {
	g.Meta    `orm:"table:ue_device, do:true"`
	DeviceId  interface{} // 自动编号
	Title     interface{} // 分类名称
	Image     interface{} // 图片
	Status    interface{} // 状态
	Sort      interface{} // 排序
	Url       interface{} // 网址
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
