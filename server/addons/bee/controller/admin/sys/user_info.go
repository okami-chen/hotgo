// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"hotgo/addons/bee/api/admin/userinfo"
	"hotgo/addons/bee/service"
)

var (
	UserInfo = cUserInfo{}
)

type cUserInfo struct{}

// List 查看客户信息扩展列表
func (c *cUserInfo) List(ctx context.Context, req *userinfo.ListReq) (res *userinfo.ListRes, err error) {
	list, totalCount, err := service.SysUserInfo().List(ctx, &req.UserInfoListInp)
	if err != nil {
		return
	}

	res = new(userinfo.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出客户信息扩展列表
func (c *cUserInfo) Export(ctx context.Context, req *userinfo.ExportReq) (res *userinfo.ExportRes, err error) {
	err = service.SysUserInfo().Export(ctx, &req.UserInfoListInp)
	return
}

// Edit 更新客户信息扩展
func (c *cUserInfo) Edit(ctx context.Context, req *userinfo.EditReq) (res *userinfo.EditRes, err error) {
	err = service.SysUserInfo().Edit(ctx, &req.UserInfoEditInp)
	return
}

// View 获取指定客户信息扩展信息
func (c *cUserInfo) View(ctx context.Context, req *userinfo.ViewReq) (res *userinfo.ViewRes, err error) {
	data, err := service.SysUserInfo().View(ctx, &req.UserInfoViewInp)
	if err != nil {
		return
	}

	res = new(userinfo.ViewRes)
	res.UserInfoViewModel = data
	return
}

// Delete 删除客户信息扩展
func (c *cUserInfo) Delete(ctx context.Context, req *userinfo.DeleteReq) (res *userinfo.DeleteRes, err error) {
	err = service.SysUserInfo().Delete(ctx, &req.UserInfoDeleteInp)
	return
}
