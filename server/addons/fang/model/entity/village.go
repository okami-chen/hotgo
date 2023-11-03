// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Village is the golang structure for table village.
type Village struct {
	Id        uint64      `json:"id"         description:"编号"`
	Name      string      `json:"name"       description:"名称"`
	Data1     string      `json:"data_1"     description:"数据_1"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" description:"删除时间"`
}
