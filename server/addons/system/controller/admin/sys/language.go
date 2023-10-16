// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/system/api/admin/language"
	"hotgo/addons/system/service"
)

var (
	Language = cLanguage{}
)

type cLanguage struct{}

// List 查看语言列表
func (c *cLanguage) List(ctx context.Context, req *language.ListReq) (res *language.ListRes, err error) {
	list, totalCount, err := service.SysLanguage().List(ctx, &req.LanguageListInp)
	if err != nil {
		return
	}

	res = new(language.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出语言列表
func (c *cLanguage) Export(ctx context.Context, req *language.ExportReq) (res *language.ExportRes, err error) {
	err = service.SysLanguage().Export(ctx, &req.LanguageListInp)
	return
}

// Edit 更新语言
func (c *cLanguage) Edit(ctx context.Context, req *language.EditReq) (res *language.EditRes, err error) {
	err = service.SysLanguage().Edit(ctx, &req.LanguageEditInp)
	return
}

// View 获取指定语言信息
func (c *cLanguage) View(ctx context.Context, req *language.ViewReq) (res *language.ViewRes, err error) {
	data, err := service.SysLanguage().View(ctx, &req.LanguageViewInp)
	if err != nil {
		return
	}

	res = new(language.ViewRes)
	res.LanguageViewModel = data
	return
}

// Delete 删除语言
func (c *cLanguage) Delete(ctx context.Context, req *language.DeleteReq) (res *language.DeleteRes, err error) {
	err = service.SysLanguage().Delete(ctx, &req.LanguageDeleteInp)
	return
}
