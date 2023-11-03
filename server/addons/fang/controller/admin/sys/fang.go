// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package sys

import (
	"context"
	"hotgo/addons/fang/api/admin/fang"
	"hotgo/addons/fang/service"
)

var (
	Fang = cFang{}
)

type cFang struct{}

// List 查看租房列表
func (c *cFang) List(ctx context.Context, req *fang.ListReq) (res *fang.ListRes, err error) {
	list, totalCount, err := service.SysFang().List(ctx, &req.FangListInp)
	if err != nil {
		return
	}

	res = new(fang.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出租房列表
func (c *cFang) Export(ctx context.Context, req *fang.ExportReq) (res *fang.ExportRes, err error) {
	err = service.SysFang().Export(ctx, &req.FangListInp)
	return
}

// Edit 更新租房
func (c *cFang) Edit(ctx context.Context, req *fang.EditReq) (res *fang.EditRes, err error) {
	err = service.SysFang().Edit(ctx, &req.FangEditInp)
	return
}

// View 获取指定租房信息
func (c *cFang) View(ctx context.Context, req *fang.ViewReq) (res *fang.ViewRes, err error) {
	data, err := service.SysFang().View(ctx, &req.FangViewInp)
	if err != nil {
		return
	}

	res = new(fang.ViewRes)
	res.FangViewModel = data
	return
}

// Delete 删除租房
func (c *cFang) Delete(ctx context.Context, req *fang.DeleteReq) (res *fang.DeleteRes, err error) {
	err = service.SysFang().Delete(ctx, &req.FangDeleteInp)
	return
}
