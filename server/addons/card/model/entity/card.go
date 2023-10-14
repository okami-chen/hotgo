// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Card is the golang structure for table card.
type Card struct {
	Id        uint64      `json:"id"        description:"自动编号"`
	Name      string      `json:"name"      description:"持卡"`
	Title     string      `json:"title"     description:"名称"`
	Bank      string      `json:"bank"      description:"银行"`
	Organize  string      `json:"organize"  description:"组织"`
	CardNo    string      `json:"cardNo"    description:"卡号"`
	ExpireAt  *gtime.Time `json:"expireAt"  description:"过期时间"`
	Code      string      `json:"code"      description:"识别码"`
	Remark    string      `json:"remark"    description:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
