// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/game/dao"
	"hotgo/addons/game/model/input/sysin"
	"hotgo/addons/game/service"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysGame struct{}

func NewSysGame() *sSysGame {
	return &sSysGame{}
}

func init() {
	service.RegisterSysGame(NewSysGame())
}

// Model 游戏ORM模型
func (s *sSysGame) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Game.Ctx(ctx), option...)
}

// List 获取游戏列表
func (s *sSysGame) List(ctx context.Context, in *sysin.GameListInp) (list []*sysin.GameListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.GameId > 0 {
		mod = mod.Where(dao.Game.Columns().GameId, in.GameId)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.Game.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Game.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取游戏数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.GameListModel{}).Page(in.Page, in.PerPage).OrderAsc(dao.Game.Columns().Sort).OrderDesc(dao.Game.Columns().GameId).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取游戏列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出游戏
func (s *sSysGame) Export(ctx context.Context, in *sysin.GameListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.GameExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出游戏-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.GameExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增游戏
func (s *sSysGame) Edit(ctx context.Context, in *sysin.GameEditInp) (err error) {
	// 验证'GameSku'唯一
	if err = hgorm.IsUnique(ctx, &dao.Game, g.Map{dao.Game.Columns().GameSku: in.GameSku}, "唯一值已存在", in.GameId); err != nil {
		return
	}
	// 修改
	if in.GameId > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.GameUpdateFields{}).
			WherePri(in.GameId).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改游戏失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.GameInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增游戏失败，请稍后重试！")
	}
	return
}

// Delete 删除游戏
func (s *sSysGame) Delete(ctx context.Context, in *sysin.GameDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.GameId).Delete(); err != nil {
		err = gerror.Wrap(err, "删除游戏失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取游戏最大排序
func (s *sSysGame) MaxSort(ctx context.Context, in *sysin.GameMaxSortInp) (res *sysin.GameMaxSortModel, err error) {
	if err = dao.Game.Ctx(ctx).Fields(dao.Game.Columns().Sort).OrderDesc(dao.Game.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取游戏最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.GameMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取游戏指定信息
func (s *sSysGame) View(ctx context.Context, in *sysin.GameViewInp) (res *sysin.GameViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.GameId).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取游戏信息，请稍后重试！")
		return
	}
	return
}

// Status 更新游戏状态
func (s *sSysGame) Status(ctx context.Context, in *sysin.GameStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.GameId).Data(g.Map{
		dao.Game.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新游戏状态失败，请稍后重试！")
		return
	}
	return
}
