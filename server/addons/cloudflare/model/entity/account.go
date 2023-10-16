// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure for table account.
type Account struct {
	Id        uint64      `json:"id"         description:"自动编号"`
	Email     string      `json:"email"      description:"email"`
	Token     string      `json:"token"      description:"令牌"`
	Remark    string      `json:"remark"     description:"备注"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`
}
