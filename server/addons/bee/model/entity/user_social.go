// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSocial is the golang structure for table user_social.
type UserSocial struct {
	Id        uint64      `json:"id"        description:"自动编号"`
	Uid       uint64      `json:"uid"       description:"用户"`
	Skey      string      `json:"skey"      description:"键名"`
	Svalue    string      `json:"svalue"    description:"键值"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}

