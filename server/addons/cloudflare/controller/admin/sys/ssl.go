// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/cloudflare/api/admin/ssl"
	"hotgo/addons/cloudflare/service"
)

var (
	Ssl = cSsl{}
)

type cSsl struct{}

// List 查看证书列表
func (c *cSsl) List(ctx context.Context, req *ssl.ListReq) (res *ssl.ListRes, err error) {
	list, totalCount, err := service.SysSsl().List(ctx, &req.SslListInp)
	if err != nil {
		return
	}

	res = new(ssl.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出证书列表
func (c *cSsl) Export(ctx context.Context, req *ssl.ExportReq) (res *ssl.ExportRes, err error) {
	err = service.SysSsl().Export(ctx, &req.SslListInp)
	return
}

// Edit 更新证书
func (c *cSsl) Edit(ctx context.Context, req *ssl.EditReq) (res *ssl.EditRes, err error) {
	err = service.SysSsl().Edit(ctx, &req.SslEditInp)
	return
}

// View 获取指定证书信息
func (c *cSsl) View(ctx context.Context, req *ssl.ViewReq) (res *ssl.ViewRes, err error) {
	data, err := service.SysSsl().View(ctx, &req.SslViewInp)
	if err != nil {
		return
	}

	res = new(ssl.ViewRes)
	res.SslViewModel = data
	return
}

// Delete 删除证书
func (c *cSsl) Delete(ctx context.Context, req *ssl.DeleteReq) (res *ssl.DeleteRes, err error) {
	err = service.SysSsl().Delete(ctx, &req.SslDeleteInp)
	return
}

// Status 更新证书状态
func (c *cSsl) Status(ctx context.Context, req *ssl.StatusReq) (res *ssl.StatusRes, err error) {
	err = service.SysSsl().Status(ctx, &req.SslStatusInp)
	return
}
