// Package game
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package game

import (
	"hotgo/addons/game/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询游戏列表
type ListReq struct {
	g.Meta `path:"/game/list" method:"get" tags:"游戏" summary:"获取游戏列表"`
	sysin.GameListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.GameListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出游戏列表
type ExportReq struct {
	g.Meta `path:"/game/export" method:"get" tags:"游戏" summary:"导出游戏列表"`
	sysin.GameListInp
}

type ExportRes struct{}

// ViewReq 获取游戏指定信息
type ViewReq struct {
	g.Meta `path:"/game/view" method:"get" tags:"游戏" summary:"获取游戏指定信息"`
	sysin.GameViewInp
}

type ViewRes struct {
	*sysin.GameViewModel
}

// EditReq 修改/新增游戏
type EditReq struct {
	g.Meta `path:"/game/edit" method:"post" tags:"游戏" summary:"修改/新增游戏"`
	sysin.GameEditInp
}

type EditRes struct{}

// DeleteReq 删除游戏
type DeleteReq struct {
	g.Meta `path:"/game/delete" method:"post" tags:"游戏" summary:"删除游戏"`
	sysin.GameDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取游戏最大排序
type MaxSortReq struct {
	g.Meta `path:"/game/maxSort" method:"get" tags:"游戏" summary:"获取游戏最大排序"`
	sysin.GameMaxSortInp
}

type MaxSortRes struct {
	*sysin.GameMaxSortModel
}

// StatusReq 更新游戏状态
type StatusReq struct {
	g.Meta `path:"/game/status" method:"post" tags:"游戏" summary:"更新游戏状态"`
	sysin.GameStatusInp
}

type StatusRes struct{}
