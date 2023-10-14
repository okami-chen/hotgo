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

type sSysDevice struct{}

func NewSysDevice() *sSysDevice {
	return &sSysDevice{}
}

func init() {
	service.RegisterSysDevice(NewSysDevice())
}

// Model 设备ORM模型
func (s *sSysDevice) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Device.Ctx(ctx), option...)
}

// List 获取设备列表
func (s *sSysDevice) List(ctx context.Context, in *sysin.DeviceListInp) (list []*sysin.DeviceListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.DeviceId > 0 {
		mod = mod.Where(dao.Device.Columns().DeviceId, in.DeviceId)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.Device.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Device.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取设备数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.DeviceListModel{}).Page(in.Page, in.PerPage).OrderAsc(dao.Device.Columns().Sort).OrderDesc(dao.Device.Columns().DeviceId).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取设备列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出设备
func (s *sSysDevice) Export(ctx context.Context, in *sysin.DeviceListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.DeviceExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出设备-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.DeviceExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增设备
func (s *sSysDevice) Edit(ctx context.Context, in *sysin.DeviceEditInp) (err error) {
	// 修改
	if in.DeviceId > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.DeviceUpdateFields{}).
			WherePri(in.DeviceId).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改设备失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.DeviceInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增设备失败，请稍后重试！")
	}
	return
}

// Delete 删除设备
func (s *sSysDevice) Delete(ctx context.Context, in *sysin.DeviceDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.DeviceId).Delete(); err != nil {
		err = gerror.Wrap(err, "删除设备失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取设备最大排序
func (s *sSysDevice) MaxSort(ctx context.Context, in *sysin.DeviceMaxSortInp) (res *sysin.DeviceMaxSortModel, err error) {
	if err = dao.Device.Ctx(ctx).Fields(dao.Device.Columns().Sort).OrderDesc(dao.Device.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取设备最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.DeviceMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取设备指定信息
func (s *sSysDevice) View(ctx context.Context, in *sysin.DeviceViewInp) (res *sysin.DeviceViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.DeviceId).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取设备信息，请稍后重试！")
		return
	}
	return
}

// Status 更新设备状态
func (s *sSysDevice) Status(ctx context.Context, in *sysin.DeviceStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.DeviceId).Data(g.Map{
		dao.Device.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新设备状态失败，请稍后重试！")
		return
	}
	return
}
