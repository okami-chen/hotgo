// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table ue_user for DAO operations like Where/Data.
type User struct {
	g.Meta         `orm:"table:ue_user, do:true"`
	Uid            interface{} // 自动编号
	SiteId         interface{} // 站点编号
	ShareUid       interface{} // 分享用户
	Email          interface{} // 电子邮箱
	Password       interface{} // 密码
	Grade          interface{} // 等级
	Status         interface{} // 状态
	ConsumeMoney   interface{} // 消费金额
	NonMoney       interface{} // 非货币
	FreezeMoney    interface{} // 冻结金额
	Amount         interface{} // 账户余额
	Payment        interface{} // 支付方式
	DefaultPayment interface{} // 默认支付方式
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
