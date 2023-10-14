// Package systemevent
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package system

import (
	"hotgo/addons/bee/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询系统事件列表
type ListReq struct {
	g.Meta `path:"/system/event/list" method:"get" tags:"系统事件" summary:"获取系统事件列表"`
	sysin.SystemEventListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.SystemEventListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出系统事件列表
type ExportReq struct {
	g.Meta `path:"/system/event/export" method:"get" tags:"系统事件" summary:"导出系统事件列表"`
	sysin.SystemEventListInp
}

type ExportRes struct{}

// ViewReq 获取系统事件指定信息
type ViewReq struct {
	g.Meta `path:"/system/event/view" method:"get" tags:"系统事件" summary:"获取系统事件指定信息"`
	sysin.SystemEventViewInp
}

type ViewRes struct {
	*sysin.SystemEventViewModel
}

// EditReq 修改/新增系统事件
type EditReq struct {
	g.Meta `path:"/system/event/edit" method:"post" tags:"系统事件" summary:"修改/新增系统事件"`
	sysin.SystemEventEditInp
}

type EditRes struct{}

// DeleteReq 删除系统事件
type DeleteReq struct {
	g.Meta `path:"/system/event/delete" method:"post" tags:"系统事件" summary:"删除系统事件"`
	sysin.SystemEventDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新系统事件状态
type StatusReq struct {
	g.Meta `path:"/system/event/status" method:"post" tags:"系统事件" summary:"更新系统事件状态"`
	sysin.SystemEventStatusInp
}

type StatusRes struct{}
