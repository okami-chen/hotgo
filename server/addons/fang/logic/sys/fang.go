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

type sSysFang struct{}

func NewSysFang() *sSysFang {
	return &sSysFang{}
}

func init() {
	service.RegisterSysFang(NewSysFang())
}

// Model 租房ORM模型
func (s *sSysFang) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Fang.Ctx(ctx), option...)
}

// List 获取租房列表
func (s *sSysFang) List(ctx context.Context, in *sysin.FangListInp) (list []*sysin.FangListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询编号
	if in.Sid != "" {
		mod = mod.Where(dao.Fang.Columns().Sid, in.Sid)
	}

	// 查询小区
	if in.Village != "" {
		mod = mod.Where(dao.Fang.Columns().Village, in.Village)
	}

	// 查询标题
	if in.Title != "" {
		mod = mod.WhereLike(dao.Fang.Columns().Title, "%"+in.Title+"%")
	}

	// 查询户型
	if in.HouseType != "" {
		mod = mod.Where(dao.Fang.Columns().HouseType, in.HouseType)
	}

	// 查询朝向
	if in.ToWard != "" {
		mod = mod.Where(dao.Fang.Columns().ToWard, in.ToWard)
	}

	// 查询旗帜
	if in.Flag > 0 {
		mod = mod.Where(dao.Fang.Columns().Flag, in.Flag)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取租房数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.FangListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Fang.Columns().CreatedAt).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取租房列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出租房
func (s *sSysFang) Export(ctx context.Context, in *sysin.FangListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.FangExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出租房-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.FangExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增租房
func (s *sSysFang) Edit(ctx context.Context, in *sysin.FangEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.FangUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改租房失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.FangInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增租房失败，请稍后重试！")
	}
	return
}

// Delete 删除租房
func (s *sSysFang) Delete(ctx context.Context, in *sysin.FangDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除租房失败，请稍后重试！")
		return
	}
	return
}

// View 获取租房指定信息
func (s *sSysFang) View(ctx context.Context, in *sysin.FangViewInp) (res *sysin.FangViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取租房信息，请稍后重试！")
		return
	}
	return
}
