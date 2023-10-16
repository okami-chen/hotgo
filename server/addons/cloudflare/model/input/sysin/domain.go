// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sysin

import (
	"context"
	"hotgo/addons/cloudflare/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DomainUpdateFields 修改名字段过滤
type DomainUpdateFields struct {
	AccountId int64  `json:"account_id" dc:"关联账号"`
	DomainId  string `json:"domain_id"  dc:"域名编号"`
	Domain    string `json:"domain"     dc:"域名名称"`
}

// DomainInsertFields 新增名字段过滤
type DomainInsertFields struct {
	AccountId int64  `json:"account_id" dc:"关联账号"`
	DomainId  string `json:"domain_id"  dc:"域名编号"`
	Domain    string `json:"domain"     dc:"域名名称"`
}

// DomainEditInp 修改/新增名
type DomainEditInp struct {
	entity.Domain
}

func (in *DomainEditInp) Filter(ctx context.Context) (err error) {
	// 验证关联账号
	if err := g.Validator().Rules("required").Data(in.AccountId).Messages("关联账号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证域名编号
	if err := g.Validator().Rules("required").Data(in.DomainId).Messages("域名编号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证域名名称
	if err := g.Validator().Rules("required").Data(in.Domain).Messages("域名名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type DomainEditModel struct{}

// DomainDeleteInp 删除名
type DomainDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *DomainDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type DomainDeleteModel struct{}

// DomainViewInp 获取指定名信息
type DomainViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *DomainViewInp) Filter(ctx context.Context) (err error) {
	return
}

type DomainViewModel struct {
	entity.Domain
}

// DomainListInp 获取名列表
type DomainListInp struct {
	form.PageReq
	Domain string `json:"domain" dc:"域名名称"`
}

func (in *DomainListInp) Filter(ctx context.Context) (err error) {
	return
}

type DomainListModel struct {
	Id        int64       `json:"id"         dc:"自动编号"`
	AccountId int64       `json:"account_id" dc:"关联账号"`
	DomainId  string      `json:"domain_id"  dc:"域名编号"`
	Domain    string      `json:"domain"     dc:"域名名称"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
}

// DomainExportModel 导出名
type DomainExportModel struct {
	Id        int64       `json:"id"         dc:"自动编号"`
	AccountId int64       `json:"account_id" dc:"关联账号"`
	DomainId  string      `json:"domain_id"  dc:"域名编号"`
	Domain    string      `json:"domain"     dc:"域名名称"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" dc:"更新时间"`
}
