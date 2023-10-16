// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
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

// CurrencyUpdateFields 修改货币字段过滤
type CurrencyUpdateFields struct {
	Status int     `json:"status" dc:"状态"`
	Code   string  `json:"code"   dc:"代码"`
	Name   string  `json:"name"   dc:"名称"`
	Symbol string  `json:"symbol" dc:"符号"`
	Rate   float64 `json:"rate"   dc:"汇率"`
}

// CurrencyInsertFields 新增货币字段过滤
type CurrencyInsertFields struct {
	Status int     `json:"status" dc:"状态"`
	Code   string  `json:"code"   dc:"代码"`
	Name   string  `json:"name"   dc:"名称"`
	Symbol string  `json:"symbol" dc:"符号"`
	Rate   float64 `json:"rate"   dc:"汇率"`
}

// CurrencyEditInp 修改/新增货币
type CurrencyEditInp struct {
	entity.Currency
}

func (in *CurrencyEditInp) Filter(ctx context.Context) (err error) {
	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:0,1").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证代码
	if err := g.Validator().Rules("required").Data(in.Code).Messages("代码不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证符号
	if err := g.Validator().Rules("required").Data(in.Symbol).Messages("符号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证汇率
	if err := g.Validator().Rules("required").Data(in.Rate).Messages("汇率不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type CurrencyEditModel struct{}

// CurrencyDeleteInp 删除货币
type CurrencyDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *CurrencyDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CurrencyDeleteModel struct{}

// CurrencyViewInp 获取指定货币信息
type CurrencyViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *CurrencyViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CurrencyViewModel struct {
	entity.Currency
}

// CurrencyListInp 获取货币列表
type CurrencyListInp struct {
	form.PageReq
	Id     int64 `json:"id"     dc:"自动编号"`
	Status int   `json:"status" dc:"状态"`
}

func (in *CurrencyListInp) Filter(ctx context.Context) (err error) {
	return
}

type CurrencyListModel struct {
	Id     int64   `json:"id"     dc:"自动编号"`
	Status int     `json:"status" dc:"状态"`
	Code   string  `json:"code"   dc:"代码"`
	Name   string  `json:"name"   dc:"名称"`
	Symbol string  `json:"symbol" dc:"符号"`
	Rate   float64 `json:"rate"   dc:"汇率"`
}

// CurrencyExportModel 导出货币
type CurrencyExportModel struct {
	Id        int64       `json:"id"         dc:"自动编号"`
	Status    int         `json:"status"     dc:"状态"`
	Code      string      `json:"code"       dc:"代码"`
	Name      string      `json:"name"       dc:"名称"`
	Symbol    string      `json:"symbol"     dc:"符号"`
	Rate      float64     `json:"rate"       dc:"汇率"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}

// CurrencyStatusInp 更新货币状态
type CurrencyStatusInp struct {
	Id     int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
	Status int   `json:"status" dc:"状态"`
}

func (in *CurrencyStatusInp) Filter(ctx context.Context) (err error) {
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

type CurrencyStatusModel struct{}
