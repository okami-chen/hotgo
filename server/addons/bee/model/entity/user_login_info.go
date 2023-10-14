// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginInfo is the golang structure for table user_login_info.
type UserLoginInfo struct {
	Id            uint64      `json:"id"            description:"自动编号"`
	Uid           uint64      `json:"uid"           description:"用户"`
	LastLoginIp   string      `json:"lastLoginIp"   description:"最后IP"`
	RegisterIp    string      `json:"registerIp"    description:"注册IP"`
	RegisterTime  *gtime.Time `json:"registerTime"  description:"最后登录"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" description:"最后登录"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
}
