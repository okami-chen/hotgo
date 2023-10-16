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
	"hotgo/addons/cloudflare/dao"
	"hotgo/addons/cloudflare/model/input/sysin"
	"hotgo/addons/cloudflare/service"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDomain struct{}

func NewSysDomain() *sSysDomain {
	return &sSysDomain{}
}

func init() {
	service.RegisterSysDomain(NewSysDomain())
}

// Model 名ORM模型
func (s *sSysDomain) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Domain.Ctx(ctx), option...)
}

// List 获取名列表
func (s *sSysDomain) List(ctx context.Context, in *sysin.DomainListInp) (list []*sysin.DomainListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询域名名称
	if in.Domain != "" {
		mod = mod.WhereLike(dao.Domain.Columns().Domain, in.Domain)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取名数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.DomainListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Domain.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取名列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出名
func (s *sSysDomain) Export(ctx context.Context, in *sysin.DomainListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.DomainExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出名-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.DomainExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增名
func (s *sSysDomain) Edit(ctx context.Context, in *sysin.DomainEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.DomainUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改名失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.DomainInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增名失败，请稍后重试！")
	}
	return
}

// Delete 删除名
func (s *sSysDomain) Delete(ctx context.Context, in *sysin.DomainDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除名失败，请稍后重试！")
		return
	}
	return
}

// View 获取名指定信息
func (s *sSysDomain) View(ctx context.Context, in *sysin.DomainViewInp) (res *sysin.DomainViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取名信息，请稍后重试！")
		return
	}
	return
}
