// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/card/model"
	"hotgo/addons/card/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
	ISysCard interface {
		// Model 卡片ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取卡片列表
		List(ctx context.Context, in *sysin.CardListInp) (list []*sysin.CardListModel, totalCount int, err error)
		// Export 导出卡片
		Export(ctx context.Context, in *sysin.CardListInp) (err error)
		// Edit 修改/新增卡片
		Edit(ctx context.Context, in *sysin.CardEditInp) (err error)
		// Delete 删除卡片
		Delete(ctx context.Context, in *sysin.CardDeleteInp) (err error)
		// View 获取卡片指定信息
		View(ctx context.Context, in *sysin.CardViewInp) (res *sysin.CardViewModel, err error)
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
	localSysCard   ISysCard
	localSysConfig ISysConfig
	localSysIndex  ISysIndex
)

func SysCard() ISysCard {
	if localSysCard == nil {
		panic("implement not found for interface ISysCard, forgot register?")
	}
	return localSysCard
}

func RegisterSysCard(i ISysCard) {
	localSysCard = i
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
