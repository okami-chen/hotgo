// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Social is the golang structure for table social.
type Social struct {
	Id        uint64      `json:"id"         description:"自动编号"`
	Uid       uint64      `json:"uid"        description:"用户"`
	Skey      string      `json:"skey"       description:"键名"`
	Svalue    string      `json:"svalue"     description:"键值"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" description:"删除时间"`
}
