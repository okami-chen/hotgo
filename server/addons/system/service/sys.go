// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/system/model"
	"hotgo/addons/system/model/input/sysin"
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
	ISysEvent interface {
		// Model 事件ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取事件列表
		List(ctx context.Context, in *sysin.EventListInp) (list []*sysin.EventListModel, totalCount int, err error)
		// Export 导出事件
		Export(ctx context.Context, in *sysin.EventListInp) (err error)
		// Edit 修改/新增事件
		Edit(ctx context.Context, in *sysin.EventEditInp) (err error)
		// Delete 删除事件
		Delete(ctx context.Context, in *sysin.EventDeleteInp) (err error)
		// View 获取事件指定信息
		View(ctx context.Context, in *sysin.EventViewInp) (res *sysin.EventViewModel, err error)
		// Status 更新事件状态
		Status(ctx context.Context, in *sysin.EventStatusInp) (err error)
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
)

var (
	localSysEvent  ISysEvent
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

func SysEvent() ISysEvent {
	if localSysEvent == nil {
		panic("implement not found for interface ISysEvent, forgot register?")
	}
	return localSysEvent
}

func RegisterSysEvent(i ISysEvent) {
	localSysEvent = i
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
