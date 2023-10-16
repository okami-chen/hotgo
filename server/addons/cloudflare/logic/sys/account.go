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

type sSysAccount struct{}

func NewSysAccount() *sSysAccount {
	return &sSysAccount{}
}

func init() {
	service.RegisterSysAccount(NewSysAccount())
}

// Model 账号ORM模型
func (s *sSysAccount) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Account.Ctx(ctx), option...)
}

// List 获取账号列表
func (s *sSysAccount) List(ctx context.Context, in *sysin.AccountListInp) (list []*sysin.AccountListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询自动编号
	if in.Id > 0 {
		mod = mod.Where(dao.Account.Columns().Id, in.Id)
	}

	// 查询email
	if in.Email != "" {
		mod = mod.WhereLike(dao.Account.Columns().Email, in.Email)
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取账号数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.AccountListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Account.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取账号列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出账号
func (s *sSysAccount) Export(ctx context.Context, in *sysin.AccountListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.AccountExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出账号-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.AccountExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增账号
func (s *sSysAccount) Edit(ctx context.Context, in *sysin.AccountEditInp) (err error) {
	// 验证'Email'唯一
	if err = hgorm.IsUnique(ctx, &dao.Account, g.Map{dao.Account.Columns().Email: in.Email}, "email已存在", in.Id); err != nil {
		return
	}
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.AccountUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改账号失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.AccountInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增账号失败，请稍后重试！")
	}
	return
}

// Delete 删除账号
func (s *sSysAccount) Delete(ctx context.Context, in *sysin.AccountDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除账号失败，请稍后重试！")
		return
	}
	return
}

// View 获取账号指定信息
func (s *sSysAccount) View(ctx context.Context, in *sysin.AccountViewInp) (res *sysin.AccountViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取账号信息，请稍后重试！")
		return
	}
	return
}
