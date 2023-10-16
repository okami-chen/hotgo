// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Country is the golang structure of table ue_system_country for DAO operations like Where/Data.
type Country struct {
	g.Meta       `orm:"table:ue_system_country, do:true"`
	Id           interface{} // 自动编号
	Country      interface{} // 国家缩写
	CountryName  interface{} // 国家名称全程
	DiallingCode interface{} // 国家区号
	Img          interface{} // 国旗图片
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
