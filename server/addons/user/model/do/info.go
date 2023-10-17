// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Info is the golang structure of table ue_user_info for DAO operations like Where/Data.
type Info struct {
	g.Meta      `orm:"table:ue_user_info, do:true"`
	Id          interface{} // 自动编号
	Uid         interface{} // 用户
	Username    interface{} // 用户名
	FirstName   interface{} // 用户名
	LastName    interface{} // 用户姓
	Mobile      interface{} // 手机
	Country     interface{} // 国家
	Avatar      interface{} // 头像
	Account     interface{} // 账号
	Gender      interface{} // 性别
	SocialMedia interface{} // 联系
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
