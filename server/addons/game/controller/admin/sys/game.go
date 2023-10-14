// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"hotgo/addons/game/api/admin/game"
	"hotgo/addons/game/service"
)

var (
	Game = cGame{}
)

type cGame struct{}

// List 查看游戏列表
func (c *cGame) List(ctx context.Context, req *game.ListReq) (res *game.ListRes, err error) {
	list, totalCount, err := service.SysGame().List(ctx, &req.GameListInp)
	if err != nil {
		return
	}

	res = new(game.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出游戏列表
func (c *cGame) Export(ctx context.Context, req *game.ExportReq) (res *game.ExportRes, err error) {
	err = service.SysGame().Export(ctx, &req.GameListInp)
	return
}

// Edit 更新游戏
func (c *cGame) Edit(ctx context.Context, req *game.EditReq) (res *game.EditRes, err error) {
	err = service.SysGame().Edit(ctx, &req.GameEditInp)
	return
}

// MaxSort 获取游戏最大排序
func (c *cGame) MaxSort(ctx context.Context, req *game.MaxSortReq) (res *game.MaxSortRes, err error) {
	data, err := service.SysGame().MaxSort(ctx, &req.GameMaxSortInp)
	if err != nil {
		return
	}

	res = new(game.MaxSortRes)
	res.GameMaxSortModel = data
	return
}

// View 获取指定游戏信息
func (c *cGame) View(ctx context.Context, req *game.ViewReq) (res *game.ViewRes, err error) {
	data, err := service.SysGame().View(ctx, &req.GameViewInp)
	if err != nil {
		return
	}

	res = new(game.ViewRes)
	res.GameViewModel = data
	return
}

// Delete 删除游戏
func (c *cGame) Delete(ctx context.Context, req *game.DeleteReq) (res *game.DeleteRes, err error) {
	err = service.SysGame().Delete(ctx, &req.GameDeleteInp)
	return
}

// Status 更新游戏状态
func (c *cGame) Status(ctx context.Context, req *game.StatusReq) (res *game.StatusRes, err error) {
	err = service.SysGame().Status(ctx, &req.GameStatusInp)
	return
}
