// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Language is the golang structure of table ue_system_language for DAO operations like Where/Data.
type Language struct {
	g.Meta     `orm:"table:ue_system_language, do:true"`
	Id         interface{} // 自动编号
	Language   interface{} // 语言全称-英文
	LanguageZh interface{} // 语言全称-中文
	Code       interface{} // 语言简称
	Image      interface{} // 国旗图片
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
