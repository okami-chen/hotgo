// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package sysin

import (
	"context"
	"hotgo/addons/bee/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemSiteUpdateFields 修改站群字段过滤
type SystemSiteUpdateFields struct {
	MessageType int64  `json:"messageType" dc:"客服类型"`
	Sort        int64  `json:"sort"        dc:"排序"`
	SiteName    string `json:"siteName"    dc:"网站名称"`
	SiteUrl     string `json:"siteUrl"     dc:"网站网址"`
	SiteLogo    string `json:"siteLogo"    dc:"网站logo"`
	Currencys   string `json:"currencys"   dc:"货币，多个逗号分隔"`
	Languages   string `json:"languages"   dc:"语言，多个逗号分隔"`
	SiteRemark  string `json:"siteRemark"  dc:"网站备注"`
}

// SystemSiteInsertFields 新增站群字段过滤
type SystemSiteInsertFields struct {
	MessageType int64  `json:"messageType" dc:"客服类型"`
	Sort        int64  `json:"sort"        dc:"排序"`
	SiteName    string `json:"siteName"    dc:"网站名称"`
	SiteUrl     string `json:"siteUrl"     dc:"网站网址"`
	SiteLogo    string `json:"siteLogo"    dc:"网站logo"`
	Currencys   string `json:"currencys"   dc:"货币，多个逗号分隔"`
	Languages   string `json:"languages"   dc:"语言，多个逗号分隔"`
	SiteRemark  string `json:"siteRemark"  dc:"网站备注"`
}

// SystemSiteEditInp 修改/新增站群
type SystemSiteEditInp struct {
	entity.SystemSite
}

func (in *SystemSiteEditInp) Filter(ctx context.Context) (err error) {
	// 验证客服类型
	if err := g.Validator().Rules("required").Data(in.MessageType).Messages("客服类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证网站名称
	if err := g.Validator().Rules("required").Data(in.SiteName).Messages("网站名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证网站网址
	if err := g.Validator().Rules("required").Data(in.SiteUrl).Messages("网站网址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证货币，多个逗号分隔
	if err := g.Validator().Rules("required").Data(in.Currencys).Messages("货币，多个逗号分隔不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证语言，多个逗号分隔
	if err := g.Validator().Rules("required").Data(in.Languages).Messages("语言，多个逗号分隔不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type SystemSiteEditModel struct{}

// SystemSiteDeleteInp 删除站群
type SystemSiteDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *SystemSiteDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SystemSiteDeleteModel struct{}

// SystemSiteViewInp 获取指定站群信息
type SystemSiteViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *SystemSiteViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SystemSiteViewModel struct {
	entity.SystemSite
}

// SystemSiteListInp 获取站群列表
type SystemSiteListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自动编号"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *SystemSiteListInp) Filter(ctx context.Context) (err error) {
	return
}

type SystemSiteListModel struct {
	Id          int64       `json:"id"          dc:"自动编号"`
	MessageType int64       `json:"messageType" dc:"客服类型"`
	Sort        int64       `json:"sort"        dc:"排序"`
	SiteName    string      `json:"siteName"    dc:"网站名称"`
	SiteUrl     string      `json:"siteUrl"     dc:"网站网址"`
	SiteLogo    string      `json:"siteLogo"    dc:"网站logo"`
	Currencys   string      `json:"currencys"   dc:"货币，多个逗号分隔"`
	Languages   string      `json:"languages"   dc:"语言，多个逗号分隔"`
	SiteRemark  string      `json:"siteRemark"  dc:"网站备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}

// SystemSiteExportModel 导出站群
type SystemSiteExportModel struct {
	Id          int64       `json:"id"          dc:"自动编号"`
	MessageType int64       `json:"messageType" dc:"客服类型"`
	Sort        int64       `json:"sort"        dc:"排序"`
	SiteName    string      `json:"siteName"    dc:"网站名称"`
	SiteUrl     string      `json:"siteUrl"     dc:"网站网址"`
	SiteLogo    string      `json:"siteLogo"    dc:"网站logo"`
	Currencys   string      `json:"currencys"   dc:"货币，多个逗号分隔"`
	Languages   string      `json:"languages"   dc:"语言，多个逗号分隔"`
	SiteRemark  string      `json:"siteRemark"  dc:"网站备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}

// SystemSiteMaxSortInp 获取站群最大排序
type SystemSiteMaxSortInp struct{}

func (in *SystemSiteMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type SystemSiteMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}
