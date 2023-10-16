// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure of table ue_cf_account for DAO operations like Where/Data.
type Account struct {
	g.Meta    `orm:"table:ue_cf_account, do:true"`
	Id        interface{} // 自动编号
	Email     interface{} // email
	Token     interface{} // 令牌
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
