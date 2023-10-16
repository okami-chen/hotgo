// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/system/api/admin/country"
	"hotgo/addons/system/service"
)

var (
	Country = cCountry{}
)

type cCountry struct{}

// List 查看国家列表
func (c *cCountry) List(ctx context.Context, req *country.ListReq) (res *country.ListRes, err error) {
	list, totalCount, err := service.SysCountry().List(ctx, &req.CountryListInp)
	if err != nil {
		return
	}

	res = new(country.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出国家列表
func (c *cCountry) Export(ctx context.Context, req *country.ExportReq) (res *country.ExportRes, err error) {
	err = service.SysCountry().Export(ctx, &req.CountryListInp)
	return
}

// Edit 更新国家
func (c *cCountry) Edit(ctx context.Context, req *country.EditReq) (res *country.EditRes, err error) {
	err = service.SysCountry().Edit(ctx, &req.CountryEditInp)
	return
}

// View 获取指定国家信息
func (c *cCountry) View(ctx context.Context, req *country.ViewReq) (res *country.ViewRes, err error) {
	data, err := service.SysCountry().View(ctx, &req.CountryViewInp)
	if err != nil {
		return
	}

	res = new(country.ViewRes)
	res.CountryViewModel = data
	return
}

// Delete 删除国家
func (c *cCountry) Delete(ctx context.Context, req *country.DeleteReq) (res *country.DeleteRes, err error) {
	err = service.SysCountry().Delete(ctx, &req.CountryDeleteInp)
	return
}
