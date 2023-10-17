// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Social is the golang structure of table ue_user_social for DAO operations like Where/Data.
type Social struct {
	g.Meta    `orm:"table:ue_user_social, do:true"`
	Id        interface{} // 自动编号
	Uid       interface{} // 用户
	Skey      interface{} // 键名
	Svalue    interface{} // 键值
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
