// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Currency is the golang structure for table currency.
type Currency struct {
	Id        uint64      `json:"id"         description:"自动编号"`
	Status    int         `json:"status"     description:"状态"`
	Code      string      `json:"code"       description:"代码"`
	Name      string      `json:"name"       description:"名称"`
	Symbol    string      `json:"symbol"     description:"符号"`
	Rate      float64     `json:"rate"       description:"汇率"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" description:"删除时间"`
}
