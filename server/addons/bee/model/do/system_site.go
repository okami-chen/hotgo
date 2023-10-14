// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSite is the golang structure of table ue_system_site for DAO operations like Where/Data.
type SystemSite struct {
	g.Meta      `orm:"table:ue_system_site, do:true"`
	Id          interface{} // 自动编号
	MessageType interface{} // 客服类型
	Sort        interface{} // 排序
	SiteName    interface{} // 网站名称
	SiteUrl     interface{} // 网站网址
	SiteLogo    interface{} // 网站logo
	Currencys   interface{} // 货币，多个逗号分隔
	Languages   interface{} // 语言，多个逗号分隔
	SiteRemark  interface{} // 网站备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
