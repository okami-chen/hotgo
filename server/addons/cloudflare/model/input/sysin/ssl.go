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
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SslUpdateFields 修改证书字段过滤
type SslUpdateFields struct {
	AccountId int64       `json:"account_id"  dc:"账号"`
	Domain    string      `json:"domain"      dc:"域名"`
	Status    string      `json:"status"      dc:"状态"`
	DomainId  string      `json:"domain_id"   dc:"域名"`
	CertId    string      `json:"cert_id"     dc:"证书"`
	CertSubId string      `json:"cert_sub_id" dc:"证书"`
	Type      string      `json:"type"        dc:"类型"`
	Issuer    string      `json:"issuer"      dc:"签发组织"`
	Authority string      `json:"authority"   dc:"授权组织"`
	Signature string      `json:"signature"   dc:"签名方式"`
	ExpireAt  *gtime.Time `json:"expire_at"   dc:"过期时间"`
}

// SslInsertFields 新增证书字段过滤
type SslInsertFields struct {
	AccountId int64       `json:"account_id"  dc:"账号"`
	Domain    string      `json:"domain"      dc:"域名"`
	Status    string      `json:"status"      dc:"状态"`
	DomainId  string      `json:"domain_id"   dc:"域名"`
	CertId    string      `json:"cert_id"     dc:"证书"`
	CertSubId string      `json:"cert_sub_id" dc:"证书"`
	Type      string      `json:"type"        dc:"类型"`
	Issuer    string      `json:"issuer"      dc:"签发组织"`
	Authority string      `json:"authority"   dc:"授权组织"`
	Signature string      `json:"signature"   dc:"签名方式"`
	ExpireAt  *gtime.Time `json:"expire_at"   dc:"过期时间"`
}

// SslEditInp 修改/新增证书
type SslEditInp struct {
	entity.Ssl
}

func (in *SslEditInp) Filter(ctx context.Context) (err error) {
	// 验证账号
	if err := g.Validator().Rules("required").Data(in.AccountId).Messages("账号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证域名
	if err := g.Validator().Rules("required").Data(in.Domain).Messages("域名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:active,backup_issued").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证域名
	if err := g.Validator().Rules("required").Data(in.DomainId).Messages("域名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证证书
	if err := g.Validator().Rules("required").Data(in.CertId).Messages("证书不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证证书
	if err := g.Validator().Rules("required").Data(in.CertSubId).Messages("证书不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证类型
	if err := g.Validator().Rules("required").Data(in.Type).Messages("类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:universal").Data(in.Type).Messages("类型值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证签发组织
	if err := g.Validator().Rules("required").Data(in.Issuer).Messages("签发组织不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:LetsEncrypt").Data(in.Issuer).Messages("签发组织值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证授权组织
	if err := g.Validator().Rules("required").Data(in.Authority).Messages("授权组织不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:lets_encrypt").Data(in.Authority).Messages("授权组织值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证签名方式
	if err := g.Validator().Rules("required").Data(in.Signature).Messages("签名方式不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:ECDSAWithSHA384,SHA256WithRSA").Data(in.Signature).Messages("签名方式值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type SslEditModel struct{}

// SslDeleteInp 删除证书
type SslDeleteInp struct {
	Id interface{} `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *SslDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SslDeleteModel struct{}

// SslViewInp 获取指定证书信息
type SslViewInp struct {
	Id int64 `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *SslViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SslViewModel struct {
	entity.Ssl
}

// SslListInp 获取证书列表
type SslListInp struct {
	form.PageReq
	Id        int64         `json:"id"         dc:"编号"`
	Domain    string        `json:"domain"     dc:"域名"`
	Status    string        `json:"status"     dc:"状态"`
	Type      string        `json:"type"       dc:"类型"`
	Issuer    string        `json:"issuer"     dc:"签发组织"`
	Authority string        `json:"authority"  dc:"授权组织"`
	Signature string        `json:"signature"  dc:"签名方式"`
	CreatedAt []*gtime.Time `json:"created_at" dc:"创建时间"`
}

func (in *SslListInp) Filter(ctx context.Context) (err error) {
	return
}

type SslListModel struct {
	Id        int64       `json:"id"          dc:"编号"`
	AccountId int64       `json:"account_id"  dc:"账号"`
	Domain    string      `json:"domain"      dc:"域名"`
	Status    string      `json:"status"      dc:"状态"`
	DomainId  string      `json:"domain_id"   dc:"域名"`
	CertSubId string      `json:"cert_sub_id" dc:"证书"`
	Type      string      `json:"type"        dc:"类型"`
	Issuer    string      `json:"issuer"      dc:"签发组织"`
	Authority string      `json:"authority"   dc:"授权组织"`
	Signature string      `json:"signature"   dc:"签名方式"`
	CreatedAt *gtime.Time `json:"created_at"  dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at"  dc:"更新时间"`
	ExpireAt  *gtime.Time `json:"expire_at"   dc:"过期时间"`
}

// SslExportModel 导出证书
type SslExportModel struct {
	Id        int64       `json:"id"          dc:"编号"`
	AccountId int64       `json:"account_id"  dc:"账号"`
	Domain    string      `json:"domain"      dc:"域名"`
	Status    string      `json:"status"      dc:"状态"`
	DomainId  string      `json:"domain_id"   dc:"域名"`
	CertSubId string      `json:"cert_sub_id" dc:"证书"`
	Type      string      `json:"type"        dc:"类型"`
	Issuer    string      `json:"issuer"      dc:"签发组织"`
	Authority string      `json:"authority"   dc:"授权组织"`
	Signature string      `json:"signature"   dc:"签名方式"`
	CreatedAt *gtime.Time `json:"created_at"  dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at"  dc:"更新时间"`
	ExpireAt  *gtime.Time `json:"expire_at"   dc:"过期时间"`
}

// SslStatusInp 更新证书状态
type SslStatusInp struct {
	Id     int64 `json:"id" v:"required#编号不能为空" dc:"编号"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SslStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("编号不能为空")
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

type SslStatusModel struct{}
