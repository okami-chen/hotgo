// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/fang/dao"
	"hotgo/addons/fang/model/input/sysin"
	"hotgo/addons/fang/service"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysVillage struct{}

func NewSysVillage() *sSysVillage {
	return &sSysVillage{}
}

func init() {
	service.RegisterSysVillage(NewSysVillage())
}

// Model 小区ORM模型
func (s *sSysVillage) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Village.Ctx(ctx), option...)
}

// List 获取小区列表
func (s *sSysVillage) List(ctx context.Context, in *sysin.VillageListInp) (list []*sysin.VillageListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询名称
	if in.Name != "" {
		mod = mod.WhereLike(dao.Village.Columns().Name, in.Name)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取小区数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.VillageListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Village.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取小区列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出小区
func (s *sSysVillage) Export(ctx context.Context, in *sysin.VillageListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.VillageExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出小区-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.VillageExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增小区
func (s *sSysVillage) Edit(ctx context.Context, in *sysin.VillageEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.VillageUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改小区失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.VillageInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增小区失败，请稍后重试！")
	}
	return
}

// Delete 删除小区
func (s *sSysVillage) Delete(ctx context.Context, in *sysin.VillageDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除小区失败，请稍后重试！")
		return
	}
	return
}

// View 获取小区指定信息
func (s *sSysVillage) View(ctx context.Context, in *sysin.VillageViewInp) (res *sysin.VillageViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取小区信息，请稍后重试！")
		return
	}
	return
}
