// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sysin

import (
	"context"
	"hotgo/addons/card/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CardUpdateFields 修改卡片字段过滤
type CardUpdateFields struct {
	Name     string      `json:"name"     dc:"持卡"`
	Title    string      `json:"title"    dc:"名称"`
	Bank     string      `json:"bank"     dc:"银行"`
	Organize string      `json:"organize" dc:"组织"`
	CardNo   string      `json:"cardNo"   dc:"卡号"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"过期时间"`
	Code     string      `json:"code"     dc:"识别码"`
	Remark   string      `json:"remark"   dc:"备注"`
}

// CardInsertFields 新增卡片字段过滤
type CardInsertFields struct {
	Name     string      `json:"name"     dc:"持卡"`
	Title    string      `json:"title"    dc:"名称"`
	Bank     string      `json:"bank"     dc:"银行"`
	Organize string      `json:"organize" dc:"组织"`
	CardNo   string      `json:"cardNo"   dc:"卡号"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"过期时间"`
	Code     string      `json:"code"     dc:"识别码"`
	Remark   string      `json:"remark"   dc:"备注"`
}

// CardEditInp 修改/新增卡片
type CardEditInp struct {
	entity.Card
}

func (in *CardEditInp) Filter(ctx context.Context) (err error) {
	// 验证持卡
	if err := g.Validator().Rules("required").Data(in.Name).Messages("持卡不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Title).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证银行
	if err := g.Validator().Rules("required").Data(in.Bank).Messages("银行不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证组织
	if err := g.Validator().Rules("required").Data(in.Organize).Messages("组织不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证卡号
	if err := g.Validator().Rules("required").Data(in.CardNo).Messages("卡号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证过期时间
	if err := g.Validator().Rules("required").Data(in.ExpireAt).Messages("过期时间不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证识别码
	if err := g.Validator().Rules("required").Data(in.Code).Messages("识别码不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type CardEditModel struct{}

// CardDeleteInp 删除卡片
type CardDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *CardDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CardDeleteModel struct{}

// CardViewInp 获取指定卡片信息
type CardViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *CardViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CardViewModel struct {
	entity.Card
}

// CardListInp 获取卡片列表
type CardListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自动编号"`
	CardNo    string        `json:"card_no"	dc:"卡号"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"created_at"`
}

func (in *CardListInp) Filter(ctx context.Context) (err error) {
	return
}

type CardListModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	Name      string      `json:"name"      dc:"持卡"`
	Title     string      `json:"title"     dc:"名称"`
	Bank      string      `json:"bank"      dc:"银行"`
	Organize  string      `json:"organize"  dc:"组织"`
	CardNo    string      `json:"cardNo"    dc:"卡号"`
	ExpireAt  *gtime.Time `json:"expireAt"  dc:"过期时间"`
	Code      string      `json:"code"      dc:"识别码"`
	Remark    string      `json:"remark"    dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}

// CardExportModel 导出卡片
type CardExportModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	Name      string      `json:"name"      dc:"持卡"`
	Title     string      `json:"title"     dc:"名称"`
	Bank      string      `json:"bank"      dc:"银行"`
	Organize  string      `json:"organize"  dc:"组织"`
	CardNo    string      `json:"cardNo"    dc:"卡号"`
	ExpireAt  *gtime.Time `json:"expireAt"  dc:"过期时间"`
	Code      string      `json:"code"      dc:"识别码"`
	Remark    string      `json:"remark"    dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}
