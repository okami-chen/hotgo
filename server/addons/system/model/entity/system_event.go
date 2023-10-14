// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemEvent is the golang structure for table system_event.
type SystemEvent struct {
	Id        uint64      `json:"id"        description:"自动编号"`
	Status    int         `json:"status"    description:"状态"`
	Name      string      `json:"name"      description:"表名"`
	Pk        string      `json:"pk"        description:"主键"`
	Event     string      `json:"event"     description:"事件"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
