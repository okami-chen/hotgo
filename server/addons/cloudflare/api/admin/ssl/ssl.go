// Package ssl
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package ssl

import (
	"hotgo/addons/cloudflare/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询证书列表
type ListReq struct {
	g.Meta `path:"/ssl/list" method:"get" tags:"证书" summary:"获取证书列表"`
	sysin.SslListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.SslListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出证书列表
type ExportReq struct {
	g.Meta `path:"/ssl/export" method:"get" tags:"证书" summary:"导出证书列表"`
	sysin.SslListInp
}

type ExportRes struct{}

// ViewReq 获取证书指定信息
type ViewReq struct {
	g.Meta `path:"/ssl/view" method:"get" tags:"证书" summary:"获取证书指定信息"`
	sysin.SslViewInp
}

type ViewRes struct {
	*sysin.SslViewModel
}

// EditReq 修改/新增证书
type EditReq struct {
	g.Meta `path:"/ssl/edit" method:"post" tags:"证书" summary:"修改/新增证书"`
	sysin.SslEditInp
}

type EditRes struct{}

// DeleteReq 删除证书
type DeleteReq struct {
	g.Meta `path:"/ssl/delete" method:"post" tags:"证书" summary:"删除证书"`
	sysin.SslDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新证书状态
type StatusReq struct {
	g.Meta `path:"/ssl/status" method:"post" tags:"证书" summary:"更新证书状态"`
	sysin.SslStatusInp
}

type StatusRes struct{}
