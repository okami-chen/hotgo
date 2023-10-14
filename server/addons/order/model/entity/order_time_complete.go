// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderTimeComplete is the golang structure for table order_time_complete.
type OrderTimeComplete struct {
	Id        uint64      `json:"id"        description:"自动编号"`
	OrderId   uint64      `json:"orderId"   description:"订单编号"`
	AdminId   uint64      `json:"adminId"   description:"操作人"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
