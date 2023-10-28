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
	"hotgo/addons/card/dao"
	"hotgo/addons/card/model/input/sysin"
	"hotgo/addons/card/service"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysCard struct{}

func NewSysCard() *sSysCard {
	return &sSysCard{}
}

func init() {
	service.RegisterSysCard(NewSysCard())
}

// Model 卡片ORM模型
func (s *sSysCard) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Card.Ctx(ctx), option...)
}

// List 获取卡片列表
func (s *sSysCard) List(ctx context.Context, in *sysin.CardListInp) (list []*sysin.CardListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.Card.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Card.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取卡片数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.CardListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Card.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取卡片列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出卡片
func (s *sSysCard) Export(ctx context.Context, in *sysin.CardListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.CardExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出卡片-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.CardExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增卡片
func (s *sSysCard) Edit(ctx context.Context, in *sysin.CardEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.CardUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改卡片失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.CardInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增卡片失败，请稍后重试！")
	}
	return
}

// Delete 删除卡片
func (s *sSysCard) Delete(ctx context.Context, in *sysin.CardDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除卡片失败，请稍后重试！")
		return
	}
	return
}

// View 获取卡片指定信息
func (s *sSysCard) View(ctx context.Context, in *sysin.CardViewInp) (res *sysin.CardViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取卡片信息，请稍后重试！")
		return
	}
	return
}
