// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Language is the golang structure for table language.
type Language struct {
	Id         uint        `json:"id"          description:"自动编号"`
	Language   string      `json:"language"    description:"语言全称-英文"`
	LanguageZh string      `json:"language_zh" description:"语言全称-中文"`
	Code       string      `json:"code"        description:"语言简称"`
	Image      string      `json:"image"       description:"国旗图片"`
	CreatedAt  *gtime.Time `json:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updated_at"  description:"更新时间"`
	DeletedAt  *gtime.Time `json:"deleted_at"  description:"删除时间"`
}
