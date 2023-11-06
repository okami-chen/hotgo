// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Fang is the golang structure for table fang.
type Fang struct {
	Id          uint64      `json:"id"          description:"编号"`
	Sid         string      `json:"sid"         description:"编号"`
	Village     string      `json:"village"     description:"小区"`
	Title       string      `json:"title"       description:"标题"`
	Price       uint64      `json:"price"       description:"价格"`
	HouseType   string      `json:"house_type"  description:"户型"`
	Area        string      `json:"area"        description:"面积"`
	ToWard      string      `json:"to_ward"     description:"朝向"`
	Lift        string      `json:"lift"        description:"电梯"`
	Water       string      `json:"water"       description:"用水"`
	Electricity string      `json:"electricity" description:"用电"`
	Tenancy     string      `json:"tenancy"     description:"租期"`
	Verify      string      `json:"verify"      description:"核验"`
	CheckIn     string      `json:"check_in"    description:"入住"`
	Flag        uint        `json:"flag"        description:"旗帜"`
	Url         string      `json:"url"         description:"网址"`
	Remark      string      `json:"remark"      description:"备注"`
	CreatedAt   *gtime.Time `json:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at"  description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deleted_at"  description:"删除时间"`
}
