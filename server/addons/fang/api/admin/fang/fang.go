// Package fang
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package fang

import (
	"hotgo/addons/fang/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询租房列表
type ListReq struct {
	g.Meta `path:"/fang/list" method:"get" tags:"租房" summary:"获取租房列表"`
	sysin.FangListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FangListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出租房列表
type ExportReq struct {
	g.Meta `path:"/fang/export" method:"get" tags:"租房" summary:"导出租房列表"`
	sysin.FangListInp
}

type ExportRes struct{}

// ViewReq 获取租房指定信息
type ViewReq struct {
	g.Meta `path:"/fang/view" method:"get" tags:"租房" summary:"获取租房指定信息"`
	sysin.FangViewInp
}

type ViewRes struct {
	*sysin.FangViewModel
}

// EditReq 修改/新增租房
type EditReq struct {
	g.Meta `path:"/fang/edit" method:"post" tags:"租房" summary:"修改/新增租房"`
	sysin.FangEditInp
}

type EditRes struct{}

// DeleteReq 删除租房
type DeleteReq struct {
	g.Meta `path:"/fang/delete" method:"post" tags:"租房" summary:"删除租房"`
	sysin.FangDeleteInp
}

type DeleteRes struct{}
