// Package currency
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package currency

import (
	"hotgo/addons/system/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询货币列表
type ListReq struct {
	g.Meta `path:"/currency/list" method:"get" tags:"货币" summary:"获取货币列表"`
	sysin.CurrencyListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.CurrencyListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出货币列表
type ExportReq struct {
	g.Meta `path:"/currency/export" method:"get" tags:"货币" summary:"导出货币列表"`
	sysin.CurrencyListInp
}

type ExportRes struct{}

// ViewReq 获取货币指定信息
type ViewReq struct {
	g.Meta `path:"/currency/view" method:"get" tags:"货币" summary:"获取货币指定信息"`
	sysin.CurrencyViewInp
}

type ViewRes struct {
	*sysin.CurrencyViewModel
}

// EditReq 修改/新增货币
type EditReq struct {
	g.Meta `path:"/currency/edit" method:"post" tags:"货币" summary:"修改/新增货币"`
	sysin.CurrencyEditInp
}

type EditRes struct{}

// DeleteReq 删除货币
type DeleteReq struct {
	g.Meta `path:"/currency/delete" method:"post" tags:"货币" summary:"删除货币"`
	sysin.CurrencyDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新货币状态
type StatusReq struct {
	g.Meta `path:"/currency/status" method:"post" tags:"货币" summary:"更新货币状态"`
	sysin.CurrencyStatusInp
}

type StatusRes struct{}
