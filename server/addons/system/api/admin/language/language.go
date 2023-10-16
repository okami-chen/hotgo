// Package language
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package language

import (
	"hotgo/addons/system/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询语言列表
type ListReq struct {
	g.Meta `path:"/language/list" method:"get" tags:"语言" summary:"获取语言列表"`
	sysin.LanguageListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.LanguageListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出语言列表
type ExportReq struct {
	g.Meta `path:"/language/export" method:"get" tags:"语言" summary:"导出语言列表"`
	sysin.LanguageListInp
}

type ExportRes struct{}

// ViewReq 获取语言指定信息
type ViewReq struct {
	g.Meta `path:"/language/view" method:"get" tags:"语言" summary:"获取语言指定信息"`
	sysin.LanguageViewInp
}

type ViewRes struct {
	*sysin.LanguageViewModel
}

// EditReq 修改/新增语言
type EditReq struct {
	g.Meta `path:"/language/edit" method:"post" tags:"语言" summary:"修改/新增语言"`
	sysin.LanguageEditInp
}

type EditRes struct{}

// DeleteReq 删除语言
type DeleteReq struct {
	g.Meta `path:"/language/delete" method:"post" tags:"语言" summary:"删除语言"`
	sysin.LanguageDeleteInp
}

type DeleteRes struct{}
