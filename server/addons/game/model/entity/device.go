// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	DeviceId  uint64      `json:"deviceId"  description:"自动编号"`
	Title     string      `json:"title"     description:"分类名称"`
	Image     string      `json:"image"     description:"图片"`
	Status    uint        `json:"status"    description:"状态"`
	Sort      uint64      `json:"sort"      description:"排序"`
	Url       string      `json:"url"       description:"网址"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
