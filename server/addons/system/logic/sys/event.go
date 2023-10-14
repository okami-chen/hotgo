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
	"hotgo/addons/system/dao"
	"hotgo/addons/system/model/input/sysin"
	"hotgo/addons/system/service"
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

type sSysEvent struct{}

func NewSysEvent() *sSysEvent {
	return &sSysEvent{}
}

func init() {
	service.RegisterSysEvent(NewSysEvent())
}

// Model 事件ORM模型
func (s *sSysEvent) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SystemEvent.Ctx(ctx), option...)
}

// List 获取事件列表
func (s *sSysEvent) List(ctx context.Context, in *sysin.EventListInp) (list []*sysin.EventListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.SystemEvent.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.SystemEvent.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.SystemEvent.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取事件数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.EventListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.SystemEvent.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取事件列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出事件
func (s *sSysEvent) Export(ctx context.Context, in *sysin.EventListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.EventExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出事件-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.EventExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增事件
func (s *sSysEvent) Edit(ctx context.Context, in *sysin.EventEditInp) (err error) {
	// 验证'Name'唯一
	if err = hgorm.IsUnique(ctx, &dao.SystemEvent, g.Map{dao.SystemEvent.Columns().Name: in.Name}, "表名已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.EventUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改事件失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.EventInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增事件失败，请稍后重试！")
	}
	return
}

// Delete 删除事件
func (s *sSysEvent) Delete(ctx context.Context, in *sysin.EventDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除事件失败，请稍后重试！")
		return
	}
	return
}

// View 获取事件指定信息
func (s *sSysEvent) View(ctx context.Context, in *sysin.EventViewInp) (res *sysin.EventViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取事件信息，请稍后重试！")
		return
	}
	return
}

// Status 更新事件状态
func (s *sSysEvent) Status(ctx context.Context, in *sysin.EventStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SystemEvent.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新事件状态失败，请稍后重试！")
		return
	}
	return
}
