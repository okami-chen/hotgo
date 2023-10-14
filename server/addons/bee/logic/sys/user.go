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

type sSysUser struct{}

func NewSysUser() *sSysUser {
	return &sSysUser{}
}

func init() {
	service.RegisterSysUser(NewSysUser())
}

// Model 用户信息ORM模型
func (s *sSysUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.User.Ctx(ctx), option...)
}

// List 获取用户信息列表
func (s *sSysUser) List(ctx context.Context, in *sysin.UserListInp) (list []*sysin.UserListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Uid > 0 {
		mod = mod.Where(dao.User.Columns().Uid, in.Uid)
	}

	// 查询状态
	if in.Status >= 0 {
		mod = mod.Where(dao.User.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.User.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 关联表info
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.User.Table(), dao.User.Columns().Uid, // 主表表名,关联字段
		dao.UserInfo.Table(), "info", dao.UserInfo.Columns().Uid, // 关联表表名,别名,关联字段
	)...)

	// 关联表social
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.User.Table(), dao.User.Columns().Uid, // 主表表名,关联字段
		dao.UserSocial.Table(), "social", dao.UserSocial.Columns().Uid, // 关联表表名,别名,关联字段
	)...)

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取用户信息数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	// 关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.UserListModel{}, &dao.User, []*hgorm.Join{
		{Dao: &dao.UserInfo, Alias: "info"},
		{Dao: &dao.UserSocial, Alias: "social"},
	})

	if err != nil {
		err = gerror.Wrap(err, "获取用户信息关联字段失败，请稍后重试！")
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.User.Columns().Uid).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取用户信息列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出用户信息
func (s *sSysUser) Export(ctx context.Context, in *sysin.UserListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.UserExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出用户信息-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.UserExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增用户信息
func (s *sSysUser) Edit(ctx context.Context, in *sysin.UserEditInp) (err error) {
	// 修改
	if in.Uid > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.UserUpdateFields{}).
			WherePri(in.Uid).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改用户信息失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.UserInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增用户信息失败，请稍后重试！")
	}
	return
}

// Delete 删除用户信息
func (s *sSysUser) Delete(ctx context.Context, in *sysin.UserDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Uid).Delete(); err != nil {
		err = gerror.Wrap(err, "删除用户信息失败，请稍后重试！")
		return
	}
	return
}

// View 获取用户信息指定信息
func (s *sSysUser) View(ctx context.Context, in *sysin.UserViewInp) (res *sysin.UserViewModel, err error) {
	mod := s.Model(ctx)
	//// 关联表info
	//mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
	//	dao.User.Table(), dao.User.Columns().Uid, // 主表表名,关联字段
	//	dao.UserInfo.Table(), "info", dao.UserInfo.Columns().Uid, // 关联表表名,别名,关联字段
	//)...)
	//
	//// 关联表select
	//fields, err := hgorm.GenJoinSelect(ctx, sysin.UserListModel{}, &dao.User, []*hgorm.Join{
	//	{Dao: &dao.UserInfo, Alias: "info"},
	//})

	if err = mod.WherePri(in.Uid).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取用户信息信息，请稍后重试！"+err.Error())
		return
	}

	if err = dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Uid, res.Uid).Scan(&res.Info); err != nil {
		err = gerror.Wrap(err, "获取用户信息信息，请稍后重试！")
		return
	}
	if err = dao.UserSocial.Ctx(ctx).Where(dao.UserSocial.Columns().Uid, res.Uid).Scan(&res.Social); err != nil {
		err = gerror.Wrap(err, "获取用户信息信息，请稍后重试！"+err.Error())
		return
	}

	return
}

// Status 更新用户信息状态
func (s *sSysUser) Status(ctx context.Context, in *sysin.UserStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Uid).Data(g.Map{
		dao.User.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新用户信息状态失败，请稍后重试！")
		return
	}
	return
}
