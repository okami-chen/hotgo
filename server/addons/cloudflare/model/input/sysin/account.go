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

// AccountUpdateFields 修改账号字段过滤
type AccountUpdateFields struct {
	Email  string `json:"email"  dc:"email"`
	Token  string `json:"token"  dc:"令牌"`
	Remark string `json:"remark" dc:"备注"`
}

// AccountInsertFields 新增账号字段过滤
type AccountInsertFields struct {
	Email  string `json:"email"  dc:"email"`
	Token  string `json:"token"  dc:"令牌"`
	Remark string `json:"remark" dc:"备注"`
}

// AccountEditInp 修改/新增账号
type AccountEditInp struct {
	entity.Account
}

func (in *AccountEditInp) Filter(ctx context.Context) (err error) {
	// 验证email
	if err := g.Validator().Rules("email").Data(in.Email).Messages("email不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证令牌
	if err := g.Validator().Rules("required").Data(in.Token).Messages("令牌不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type AccountEditModel struct{}

// AccountDeleteInp 删除账号
type AccountDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *AccountDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AccountDeleteModel struct{}

// AccountViewInp 获取指定账号信息
type AccountViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *AccountViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AccountViewModel struct {
	entity.Account
}

// AccountListInp 获取账号列表
type AccountListInp struct {
	form.PageReq
	Id    int64  `json:"id"    dc:"自动编号"`
	Email string `json:"email" dc:"email"`
}

func (in *AccountListInp) Filter(ctx context.Context) (err error) {
	return
}

type AccountListModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	Email     string      `json:"email"     dc:"email"`
	Token     string      `json:"token"     dc:"令牌"`
	Remark    string      `json:"remark"    dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// AccountExportModel 导出账号
type AccountExportModel struct {
	Id        int64       `json:"id"        dc:"自动编号"`
	Email     string      `json:"email"     dc:"email"`
	Token     string      `json:"token"     dc:"令牌"`
	Remark    string      `json:"remark"    dc:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
