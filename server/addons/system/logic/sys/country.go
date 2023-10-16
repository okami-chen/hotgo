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

type sSysCountry struct{}

func NewSysCountry() *sSysCountry {
	return &sSysCountry{}
}

func init() {
	service.RegisterSysCountry(NewSysCountry())
}

// Model 国家ORM模型
func (s *sSysCountry) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Country.Ctx(ctx), option...)
}

// List 获取国家列表
func (s *sSysCountry) List(ctx context.Context, in *sysin.CountryListInp) (list []*sysin.CountryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询国家缩写
	if in.Country != "" {
		mod = mod.WhereLike(dao.Country.Columns().Country, in.Country)
	}

	// 查询国家名称全程
	if in.CountryName != "" {
		mod = mod.WhereLike(dao.Country.Columns().CountryName, in.CountryName)
	}

	// 查询国家区号
	if in.DiallingCode != "" {
		mod = mod.WhereLike(dao.Country.Columns().DiallingCode, in.DiallingCode)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取国家数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.CountryListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Country.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取国家列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出国家
func (s *sSysCountry) Export(ctx context.Context, in *sysin.CountryListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.CountryExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出国家-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.CountryExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增国家
func (s *sSysCountry) Edit(ctx context.Context, in *sysin.CountryEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.CountryUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改国家失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.CountryInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增国家失败，请稍后重试！")
	}
	return
}

// Delete 删除国家
func (s *sSysCountry) Delete(ctx context.Context, in *sysin.CountryDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除国家失败，请稍后重试！")
		return
	}
	return
}

// View 获取国家指定信息
func (s *sSysCountry) View(ctx context.Context, in *sysin.CountryViewInp) (res *sysin.CountryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取国家信息，请稍后重试！")
		return
	}
	return
}
