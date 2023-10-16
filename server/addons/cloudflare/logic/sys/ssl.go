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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysSsl struct{}

func NewSysSsl() *sSysSsl {
	return &sSysSsl{}
}

func init() {
	service.RegisterSysSsl(NewSysSsl())
}

// Model 证书ORM模型
func (s *sSysSsl) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.Ssl.Ctx(ctx), option...)
}

// List 获取证书列表
func (s *sSysSsl) List(ctx context.Context, in *sysin.SslListInp) (list []*sysin.SslListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询编号
	if in.Id > 0 {
		mod = mod.Where(dao.Ssl.Columns().Id, in.Id)
	}

	// 查询域名
	if in.Domain != "" {
		mod = mod.WhereLike(dao.Ssl.Columns().Domain, in.Domain)
	}

	// 查询状态
	if in.Status != "" {
		mod = mod.Where(dao.Ssl.Columns().Status, in.Status)
	}

	// 查询类型
	if in.Type != "" {
		mod = mod.Where(dao.Ssl.Columns().Type, in.Type)
	}

	// 查询签发组织
	if in.Issuer != "" {
		mod = mod.Where(dao.Ssl.Columns().Issuer, in.Issuer)
	}

	// 查询授权组织
	if in.Authority != "" {
		mod = mod.Where(dao.Ssl.Columns().Authority, in.Authority)
	}

	// 查询签名方式
	if in.Signature != "" {
		mod = mod.Where(dao.Ssl.Columns().Signature, in.Signature)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.Ssl.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取证书数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.SslListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.Ssl.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取证书列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出证书
func (s *sSysSsl) Export(ctx context.Context, in *sysin.SslListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.SslExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出证书-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.SslExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增证书
func (s *sSysSsl) Edit(ctx context.Context, in *sysin.SslEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.SslUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改证书失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.SslInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增证书失败，请稍后重试！")
	}
	return
}

// Delete 删除证书
func (s *sSysSsl) Delete(ctx context.Context, in *sysin.SslDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除证书失败，请稍后重试！")
		return
	}
	return
}

// View 获取证书指定信息
func (s *sSysSsl) View(ctx context.Context, in *sysin.SslViewInp) (res *sysin.SslViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取证书信息，请稍后重试！")
		return
	}
	return
}

// Status 更新证书状态
func (s *sSysSsl) Status(ctx context.Context, in *sysin.SslStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.Ssl.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新证书状态失败，请稍后重试！")
		return
	}
	return
}
