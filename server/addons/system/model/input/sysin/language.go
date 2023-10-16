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
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LanguageUpdateFields 修改语言字段过滤
type LanguageUpdateFields struct {
	Language   string `json:"language"    dc:"语言英文"`
	LanguageZh string `json:"language_zh" dc:"语言中文"`
	Code       string `json:"code"        dc:"语言简称"`
	Image      string `json:"image"       dc:"国旗图片"`
}

// LanguageInsertFields 新增语言字段过滤
type LanguageInsertFields struct {
	Language   string `json:"language"    dc:"语言英文"`
	LanguageZh string `json:"language_zh" dc:"语言中文"`
	Code       string `json:"code"        dc:"语言简称"`
	Image      string `json:"image"       dc:"国旗图片"`
}

// LanguageEditInp 修改/新增语言
type LanguageEditInp struct {
	entity.Language
}

func (in *LanguageEditInp) Filter(ctx context.Context) (err error) {
	// 验证语言英文
	if err := g.Validator().Rules("required").Data(in.Language).Messages("语言英文不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证语言中文
	if err := g.Validator().Rules("required").Data(in.LanguageZh).Messages("语言中文不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证语言简称
	if err := g.Validator().Rules("required").Data(in.Code).Messages("语言简称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type LanguageEditModel struct{}

// LanguageDeleteInp 删除语言
type LanguageDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *LanguageDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type LanguageDeleteModel struct{}

// LanguageViewInp 获取指定语言信息
type LanguageViewInp struct {
	Id int `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *LanguageViewInp) Filter(ctx context.Context) (err error) {
	return
}

type LanguageViewModel struct {
	entity.Language
}

// LanguageListInp 获取语言列表
type LanguageListInp struct {
	form.PageReq
	Id         int    `json:"id"          dc:"自动编号"`
	Language   string `json:"language"    dc:"语言英文"`
	LanguageZh string `json:"language_zh" dc:"语言中文"`
	Code       string `json:"code"        dc:"语言简称"`
}

func (in *LanguageListInp) Filter(ctx context.Context) (err error) {
	return
}

type LanguageListModel struct {
	Id         int         `json:"id"          dc:"自动编号"`
	Language   string      `json:"language"    dc:"语言英文"`
	LanguageZh string      `json:"language_zh" dc:"语言中文"`
	Code       string      `json:"code"        dc:"语言简称"`
	CreatedAt  *gtime.Time `json:"created_at"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updated_at"  dc:"更新时间"`
}

// LanguageExportModel 导出语言
type LanguageExportModel struct {
	Id         int         `json:"id"          dc:"自动编号"`
	Language   string      `json:"language"    dc:"语言英文"`
	LanguageZh string      `json:"language_zh" dc:"语言中文"`
	Code       string      `json:"code"        dc:"语言简称"`
	CreatedAt  *gtime.Time `json:"created_at"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updated_at"  dc:"更新时间"`
}
