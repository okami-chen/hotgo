// Package logcomplete
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package logcomplete

import (
	"hotgo/addons/order/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询完成时间列表
type ListReq struct {
	g.Meta `path:"/log/complete/list" method:"get" tags:"完成时间" summary:"获取完成时间列表"`
	sysin.LogCompleteListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.LogCompleteListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出完成时间列表
type ExportReq struct {
	g.Meta `path:"/log/complete/export" method:"get" tags:"完成时间" summary:"导出完成时间列表"`
	sysin.LogCompleteListInp
}

type ExportRes struct{}

// ViewReq 获取完成时间指定信息
type ViewReq struct {
	g.Meta `path:"/log/complete/view" method:"get" tags:"完成时间" summary:"获取完成时间指定信息"`
	sysin.LogCompleteViewInp
}

type ViewRes struct {
	*sysin.LogCompleteViewModel
}

// EditReq 修改/新增完成时间
type EditReq struct {
	g.Meta `path:"/log/complete/edit" method:"post" tags:"完成时间" summary:"修改/新增完成时间"`
	sysin.LogCompleteEditInp
}

type EditRes struct{}

// DeleteReq 删除完成时间
type DeleteReq struct {
	g.Meta `path:"/log/complete/delete" method:"post" tags:"完成时间" summary:"删除完成时间"`
	sysin.LogCompleteDeleteInp
}

type DeleteRes struct{}
