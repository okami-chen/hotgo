// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/order/model"
	"hotgo/addons/order/model/input/sysin"
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
	ISysLogComplete interface {
		// Model 完成时间ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取完成时间列表
		List(ctx context.Context, in *sysin.LogCompleteListInp) (list []*sysin.LogCompleteListModel, totalCount int, err error)
		// Export 导出完成时间
		Export(ctx context.Context, in *sysin.LogCompleteListInp) (err error)
		// Edit 修改/新增完成时间
		Edit(ctx context.Context, in *sysin.LogCompleteEditInp) (err error)
		// Delete 删除完成时间
		Delete(ctx context.Context, in *sysin.LogCompleteDeleteInp) (err error)
		// View 获取完成时间指定信息
		View(ctx context.Context, in *sysin.LogCompleteViewInp) (res *sysin.LogCompleteViewModel, err error)
	}
)

var (
	localSysConfig      ISysConfig
	localSysIndex       ISysIndex
	localSysLogComplete ISysLogComplete
)

func SysLogComplete() ISysLogComplete {
	if localSysLogComplete == nil {
		panic("implement not found for interface ISysLogComplete, forgot register?")
	}
	return localSysLogComplete
}

func RegisterSysLogComplete(i ISysLogComplete) {
	localSysLogComplete = i
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

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}
