// Package systemsite
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package systemsite

import (
	"hotgo/addons/bee/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询站群列表
type ListReq struct {
	g.Meta `path:"/systemSite/list" method:"get" tags:"站群" summary:"获取站群列表"`
	sysin.SystemSiteListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.SystemSiteListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出站群列表
type ExportReq struct {
	g.Meta `path:"/systemSite/export" method:"get" tags:"站群" summary:"导出站群列表"`
	sysin.SystemSiteListInp
}

type ExportRes struct{}

// ViewReq 获取站群指定信息
type ViewReq struct {
	g.Meta `path:"/systemSite/view" method:"get" tags:"站群" summary:"获取站群指定信息"`
	sysin.SystemSiteViewInp
}

type ViewRes struct {
	*sysin.SystemSiteViewModel
}

// EditReq 修改/新增站群
type EditReq struct {
	g.Meta `path:"/systemSite/edit" method:"post" tags:"站群" summary:"修改/新增站群"`
	sysin.SystemSiteEditInp
}

type EditRes struct{}

// DeleteReq 删除站群
type DeleteReq struct {
	g.Meta `path:"/systemSite/delete" method:"post" tags:"站群" summary:"删除站群"`
	sysin.SystemSiteDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取站群最大排序
type MaxSortReq struct {
	g.Meta `path:"/systemSite/maxSort" method:"get" tags:"站群" summary:"获取站群最大排序"`
	sysin.SystemSiteMaxSortInp
}

type MaxSortRes struct {
	*sysin.SystemSiteMaxSortModel
}
