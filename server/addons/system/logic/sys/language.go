// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/system/dao"
	"hotgo/addons/system/model/input/sysin"
	"hotgo/addons/system/service"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysLanguage struct{}

func NewSysLanguage() *sSysLanguage {
	return &sSysLanguage{}
}

func init() {
	service.RegisterSysLanguage(NewSysLanguage())
}

// Model 语言ORM模型
func (s *sSysLanguage) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Language.Ctx(ctx), option...)
}

// List 获取语言列表
func (s *sSysLanguage) List(ctx context.Context, in *sysin.LanguageListInp) (list []*sysin.LanguageListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.Language.Columns().Id, in.Id)
	}

	// 查询语言英文
	if in.Language != "" {
		mod = mod.WhereLike(dao.Language.Columns().Language, in.Language)
	}

	// 查询语言中文
	if in.LanguageZh != "" {
		mod = mod.WhereLike(dao.Language.Columns().LanguageZh, in.LanguageZh)
	}

	// 查询语言简称
	if in.Code != "" {
		mod = mod.WhereLike(dao.Language.Columns().Code, in.Code)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取语言数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.LanguageListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Language.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取语言列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出语言
func (s *sSysLanguage) Export(ctx context.Context, in *sysin.LanguageListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.LanguageExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出语言-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.LanguageExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增语言
func (s *sSysLanguage) Edit(ctx context.Context, in *sysin.LanguageEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.LanguageUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改语言失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.LanguageInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增语言失败，请稍后重试！")
	}
	return
}

// Delete 删除语言
func (s *sSysLanguage) Delete(ctx context.Context, in *sysin.LanguageDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除语言失败，请稍后重试！")
		return
	}
	return
}

// View 获取语言指定信息
func (s *sSysLanguage) View(ctx context.Context, in *sysin.LanguageViewInp) (res *sysin.LanguageViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取语言信息，请稍后重试！")
		return
	}
	return
}
