// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sysin

import (
	"context"
	"hotgo/addons/game/model/entity"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameUpdateFields 修改游戏字段过滤
type GameUpdateFields struct {
	GameSku string `json:"gameSku" dc:"唯一值"`
	Name    string `json:"name"    dc:"游戏全名"`
	Title   string `json:"title"   dc:"游戏名称"`
	Status  int    `json:"status"  dc:"状态"`
	Sort    int64  `json:"sort"    dc:"排序"`
}

// GameInsertFields 新增游戏字段过滤
type GameInsertFields struct {
	GameSku string `json:"gameSku" dc:"唯一值"`
	Name    string `json:"name"    dc:"游戏全名"`
	Title   string `json:"title"   dc:"游戏名称"`
	Status  int    `json:"status"  dc:"状态"`
	Sort    int64  `json:"sort"    dc:"排序"`
}

// GameEditInp 修改/新增游戏
type GameEditInp struct {
	entity.Game
}

func (in *GameEditInp) Filter(ctx context.Context) (err error) {
	// 验证唯一值
	if err := g.Validator().Rules("required").Data(in.GameSku).Messages("唯一值不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证游戏全名
	if err := g.Validator().Rules("required").Data(in.Name).Messages("游戏全名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证游戏名称
	if err := g.Validator().Rules("required").Data(in.Title).Messages("游戏名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type GameEditModel struct{}

// GameDeleteInp 删除游戏
type GameDeleteInp struct {
	GameId interface{} `json:"gameId" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *GameDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type GameDeleteModel struct{}

// GameViewInp 获取指定游戏信息
type GameViewInp struct {
	GameId int64 `json:"gameId" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *GameViewInp) Filter(ctx context.Context) (err error) {
	return
}

type GameViewModel struct {
	entity.Game
}

// GameListInp 获取游戏列表
type GameListInp struct {
	form.PageReq
	GameId    int64         `json:"gameId"    dc:"自动编号"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *GameListInp) Filter(ctx context.Context) (err error) {
	return
}

type GameListModel struct {
	GameId    int64       `json:"gameId"    dc:"自动编号"`
	GameSku   string      `json:"gameSku"   dc:"唯一值"`
	Name      string      `json:"name"      dc:"游戏全名"`
	Title     string      `json:"title"     dc:"游戏名称"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int64       `json:"sort"      dc:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GameExportModel 导出游戏
type GameExportModel struct {
	GameId    int64       `json:"gameId"    dc:"自动编号"`
	GameSku   string      `json:"gameSku"   dc:"唯一值"`
	Name      string      `json:"name"      dc:"游戏全名"`
	Title     string      `json:"title"     dc:"游戏名称"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int64       `json:"sort"      dc:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GameMaxSortInp 获取游戏最大排序
type GameMaxSortInp struct{}

func (in *GameMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type GameMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// GameStatusInp 更新游戏状态
type GameStatusInp struct {
	GameId int64 `json:"gameId" v:"required#自动编号不能为空" dc:"自动编号"`
	Status int   `json:"status" dc:"状态"`
}

func (in *GameStatusInp) Filter(ctx context.Context) (err error) {
	if in.GameId <= 0 {
		err = gerror.New("自动编号不能为空")
		return
	}

	if in.Status < 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type GameStatusModel struct{}
