// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Game is the golang structure of table ue_game for DAO operations like Where/Data.
type Game struct {
	g.Meta    `orm:"table:ue_game, do:true"`
	GameId    interface{} // 自动编号
	GameSku   interface{} // 唯一值
	Name      interface{} // 游戏全名
	Title     interface{} // 游戏名称
	Status    interface{} // 状态
	Sort      interface{} // 排序
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
