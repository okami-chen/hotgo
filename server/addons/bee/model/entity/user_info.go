// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id          uint64      `json:"id"          description:"自动编号"`
	Uid         uint64      `json:"uid"         description:"用户"`
	Username    string      `json:"username"    description:"用户名"`
	FirstName   string      `json:"firstName"   description:"用户名"`
	LastName    string      `json:"lastName"    description:"用户姓"`
	Mobile      string      `json:"mobile"      description:"手机"`
	Country     string      `json:"country"     description:"国家"`
	Avatar      string      `json:"avatar"      description:"头像"`
	Account     string      `json:"account"     description:"账号"`
	Gender      string      `json:"gender"      description:"性别"`
	SocialMedia string      `json:"socialMedia" description:"联系"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
}
