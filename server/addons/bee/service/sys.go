// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/bee/model"
	"hotgo/addons/bee/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
	ISystemSite interface {
		// Model 站群ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取站群列表
		List(ctx context.Context, in *sysin.SystemSiteListInp) (list []*sysin.SystemSiteListModel, totalCount int, err error)
		// Export 导出站群
		Export(ctx context.Context, in *sysin.SystemSiteListInp) (err error)
		// Edit 修改/新增站群
		Edit(ctx context.Context, in *sysin.SystemSiteEditInp) (err error)
		// Delete 删除站群
		Delete(ctx context.Context, in *sysin.SystemSiteDeleteInp) (err error)
		// MaxSort 获取站群最大排序
		MaxSort(ctx context.Context, in *sysin.SystemSiteMaxSortInp) (res *sysin.SystemSiteMaxSortModel, err error)
		// View 获取站群指定信息
		View(ctx context.Context, in *sysin.SystemSiteViewInp) (res *sysin.SystemSiteViewModel, err error)
	}

)

var (
	localSystemSite  ISystemSite
	localSysConfig   ISysConfig
	localSysIndex    ISysIndex
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}

func SystemSite() ISystemSite {
	if localSystemSite == nil {
		panic("implement not found for interface ISystemSite, forgot register?")
	}
	return localSystemSite
}

func RegisterSystemSite(i ISystemSite) {
	localSystemSite = i
}

