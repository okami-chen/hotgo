// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Uid            uint64      `json:"uid"             description:"自动编号"`
	SiteId         uint64      `json:"site_id"         description:"站点编号"`
	ShareUid       uint64      `json:"share_uid"       description:"分享用户"`
	Email          string      `json:"email"           description:"电子邮箱"`
	Password       string      `json:"password"        description:"密码"`
	Grade          uint        `json:"grade"           description:"等级"`
	Status         int         `json:"status"          description:"状态"`
	ConsumeMoney   float64     `json:"consume_money"   description:"消费金额"`
	NonMoney       float64     `json:"non_money"       description:"非货币"`
	FreezeMoney    float64     `json:"freeze_money"    description:"冻结金额"`
	Amount         float64     `json:"amount"          description:"账户余额"`
	Payment        string      `json:"payment"         description:"支付方式"`
	DefaultPayment string      `json:"default_payment" description:"默认支付方式"`
	CreatedAt      *gtime.Time `json:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updated_at"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deleted_at"      description:"删除时间"`
}
