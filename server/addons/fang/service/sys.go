// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/fang/model"
	"hotgo/addons/fang/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysFang interface {
		// Model 租房ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取租房列表
		List(ctx context.Context, in *sysin.FangListInp) (list []*sysin.FangListModel, totalCount int, err error)
		// Export 导出租房
		Export(ctx context.Context, in *sysin.FangListInp) (err error)
		// Edit 修改/新增租房
		Edit(ctx context.Context, in *sysin.FangEditInp) (err error)
		// Delete 删除租房
		Delete(ctx context.Context, in *sysin.FangDeleteInp) (err error)
		// View 获取租房指定信息
		View(ctx context.Context, in *sysin.FangViewInp) (res *sysin.FangViewModel, err error)
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
)

var (
	localSysFang   ISysFang
	localSysIndex  ISysIndex
	localSysConfig ISysConfig
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

func SysFang() ISysFang {
	if localSysFang == nil {
		panic("implement not found for interface ISysFang, forgot register?")
	}
	return localSysFang
}

func RegisterSysFang(i ISysFang) {
	localSysFang = i
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
