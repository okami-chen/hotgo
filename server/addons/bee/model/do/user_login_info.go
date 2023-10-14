// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginInfo is the golang structure of table ue_user_login_info for DAO operations like Where/Data.
type UserLoginInfo struct {
	g.Meta        `orm:"table:ue_user_login_info, do:true"`
	Id            interface{} // 自动编号
	Uid           interface{} // 用户
	LastLoginIp   interface{} // 最后IP
	RegisterIp    interface{} // 注册IP
	RegisterTime  *gtime.Time // 最后登录
	LastLoginTime *gtime.Time // 最后登录
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
