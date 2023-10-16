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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysCurrency struct{}

func NewSysCurrency() *sSysCurrency {
	return &sSysCurrency{}
}

func init() {
	service.RegisterSysCurrency(NewSysCurrency())
}

// Model 货币ORM模型
func (s *sSysCurrency) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Currency.Ctx(ctx), option...)
}

// List 获取货币列表
func (s *sSysCurrency) List(ctx context.Context, in *sysin.CurrencyListInp) (list []*sysin.CurrencyListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.Currency.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.Currency.Columns().Status, in.Status)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取货币数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.CurrencyListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Currency.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取货币列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出货币
func (s *sSysCurrency) Export(ctx context.Context, in *sysin.CurrencyListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.CurrencyExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出货币-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.CurrencyExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增货币
func (s *sSysCurrency) Edit(ctx context.Context, in *sysin.CurrencyEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.CurrencyUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改货币失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.CurrencyInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增货币失败，请稍后重试！")
	}
	return
}

// Delete 删除货币
func (s *sSysCurrency) Delete(ctx context.Context, in *sysin.CurrencyDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除货币失败，请稍后重试！")
		return
	}
	return
}

// View 获取货币指定信息
func (s *sSysCurrency) View(ctx context.Context, in *sysin.CurrencyViewInp) (res *sysin.CurrencyViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取货币信息，请稍后重试！")
		return
	}
	return
}

// Status 更新货币状态
func (s *sSysCurrency) Status(ctx context.Context, in *sysin.CurrencyStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.Currency.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新货币状态失败，请稍后重试！")
		return
	}
	return
}
