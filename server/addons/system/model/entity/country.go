// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Country is the golang structure for table country.
type Country struct {
	Id           uint64      `json:"id"            description:"自动编号"`
	Country      string      `json:"country"       description:"国家缩写"`
	CountryName  string      `json:"country_name"  description:"国家名称全程"`
	DiallingCode string      `json:"dialling_code" description:"国家区号"`
	Img          string      `json:"img"           description:"国旗图片"`
	CreatedAt    *gtime.Time `json:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updated_at"    description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deleted_at"    description:"删除时间"`
}
