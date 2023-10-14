// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"hotgo/addons/system/api/admin/event"
	"hotgo/addons/system/service"
)

var (
	Event = cEvent{}
)

type cEvent struct{}

// List 查看事件列表
func (c *cEvent) List(ctx context.Context, req *event.ListReq) (res *event.ListRes, err error) {
	list, totalCount, err := service.SysEvent().List(ctx, &req.EventListInp)
	if err != nil {
		return
	}

	res = new(event.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出事件列表
func (c *cEvent) Export(ctx context.Context, req *event.ExportReq) (res *event.ExportRes, err error) {
	err = service.SysEvent().Export(ctx, &req.EventListInp)
	return
}

// Edit 更新事件
func (c *cEvent) Edit(ctx context.Context, req *event.EditReq) (res *event.EditRes, err error) {
	err = service.SysEvent().Edit(ctx, &req.EventEditInp)
	return
}

// View 获取指定事件信息
func (c *cEvent) View(ctx context.Context, req *event.ViewReq) (res *event.ViewRes, err error) {
	data, err := service.SysEvent().View(ctx, &req.EventViewInp)
	if err != nil {
		return
	}

	res = new(event.ViewRes)
	res.EventViewModel = data
	return
}

// Delete 删除事件
func (c *cEvent) Delete(ctx context.Context, req *event.DeleteReq) (res *event.DeleteRes, err error) {
	err = service.SysEvent().Delete(ctx, &req.EventDeleteInp)
	return
}

// Status 更新事件状态
func (c *cEvent) Status(ctx context.Context, req *event.StatusReq) (res *event.StatusRes, err error) {
	err = service.SysEvent().Status(ctx, &req.EventStatusInp)
	return
}
