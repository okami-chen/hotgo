// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sysin

import (
	"context"
	"hotgo/addons/order/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LogCompleteUpdateFields 修改完成时间字段过滤
type LogCompleteUpdateFields struct {
	OrderId int64 `json:"orderId" dc:"订单编号"`
	AdminId int64 `json:"adminId" dc:"操作人员"`
}

// LogCompleteInsertFields 新增完成时间字段过滤
type LogCompleteInsertFields struct {
	OrderId int64 `json:"orderId" dc:"订单编号"`
	AdminId int64 `json:"adminId" dc:"操作人员"`
}

// LogCompleteEditInp 修改/新增完成时间
type LogCompleteEditInp struct {
	entity.OrderTimeComplete
}

func (in *LogCompleteEditInp) Filter(ctx context.Context) (err error) {
	// 验证订单编号
	if err := g.Validator().Rules("required").Data(in.OrderId).Messages("订单编号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证操作人员
	if err := g.Validator().Rules("required").Data(in.AdminId).Messages("操作人员不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type LogCompleteEditModel struct{}

// LogCompleteDeleteInp 删除完成时间
type LogCompleteDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *LogCompleteDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type LogCompleteDeleteModel struct{}

// LogCompleteViewInp 获取指定完成时间信息
type LogCompleteViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *LogCompleteViewInp) Filter(ctx context.Context) (err error) {
	return
}

type LogCompleteViewModel struct {
	entity.OrderTimeComplete
}

// LogCompleteListInp 获取完成时间列表
type LogCompleteListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自动编号"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *LogCompleteListInp) Filter(ctx context.Context) (err error) {
	return
}

type LogCompleteListModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	OrderId   int64       `json:"orderId"   dc:"订单编号"`
	AdminId   int64       `json:"adminId"   dc:"操作人员"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// LogCompleteExportModel 导出完成时间
type LogCompleteExportModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	OrderId   int64       `json:"orderId"   dc:"订单编号"`
	AdminId   int64       `json:"adminId"   dc:"操作人员"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
