// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"hotgo/addons/bee/api/admin/systemsite"
	"hotgo/addons/bee/service"
)

var (
	SystemSite = cSystemSite{}
)

type cSystemSite struct{}

// List 查看站群列表
func (c *cSystemSite) List(ctx context.Context, req *systemsite.ListReq) (res *systemsite.ListRes, err error) {
	list, totalCount, err := service.SystemSite().List(ctx, &req.SystemSiteListInp)
	if err != nil {
		return
	}

	res = new(systemsite.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出站群列表
func (c *cSystemSite) Export(ctx context.Context, req *systemsite.ExportReq) (res *systemsite.ExportRes, err error) {
	err = service.SystemSite().Export(ctx, &req.SystemSiteListInp)
	return
}

// Edit 更新站群
func (c *cSystemSite) Edit(ctx context.Context, req *systemsite.EditReq) (res *systemsite.EditRes, err error) {
	err = service.SystemSite().Edit(ctx, &req.SystemSiteEditInp)
	return
}

// MaxSort 获取站群最大排序
func (c *cSystemSite) MaxSort(ctx context.Context, req *systemsite.MaxSortReq) (res *systemsite.MaxSortRes, err error) {
	data, err := service.SystemSite().MaxSort(ctx, &req.SystemSiteMaxSortInp)
	if err != nil {
		return
	}

	res = new(systemsite.MaxSortRes)
	res.SystemSiteMaxSortModel = data
	return
}

// View 获取指定站群信息
func (c *cSystemSite) View(ctx context.Context, req *systemsite.ViewReq) (res *systemsite.ViewRes, err error) {
	data, err := service.SystemSite().View(ctx, &req.SystemSiteViewInp)
	if err != nil {
		return
	}

	res = new(systemsite.ViewRes)
	res.SystemSiteViewModel = data
	return
}

// Delete 删除站群
func (c *cSystemSite) Delete(ctx context.Context, req *systemsite.DeleteReq) (res *systemsite.DeleteRes, err error) {
	err = service.SystemSite().Delete(ctx, &req.SystemSiteDeleteInp)
	return
}
