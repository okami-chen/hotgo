// Package village
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package village

import (
	"hotgo/addons/fang/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询小区列表
type ListReq struct {
	g.Meta `path:"/village/list" method:"get" tags:"小区" summary:"获取小区列表"`
	sysin.VillageListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.VillageListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出小区列表
type ExportReq struct {
	g.Meta `path:"/village/export" method:"get" tags:"小区" summary:"导出小区列表"`
	sysin.VillageListInp
}

type ExportRes struct{}

// ViewReq 获取小区指定信息
type ViewReq struct {
	g.Meta `path:"/village/view" method:"get" tags:"小区" summary:"获取小区指定信息"`
	sysin.VillageViewInp
}

type ViewRes struct {
	*sysin.VillageViewModel
}

// EditReq 修改/新增小区
type EditReq struct {
	g.Meta `path:"/village/edit" method:"post" tags:"小区" summary:"修改/新增小区"`
	sysin.VillageEditInp
}

type EditRes struct{}

// DeleteReq 删除小区
type DeleteReq struct {
	g.Meta `path:"/village/delete" method:"post" tags:"小区" summary:"删除小区"`
	sysin.VillageDeleteInp
}

type DeleteRes struct{}
