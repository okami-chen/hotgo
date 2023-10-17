// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Info is the golang structure for table info.
type Info struct {
	Id          uint64      `json:"id"           description:"自动编号"`
	Uid         uint64      `json:"uid"          description:"用户"`
	Username    string      `json:"username"     description:"用户名"`
	FirstName   string      `json:"first_name"   description:"用户名"`
	LastName    string      `json:"last_name"    description:"用户姓"`
	Mobile      string      `json:"mobile"       description:"手机"`
	Country     string      `json:"country"      description:"国家"`
	Avatar      string      `json:"avatar"       description:"头像"`
	Account     string      `json:"account"      description:"账号"`
	Gender      string      `json:"gender"       description:"性别"`
	SocialMedia string      `json:"social_media" description:"联系"`
	CreatedAt   *gtime.Time `json:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deleted_at"   description:"删除时间"`
}
