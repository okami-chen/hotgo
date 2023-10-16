// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/system/api/admin/currency"
	"hotgo/addons/system/service"
)

var (
	Currency = cCurrency{}
)

type cCurrency struct{}

// List 查看货币列表
func (c *cCurrency) List(ctx context.Context, req *currency.ListReq) (res *currency.ListRes, err error) {
	list, totalCount, err := service.SysCurrency().List(ctx, &req.CurrencyListInp)
	if err != nil {
		return
	}

	res = new(currency.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出货币列表
func (c *cCurrency) Export(ctx context.Context, req *currency.ExportReq) (res *currency.ExportRes, err error) {
	err = service.SysCurrency().Export(ctx, &req.CurrencyListInp)
	return
}

// Edit 更新货币
func (c *cCurrency) Edit(ctx context.Context, req *currency.EditReq) (res *currency.EditRes, err error) {
	err = service.SysCurrency().Edit(ctx, &req.CurrencyEditInp)
	return
}

// View 获取指定货币信息
func (c *cCurrency) View(ctx context.Context, req *currency.ViewReq) (res *currency.ViewRes, err error) {
	data, err := service.SysCurrency().View(ctx, &req.CurrencyViewInp)
	if err != nil {
		return
	}

	res = new(currency.ViewRes)
	res.CurrencyViewModel = data
	return
}

// Delete 删除货币
func (c *cCurrency) Delete(ctx context.Context, req *currency.DeleteReq) (res *currency.DeleteRes, err error) {
	err = service.SysCurrency().Delete(ctx, &req.CurrencyDeleteInp)
	return
}

// Status 更新货币状态
func (c *cCurrency) Status(ctx context.Context, req *currency.StatusReq) (res *currency.StatusRes, err error) {
	err = service.SysCurrency().Status(ctx, &req.CurrencyStatusInp)
	return
}
