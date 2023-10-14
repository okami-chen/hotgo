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
	"hotgo/addons/bee/dao"
	"hotgo/addons/bee/model/input/sysin"
	"hotgo/addons/bee/service"
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

type sSystemSite struct{}

func NewSystemSite() *sSystemSite {
	return &sSystemSite{}
}

func init() {
	service.RegisterSystemSite(NewSystemSite())
}

// Model 站群ORM模型
func (s *sSystemSite) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SystemSite.Ctx(ctx), option...)
}

// List 获取站群列表
func (s *sSystemSite) List(ctx context.Context, in *sysin.SystemSiteListInp) (list []*sysin.SystemSiteListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.SystemSite.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SystemSite.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取站群数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.SystemSiteListModel{}).Page(in.Page, in.PerPage).OrderAsc(dao.SystemSite.Columns().Sort).OrderDesc(dao.SystemSite.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取站群列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出站群
func (s *sSystemSite) Export(ctx context.Context, in *sysin.SystemSiteListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.SystemSiteExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出站群-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.SystemSiteExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增站群
func (s *sSystemSite) Edit(ctx context.Context, in *sysin.SystemSiteEditInp) (err error) {
	// 验证'SiteName'唯一
	if err = hgorm.IsUnique(ctx, &dao.SystemSite, g.Map{dao.SystemSite.Columns().SiteName: in.SiteName}, "网站名称已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.SystemSiteUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改站群失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.SystemSiteInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增站群失败，请稍后重试！")
	}
	return
}

// Delete 删除站群
func (s *sSystemSite) Delete(ctx context.Context, in *sysin.SystemSiteDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除站群失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取站群最大排序
func (s *sSystemSite) MaxSort(ctx context.Context, in *sysin.SystemSiteMaxSortInp) (res *sysin.SystemSiteMaxSortModel, err error) {
	if err = dao.SystemSite.Ctx(ctx).Fields(dao.SystemSite.Columns().Sort).OrderDesc(dao.SystemSite.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取站群最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.SystemSiteMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取站群指定信息
func (s *sSystemSite) View(ctx context.Context, in *sysin.SystemSiteViewInp) (res *sysin.SystemSiteViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取站群信息，请稍后重试！")
		return
	}
	return
}
