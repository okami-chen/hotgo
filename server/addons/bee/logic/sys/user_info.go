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
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysUserInfo struct{}

func NewSysUserInfo() *sSysUserInfo {
	return &sSysUserInfo{}
}

func init() {
	service.RegisterSysUserInfo(NewSysUserInfo())
}

// Model 客户信息扩展ORM模型
func (s *sSysUserInfo) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.UserInfo.Ctx(ctx), option...)
}

// List 获取客户信息扩展列表
func (s *sSysUserInfo) List(ctx context.Context, in *sysin.UserInfoListInp) (list []*sysin.UserInfoListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.UserInfo.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.UserInfo.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取客户信息扩展数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.UserInfoListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.UserInfo.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取客户信息扩展列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出客户信息扩展
func (s *sSysUserInfo) Export(ctx context.Context, in *sysin.UserInfoListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.UserInfoExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出客户信息扩展-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.UserInfoExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增客户信息扩展
func (s *sSysUserInfo) Edit(ctx context.Context, in *sysin.UserInfoEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.UserInfoUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改客户信息扩展失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.UserInfoInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增客户信息扩展失败，请稍后重试！")
	}
	return
}

// Delete 删除客户信息扩展
func (s *sSysUserInfo) Delete(ctx context.Context, in *sysin.UserInfoDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除客户信息扩展失败，请稍后重试！")
		return
	}
	return
}

// View 获取客户信息扩展指定信息
func (s *sSysUserInfo) View(ctx context.Context, in *sysin.UserInfoViewInp) (res *sysin.UserInfoViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取客户信息扩展信息，请稍后重试！")
		return
	}
	return
}
