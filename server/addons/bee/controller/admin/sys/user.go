// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"hotgo/addons/bee/api/admin/user"
	"hotgo/addons/bee/service"
)

var (
	User = cUser{}
)

type cUser struct{}

// List 查看用户信息列表
func (c *cUser) List(ctx context.Context, req *user.ListReq) (res *user.ListRes, err error) {
	list, totalCount, err := service.SysUser().List(ctx, &req.UserListInp)
	if err != nil {
		return
	}

	res = new(user.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出用户信息列表
func (c *cUser) Export(ctx context.Context, req *user.ExportReq) (res *user.ExportRes, err error) {
	err = service.SysUser().Export(ctx, &req.UserListInp)
	return
}

// Edit 更新用户信息
func (c *cUser) Edit(ctx context.Context, req *user.EditReq) (res *user.EditRes, err error) {
	err = service.SysUser().Edit(ctx, &req.UserEditInp)
	return
}

// View 获取指定用户信息信息
func (c *cUser) View(ctx context.Context, req *user.ViewReq) (res *user.ViewRes, err error) {
	data, err := service.SysUser().View(ctx, &req.UserViewInp)
	if err != nil {
		return
	}

	res = new(user.ViewRes)
	res.UserViewModel = data
	return
}

// Delete 删除用户信息
func (c *cUser) Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	err = service.SysUser().Delete(ctx, &req.UserDeleteInp)
	return
}

// Status 更新用户信息状态
func (c *cUser) Status(ctx context.Context, req *user.StatusReq) (res *user.StatusRes, err error) {
	err = service.SysUser().Status(ctx, &req.UserStatusInp)
	return
}
