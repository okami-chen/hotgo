// Package card
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package card

import (
	"hotgo/addons/card/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询卡片列表
type ListReq struct {
	g.Meta `path:"/card/list" method:"get" tags:"卡片" summary:"获取卡片列表"`
	sysin.CardListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.CardListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出卡片列表
type ExportReq struct {
	g.Meta `path:"/card/export" method:"get" tags:"卡片" summary:"导出卡片列表"`
	sysin.CardListInp
}

type ExportRes struct{}

// ViewReq 获取卡片指定信息
type ViewReq struct {
	g.Meta `path:"/card/view" method:"get" tags:"卡片" summary:"获取卡片指定信息"`
	sysin.CardViewInp
}

type ViewRes struct {
	*sysin.CardViewModel
}

// EditReq 修改/新增卡片
type EditReq struct {
	g.Meta `path:"/card/edit" method:"post" tags:"卡片" summary:"修改/新增卡片"`
	sysin.CardEditInp
}

type EditRes struct{}

// DeleteReq 删除卡片
type DeleteReq struct {
	g.Meta `path:"/card/delete" method:"post" tags:"卡片" summary:"删除卡片"`
	sysin.CardDeleteInp
}

type DeleteRes struct{}
