// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/user/model"
	"hotgo/addons/user/model/input/sysin"
)

type (
	ISysConfig interface {
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysIndex interface {
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
)

var (
	localSysConfig ISysConfig
	localSysIndex  ISysIndex
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
