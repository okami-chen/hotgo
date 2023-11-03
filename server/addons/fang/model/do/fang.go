// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Fang is the golang structure of table ue_fang for DAO operations like Where/Data.
type Fang struct {
	g.Meta      `orm:"table:ue_fang, do:true"`
	Id          interface{} // 编号
	Sid         interface{} // 编号
	Village     interface{} // 小区
	Title       interface{} // 标题
	Price       interface{} // 价格
	HouseType   interface{} // 户型
	Area        interface{} // 面积
	ToWard      interface{} // 朝向
	Lift        interface{} // 电梯
	Water       interface{} // 用水
	Electricity interface{} // 用电
	Tenancy     interface{} // 租期
	Verify      interface{} // 核验
	CheckIn     interface{} // 入住
	Flag        interface{} // 旗帜
	Url         interface{} // 网址
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
