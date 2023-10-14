// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Game is the golang structure for table game.
type Game struct {
	GameId    uint64      `json:"gameId"    description:"自动编号"`
	GameSku   string      `json:"gameSku"   description:"唯一值"`
	Name      string      `json:"name"      description:"游戏全名"`
	Title     string      `json:"title"     description:"游戏名称"`
	Status    uint        `json:"status"    description:"状态"`
	Sort      uint64      `json:"sort"      description:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
