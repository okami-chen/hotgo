// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/cloudflare/api/admin/account"
	"hotgo/addons/cloudflare/service"
)

var (
	Account = cAccount{}
)

type cAccount struct{}

// List 查看账号列表
func (c *cAccount) List(ctx context.Context, req *account.ListReq) (res *account.ListRes, err error) {
	list, totalCount, err := service.SysAccount().List(ctx, &req.AccountListInp)
	if err != nil {
		return
	}

	res = new(account.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出账号列表
func (c *cAccount) Export(ctx context.Context, req *account.ExportReq) (res *account.ExportRes, err error) {
	err = service.SysAccount().Export(ctx, &req.AccountListInp)
	return
}

// Edit 更新账号
func (c *cAccount) Edit(ctx context.Context, req *account.EditReq) (res *account.EditRes, err error) {
	err = service.SysAccount().Edit(ctx, &req.AccountEditInp)
	return
}

// View 获取指定账号信息
func (c *cAccount) View(ctx context.Context, req *account.ViewReq) (res *account.ViewRes, err error) {
	data, err := service.SysAccount().View(ctx, &req.AccountViewInp)
	if err != nil {
		return
	}

	res = new(account.ViewRes)
	res.AccountViewModel = data
	return
}

// Delete 删除账号
func (c *cAccount) Delete(ctx context.Context, req *account.DeleteReq) (res *account.DeleteRes, err error) {
	err = service.SysAccount().Delete(ctx, &req.AccountDeleteInp)
	return
}
