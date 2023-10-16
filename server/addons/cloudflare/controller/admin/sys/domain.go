// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/cloudflare/api/admin/domain"
	"hotgo/addons/cloudflare/service"
)

var (
	Domain = cDomain{}
)

type cDomain struct{}

// List 查看名列表
func (c *cDomain) List(ctx context.Context, req *domain.ListReq) (res *domain.ListRes, err error) {
	list, totalCount, err := service.SysDomain().List(ctx, &req.DomainListInp)
	if err != nil {
		return
	}

	res = new(domain.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出名列表
func (c *cDomain) Export(ctx context.Context, req *domain.ExportReq) (res *domain.ExportRes, err error) {
	err = service.SysDomain().Export(ctx, &req.DomainListInp)
	return
}

// Edit 更新名
func (c *cDomain) Edit(ctx context.Context, req *domain.EditReq) (res *domain.EditRes, err error) {
	err = service.SysDomain().Edit(ctx, &req.DomainEditInp)
	return
}

// View 获取指定名信息
func (c *cDomain) View(ctx context.Context, req *domain.ViewReq) (res *domain.ViewRes, err error) {
	data, err := service.SysDomain().View(ctx, &req.DomainViewInp)
	if err != nil {
		return
	}

	res = new(domain.ViewRes)
	res.DomainViewModel = data
	return
}

// Delete 删除名
func (c *cDomain) Delete(ctx context.Context, req *domain.DeleteReq) (res *domain.DeleteRes, err error) {
	err = service.SysDomain().Delete(ctx, &req.DomainDeleteInp)
	return
}
