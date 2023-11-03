// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package sysin

import (
	"context"
	"hotgo/addons/fang/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VillageUpdateFields 修改小区字段过滤
type VillageUpdateFields struct {
	Name  string `json:"name"   dc:"名称"`
	Data1 string `json:"data_1" dc:"数据_1"`
}

// VillageInsertFields 新增小区字段过滤
type VillageInsertFields struct {
	Name  string `json:"name"   dc:"名称"`
	Data1 string `json:"data_1" dc:"数据_1"`
}

// VillageEditInp 修改/新增小区
type VillageEditInp struct {
	entity.Village
}

func (in *VillageEditInp) Filter(ctx context.Context) (err error) {
	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:华安家园,美晨雅阁,金宏豪庭,元一名城A区,元一名城B区,元一名城C区,上城国际丁香苑,上城国际玫瑰苑,金域蓝湾,都市清华,荷塘月色,北美印象雅苑,金都华庭二期,金都华庭三期,金都华庭五期").Data(in.Name).Messages("名称值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type VillageEditModel struct{}

// VillageDeleteInp 删除小区
type VillageDeleteInp struct {
	Id interface{} `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *VillageDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type VillageDeleteModel struct{}

// VillageViewInp 获取指定小区信息
type VillageViewInp struct {
	Id int64 `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *VillageViewInp) Filter(ctx context.Context) (err error) {
	return
}

type VillageViewModel struct {
	entity.Village
}

// VillageListInp 获取小区列表
type VillageListInp struct {
	form.PageReq
	Name string `json:"name" dc:"名称"`
}

func (in *VillageListInp) Filter(ctx context.Context) (err error) {
	return
}

type VillageListModel struct {
	Id        int64       `json:"id"         dc:"编号"`
	Name      string      `json:"name"       dc:"名称"`
	Data1     string      `json:"data_1"     dc:"数据_1"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}

// VillageExportModel 导出小区
type VillageExportModel struct {
	Id        int64       `json:"id"         dc:"编号"`
	Name      string      `json:"name"       dc:"名称"`
	Data1     string      `json:"data_1"     dc:"数据_1"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}
