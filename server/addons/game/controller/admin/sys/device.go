// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"hotgo/addons/game/api/admin/device"
	"hotgo/addons/game/service"
)

var (
	Device = cDevice{}
)

type cDevice struct{}

// List 查看设备列表
func (c *cDevice) List(ctx context.Context, req *device.ListReq) (res *device.ListRes, err error) {
	list, totalCount, err := service.SysDevice().List(ctx, &req.DeviceListInp)
	if err != nil {
		return
	}

	res = new(device.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出设备列表
func (c *cDevice) Export(ctx context.Context, req *device.ExportReq) (res *device.ExportRes, err error) {
	err = service.SysDevice().Export(ctx, &req.DeviceListInp)
	return
}

// Edit 更新设备
func (c *cDevice) Edit(ctx context.Context, req *device.EditReq) (res *device.EditRes, err error) {
	err = service.SysDevice().Edit(ctx, &req.DeviceEditInp)
	return
}

// MaxSort 获取设备最大排序
func (c *cDevice) MaxSort(ctx context.Context, req *device.MaxSortReq) (res *device.MaxSortRes, err error) {
	data, err := service.SysDevice().MaxSort(ctx, &req.DeviceMaxSortInp)
	if err != nil {
		return
	}

	res = new(device.MaxSortRes)
	res.DeviceMaxSortModel = data
	return
}

// View 获取指定设备信息
func (c *cDevice) View(ctx context.Context, req *device.ViewReq) (res *device.ViewRes, err error) {
	data, err := service.SysDevice().View(ctx, &req.DeviceViewInp)
	if err != nil {
		return
	}

	res = new(device.ViewRes)
	res.DeviceViewModel = data
	return
}

// Delete 删除设备
func (c *cDevice) Delete(ctx context.Context, req *device.DeleteReq) (res *device.DeleteRes, err error) {
	err = service.SysDevice().Delete(ctx, &req.DeviceDeleteInp)
	return
}

// Status 更新设备状态
func (c *cDevice) Status(ctx context.Context, req *device.StatusReq) (res *device.StatusRes, err error) {
	err = service.SysDevice().Status(ctx, &req.DeviceStatusInp)
	return
}
