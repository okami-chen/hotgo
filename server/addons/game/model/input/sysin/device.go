// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sysin

import (
	"context"
	"hotgo/addons/game/model/entity"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceUpdateFields 修改设备字段过滤
type DeviceUpdateFields struct {
	Title  string `json:"title"  dc:"分类名称"`
	Image  string `json:"image"  dc:"图片"`
	Status int    `json:"status" dc:"状态"`
	Sort   int64  `json:"sort"   dc:"排序"`
	Url    string `json:"url"    dc:"网址"`
}

// DeviceInsertFields 新增设备字段过滤
type DeviceInsertFields struct {
	Title  string `json:"title"  dc:"分类名称"`
	Image  string `json:"image"  dc:"图片"`
	Status int    `json:"status" dc:"状态"`
	Sort   int64  `json:"sort"   dc:"排序"`
	Url    string `json:"url"    dc:"网址"`
}

// DeviceEditInp 修改/新增设备
type DeviceEditInp struct {
	entity.Device
}

func (in *DeviceEditInp) Filter(ctx context.Context) (err error) {
	// 验证分类名称
	if err := g.Validator().Rules("required").Data(in.Title).Messages("分类名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证网址
	if err := g.Validator().Rules("required").Data(in.Url).Messages("网址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type DeviceEditModel struct{}

// DeviceDeleteInp 删除设备
type DeviceDeleteInp struct {
	DeviceId interface{} `json:"deviceId" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *DeviceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type DeviceDeleteModel struct{}

// DeviceViewInp 获取指定设备信息
type DeviceViewInp struct {
	DeviceId int64 `json:"deviceId" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *DeviceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type DeviceViewModel struct {
	entity.Device
}

// DeviceListInp 获取设备列表
type DeviceListInp struct {
	form.PageReq
	DeviceId  int64         `json:"deviceId"  dc:"自动编号"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *DeviceListInp) Filter(ctx context.Context) (err error) {
	return
}

type DeviceListModel struct {
	DeviceId  int64       `json:"deviceId"  dc:"自动编号"`
	Title     string      `json:"title"     dc:"分类名称"`
	Image     string      `json:"image"     dc:"图片"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int64       `json:"sort"      dc:"排序"`
	Url       string      `json:"url"       dc:"网址"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// DeviceExportModel 导出设备
type DeviceExportModel struct {
	DeviceId  int64       `json:"deviceId"  dc:"自动编号"`
	Title     string      `json:"title"     dc:"分类名称"`
	Image     string      `json:"image"     dc:"图片"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int64       `json:"sort"      dc:"排序"`
	Url       string      `json:"url"       dc:"网址"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// DeviceMaxSortInp 获取设备最大排序
type DeviceMaxSortInp struct{}

func (in *DeviceMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type DeviceMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// DeviceStatusInp 更新设备状态
type DeviceStatusInp struct {
	DeviceId int64 `json:"deviceId" v:"required#自动编号不能为空" dc:"自动编号"`
	Status   int   `json:"status" dc:"状态"`
}

func (in *DeviceStatusInp) Filter(ctx context.Context) (err error) {
	if in.DeviceId <= 0 {
		err = gerror.New("自动编号不能为空")
		return
	}

	if in.Status < 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type DeviceStatusModel struct{}
