// Package domain
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package domain

import (
	"hotgo/addons/cloudflare/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询名列表
type ListReq struct {
	g.Meta `path:"/domain/list" method:"get" tags:"名" summary:"获取名列表"`
	sysin.DomainListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.DomainListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出名列表
type ExportReq struct {
	g.Meta `path:"/domain/export" method:"get" tags:"名" summary:"导出名列表"`
	sysin.DomainListInp
}

type ExportRes struct{}

// ViewReq 获取名指定信息
type ViewReq struct {
	g.Meta `path:"/domain/view" method:"get" tags:"名" summary:"获取名指定信息"`
	sysin.DomainViewInp
}

type ViewRes struct {
	*sysin.DomainViewModel
}

// EditReq 修改/新增名
type EditReq struct {
	g.Meta `path:"/domain/edit" method:"post" tags:"名" summary:"修改/新增名"`
	sysin.DomainEditInp
}

type EditRes struct{}

// DeleteReq 删除名
type DeleteReq struct {
	g.Meta `path:"/domain/delete" method:"post" tags:"名" summary:"删除名"`
	sysin.DomainDeleteInp
}

type DeleteRes struct{}
