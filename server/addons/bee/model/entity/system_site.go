// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSite is the golang structure for table system_site.
type SystemSite struct {
	Id          uint64      `json:"id"          description:"自动编号"`
	MessageType uint64      `json:"messageType" description:"客服类型"`
	Sort        uint64      `json:"sort"        description:"排序"`
	SiteName    string      `json:"siteName"    description:"网站名称"`
	SiteUrl     string      `json:"siteUrl"     description:"网站网址"`
	SiteLogo    string      `json:"siteLogo"    description:"网站logo"`
	Currencys   string      `json:"currencys"   description:"货币，多个逗号分隔"`
	Languages   string      `json:"languages"   description:"语言，多个逗号分隔"`
	SiteRemark  string      `json:"siteRemark"  description:"网站备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
}
