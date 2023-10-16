// Package account
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package account

import (
	"hotgo/addons/cloudflare/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询账号列表
type ListReq struct {
	g.Meta `path:"/account/list" method:"get" tags:"账号" summary:"获取账号列表"`
	sysin.AccountListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AccountListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出账号列表
type ExportReq struct {
	g.Meta `path:"/account/export" method:"get" tags:"账号" summary:"导出账号列表"`
	sysin.AccountListInp
}

type ExportRes struct{}

// ViewReq 获取账号指定信息
type ViewReq struct {
	g.Meta `path:"/account/view" method:"get" tags:"账号" summary:"获取账号指定信息"`
	sysin.AccountViewInp
}

type ViewRes struct {
	*sysin.AccountViewModel
}

// EditReq 修改/新增账号
type EditReq struct {
	g.Meta `path:"/account/edit" method:"post" tags:"账号" summary:"修改/新增账号"`
	sysin.AccountEditInp
}

type EditRes struct{}

// DeleteReq 删除账号
type DeleteReq struct {
	g.Meta `path:"/account/delete" method:"post" tags:"账号" summary:"删除账号"`
	sysin.AccountDeleteInp
}

type DeleteRes struct{}
