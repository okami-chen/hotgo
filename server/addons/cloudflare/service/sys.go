// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/cloudflare/model"
	"hotgo/addons/cloudflare/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysAccount interface {
		// Model 账号ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取账号列表
		List(ctx context.Context, in *sysin.AccountListInp) (list []*sysin.AccountListModel, totalCount int, err error)
		// Export 导出账号
		Export(ctx context.Context, in *sysin.AccountListInp) (err error)
		// Edit 修改/新增账号
		Edit(ctx context.Context, in *sysin.AccountEditInp) (err error)
		// Delete 删除账号
		Delete(ctx context.Context, in *sysin.AccountDeleteInp) (err error)
		// View 获取账号指定信息
		View(ctx context.Context, in *sysin.AccountViewInp) (res *sysin.AccountViewModel, err error)
	}
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysDomain interface {
		// Model 名ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取名列表
		List(ctx context.Context, in *sysin.DomainListInp) (list []*sysin.DomainListModel, totalCount int, err error)
		// Export 导出名
		Export(ctx context.Context, in *sysin.DomainListInp) (err error)
		// Edit 修改/新增名
		Edit(ctx context.Context, in *sysin.DomainEditInp) (err error)
		// Delete 删除名
		Delete(ctx context.Context, in *sysin.DomainDeleteInp) (err error)
		// View 获取名指定信息
		View(ctx context.Context, in *sysin.DomainViewInp) (res *sysin.DomainViewModel, err error)
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
	ISysSsl interface {
		// Model 证书ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取证书列表
		List(ctx context.Context, in *sysin.SslListInp) (list []*sysin.SslListModel, totalCount int, err error)
		// Export 导出证书
		Export(ctx context.Context, in *sysin.SslListInp) (err error)
		// Edit 修改/新增证书
		Edit(ctx context.Context, in *sysin.SslEditInp) (err error)
		// Delete 删除证书
		Delete(ctx context.Context, in *sysin.SslDeleteInp) (err error)
		// View 获取证书指定信息
		View(ctx context.Context, in *sysin.SslViewInp) (res *sysin.SslViewModel, err error)
		// Status 更新证书状态
		Status(ctx context.Context, in *sysin.SslStatusInp) (err error)
	}
)

var (
	localSysAccount ISysAccount
	localSysConfig  ISysConfig
	localSysDomain  ISysDomain
	localSysIndex   ISysIndex
	localSysSsl     ISysSsl
)

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}

func SysSsl() ISysSsl {
	if localSysSsl == nil {
		panic("implement not found for interface ISysSsl, forgot register?")
	}
	return localSysSsl
}

func RegisterSysSsl(i ISysSsl) {
	localSysSsl = i
}

func SysAccount() ISysAccount {
	if localSysAccount == nil {
		panic("implement not found for interface ISysAccount, forgot register?")
	}
	return localSysAccount
}

func RegisterSysAccount(i ISysAccount) {
	localSysAccount = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysDomain() ISysDomain {
	if localSysDomain == nil {
		panic("implement not found for interface ISysDomain, forgot register?")
	}
	return localSysDomain
}

func RegisterSysDomain(i ISysDomain) {
	localSysDomain = i
}
