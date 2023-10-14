// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sysin

import (
	"context"
	"hotgo/addons/system/model/entity"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EventUpdateFields 修改事件字段过滤
type EventUpdateFields struct {
	Status int    `json:"status" dc:"状态"`
	Name   string `json:"name"   dc:"表名"`
	Pk     string `json:"pk"     dc:"主键"`
	Event  string `json:"event"  dc:"事件"`
}

// EventInsertFields 新增事件字段过滤
type EventInsertFields struct {
	Status int    `json:"status" dc:"状态"`
	Name   string `json:"name"   dc:"表名"`
	Pk     string `json:"pk"     dc:"主键"`
	Event  string `json:"event"  dc:"事件"`
}

// EventEditInp 修改/新增事件
type EventEditInp struct {
	entity.SystemEvent
}

func (in *EventEditInp) Filter(ctx context.Context) (err error) {
	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证表名
	if err := g.Validator().Rules("required").Data(in.Name).Messages("表名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证主键
	if err := g.Validator().Rules("required").Data(in.Pk).Messages("主键不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证事件
	if err := g.Validator().Rules("required").Data(in.Event).Messages("事件不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type EventEditModel struct{}

// EventDeleteInp 删除事件
type EventDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *EventDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type EventDeleteModel struct{}

// EventViewInp 获取指定事件信息
type EventViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *EventViewInp) Filter(ctx context.Context) (err error) {
	return
}

type EventViewModel struct {
	entity.SystemEvent
}

// EventListInp 获取事件列表
type EventListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自动编号"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *EventListInp) Filter(ctx context.Context) (err error) {
	return
}

type EventListModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	Status    int         `json:"status"    dc:"状态"`
	Name      string      `json:"name"      dc:"表名"`
	Pk        string      `json:"pk"        dc:"主键"`
	Event     string      `json:"event"     dc:"事件"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// EventExportModel 导出事件
type EventExportModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	Status    int         `json:"status"    dc:"状态"`
	Name      string      `json:"name"      dc:"表名"`
	Pk        string      `json:"pk"        dc:"主键"`
	Event     string      `json:"event"     dc:"事件"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// EventStatusInp 更新事件状态
type EventStatusInp struct {
	Id     int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
	Status int   `json:"status" dc:"状态"`
}

func (in *EventStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
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

type EventStatusModel struct{}
