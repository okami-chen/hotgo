// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/card/api/admin/card"
	"hotgo/addons/card/service"
)

var (
	Card = cCard{}
)

type cCard struct{}

// List 查看卡片列表
func (c *cCard) List(ctx context.Context, req *card.ListReq) (res *card.ListRes, err error) {
	list, totalCount, err := service.SysCard().List(ctx, &req.CardListInp)
	if err != nil {
		return
	}

	res = new(card.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出卡片列表
func (c *cCard) Export(ctx context.Context, req *card.ExportReq) (res *card.ExportRes, err error) {
	err = service.SysCard().Export(ctx, &req.CardListInp)
	return
}

// Edit 更新卡片
func (c *cCard) Edit(ctx context.Context, req *card.EditReq) (res *card.EditRes, err error) {
	err = service.SysCard().Edit(ctx, &req.CardEditInp)
	return
}

// View 获取指定卡片信息
func (c *cCard) View(ctx context.Context, req *card.ViewReq) (res *card.ViewRes, err error) {
	data, err := service.SysCard().View(ctx, &req.CardViewInp)
	if err != nil {
		return
	}

	res = new(card.ViewRes)
	res.CardViewModel = data
	return
}

// Delete 删除卡片
func (c *cCard) Delete(ctx context.Context, req *card.DeleteReq) (res *card.DeleteRes, err error) {
	err = service.SysCard().Delete(ctx, &req.CardDeleteInp)
	return
}
