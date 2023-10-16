// Package country
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package country

import (
	"hotgo/addons/system/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询国家列表
type ListReq struct {
	g.Meta `path:"/country/list" method:"get" tags:"国家" summary:"获取国家列表"`
	sysin.CountryListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.CountryListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出国家列表
type ExportReq struct {
	g.Meta `path:"/country/export" method:"get" tags:"国家" summary:"导出国家列表"`
	sysin.CountryListInp
}

type ExportRes struct{}

// ViewReq 获取国家指定信息
type ViewReq struct {
	g.Meta `path:"/country/view" method:"get" tags:"国家" summary:"获取国家指定信息"`
	sysin.CountryViewInp
}

type ViewRes struct {
	*sysin.CountryViewModel
}

// EditReq 修改/新增国家
type EditReq struct {
	g.Meta `path:"/country/edit" method:"post" tags:"国家" summary:"修改/新增国家"`
	sysin.CountryEditInp
}

type EditRes struct{}

// DeleteReq 删除国家
type DeleteReq struct {
	g.Meta `path:"/country/delete" method:"post" tags:"国家" summary:"删除国家"`
	sysin.CountryDeleteInp
}

type DeleteRes struct{}
