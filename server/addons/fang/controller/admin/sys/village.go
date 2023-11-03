// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package sys

import (
	"context"
	"hotgo/addons/fang/api/admin/village"
	"hotgo/addons/fang/service"
)

var (
	Village = cVillage{}
)

type cVillage struct{}

// List 查看小区列表
func (c *cVillage) List(ctx context.Context, req *village.ListReq) (res *village.ListRes, err error) {
	list, totalCount, err := service.SysVillage().List(ctx, &req.VillageListInp)
	if err != nil {
		return
	}

	res = new(village.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出小区列表
func (c *cVillage) Export(ctx context.Context, req *village.ExportReq) (res *village.ExportRes, err error) {
	err = service.SysVillage().Export(ctx, &req.VillageListInp)
	return
}

// Edit 更新小区
func (c *cVillage) Edit(ctx context.Context, req *village.EditReq) (res *village.EditRes, err error) {
	err = service.SysVillage().Edit(ctx, &req.VillageEditInp)
	return
}

// View 获取指定小区信息
func (c *cVillage) View(ctx context.Context, req *village.ViewReq) (res *village.ViewRes, err error) {
	data, err := service.SysVillage().View(ctx, &req.VillageViewInp)
	if err != nil {
		return
	}

	res = new(village.ViewRes)
	res.VillageViewModel = data
	return
}

// Delete 删除小区
func (c *cVillage) Delete(ctx context.Context, req *village.DeleteReq) (res *village.DeleteRes, err error) {
	err = service.SysVillage().Delete(ctx, &req.VillageDeleteInp)
	return
}
