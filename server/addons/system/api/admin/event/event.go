// Package event
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package event

import (
	"hotgo/addons/system/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询事件列表
type ListReq struct {
	g.Meta `path:"/event/list" method:"get" tags:"事件" summary:"获取事件列表"`
	sysin.EventListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.EventListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出事件列表
type ExportReq struct {
	g.Meta `path:"/event/export" method:"get" tags:"事件" summary:"导出事件列表"`
	sysin.EventListInp
}

type ExportRes struct{}

// ViewReq 获取事件指定信息
type ViewReq struct {
	g.Meta `path:"/event/view" method:"get" tags:"事件" summary:"获取事件指定信息"`
	sysin.EventViewInp
}

type ViewRes struct {
	*sysin.EventViewModel
}

// EditReq 修改/新增事件
type EditReq struct {
	g.Meta `path:"/event/edit" method:"post" tags:"事件" summary:"修改/新增事件"`
	sysin.EventEditInp
}

type EditRes struct{}

// DeleteReq 删除事件
type DeleteReq struct {
	g.Meta `path:"/event/delete" method:"post" tags:"事件" summary:"删除事件"`
	sysin.EventDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新事件状态
type StatusReq struct {
	g.Meta `path:"/event/status" method:"post" tags:"事件" summary:"更新事件状态"`
	sysin.EventStatusInp
}

type StatusRes struct{}
