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

// CountryUpdateFields 修改国家字段过滤
type CountryUpdateFields struct {
	Country      string `json:"country"       dc:"国家缩写"`
	CountryName  string `json:"country_name"  dc:"国家名称全程"`
	DiallingCode string `json:"dialling_code" dc:"国家区号"`
	Img          string `json:"img"           dc:"国旗图片"`
}

// CountryInsertFields 新增国家字段过滤
type CountryInsertFields struct {
	Country      string `json:"country"       dc:"国家缩写"`
	CountryName  string `json:"country_name"  dc:"国家名称全程"`
	DiallingCode string `json:"dialling_code" dc:"国家区号"`
	Img          string `json:"img"           dc:"国旗图片"`
}

// CountryEditInp 修改/新增国家
type CountryEditInp struct {
	entity.Country
}

func (in *CountryEditInp) Filter(ctx context.Context) (err error) {
	// 验证国家缩写
	if err := g.Validator().Rules("required").Data(in.Country).Messages("国家缩写不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证国家名称全程
	if err := g.Validator().Rules("required").Data(in.CountryName).Messages("国家名称全程不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证国家区号
	if err := g.Validator().Rules("required").Data(in.DiallingCode).Messages("国家区号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证国旗图片
	if err := g.Validator().Rules("required").Data(in.Img).Messages("国旗图片不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type CountryEditModel struct{}

// CountryDeleteInp 删除国家
type CountryDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *CountryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CountryDeleteModel struct{}

// CountryViewInp 获取指定国家信息
type CountryViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *CountryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CountryViewModel struct {
	entity.Country
}

// CountryListInp 获取国家列表
type CountryListInp struct {
	form.PageReq
	Country      string `json:"country"       dc:"国家缩写"`
	CountryName  string `json:"country_name"  dc:"国家名称全程"`
	DiallingCode string `json:"dialling_code" dc:"国家区号"`
}

func (in *CountryListInp) Filter(ctx context.Context) (err error) {
	return
}

type CountryListModel struct {
	Id           int64  `json:"id"            dc:"自动编号"`
	Country      string `json:"country"       dc:"国家缩写"`
	CountryName  string `json:"country_name"  dc:"国家名称全程"`
	DiallingCode string `json:"dialling_code" dc:"国家区号"`
	Img          string `json:"img"           dc:"国旗图片"`
}

// CountryExportModel 导出国家
type CountryExportModel struct {
	Id           int64       `json:"id"            dc:"自动编号"`
	Country      string      `json:"country"       dc:"国家缩写"`
	CountryName  string      `json:"country_name"  dc:"国家名称全程"`
	DiallingCode string      `json:"dialling_code" dc:"国家区号"`
	Img          string      `json:"img"           dc:"国旗图片"`
	CreatedAt    *gtime.Time `json:"created_at"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updated_at"    dc:"更新时间"`
}
