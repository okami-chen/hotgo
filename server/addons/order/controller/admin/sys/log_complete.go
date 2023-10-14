// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	logcomplete "hotgo/addons/order/api/admin/log"
	"hotgo/addons/order/service"
)

var (
	LogComplete = cLogComplete{}
)

type cLogComplete struct{}

// List 查看完成时间列表
func (c *cLogComplete) List(ctx context.Context, req *logcomplete.ListReq) (res *logcomplete.ListRes, err error) {
	list, totalCount, err := service.SysLogComplete().List(ctx, &req.LogCompleteListInp)
	if err != nil {
		return
	}

	res = new(logcomplete.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出完成时间列表
func (c *cLogComplete) Export(ctx context.Context, req *logcomplete.ExportReq) (res *logcomplete.ExportRes, err error) {
	err = service.SysLogComplete().Export(ctx, &req.LogCompleteListInp)
	return
}

// Edit 更新完成时间
func (c *cLogComplete) Edit(ctx context.Context, req *logcomplete.EditReq) (res *logcomplete.EditRes, err error) {
	err = service.SysLogComplete().Edit(ctx, &req.LogCompleteEditInp)
	return
}

// View 获取指定完成时间信息
func (c *cLogComplete) View(ctx context.Context, req *logcomplete.ViewReq) (res *logcomplete.ViewRes, err error) {
	data, err := service.SysLogComplete().View(ctx, &req.LogCompleteViewInp)
	if err != nil {
		return
	}

	res = new(logcomplete.ViewRes)
	res.LogCompleteViewModel = data
	return
}

// Delete 删除完成时间
func (c *cLogComplete) Delete(ctx context.Context, req *logcomplete.DeleteReq) (res *logcomplete.DeleteRes, err error) {
	err = service.SysLogComplete().Delete(ctx, &req.LogCompleteDeleteInp)
	return
}
