// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.8
package sysin

import (
	"context"
	"hotgo/addons/fang/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FangUpdateFields 修改租房字段过滤
type FangUpdateFields struct {
	Sid         string `json:"sid"         dc:"编号"`
	Village     string `json:"village"     dc:"小区"`
	Title       string `json:"title"       dc:"标题"`
	Price       int64  `json:"price"       dc:"价格"`
	HouseType   string `json:"houseType"   dc:"户型"`
	Area        string `json:"area"        dc:"面积"`
	ToWard      string `json:"toWard"      dc:"朝向"`
	Lift        string `json:"lift"        dc:"电梯"`
	Water       string `json:"water"       dc:"用水"`
	Electricity string `json:"electricity" dc:"用电"`
	Tenancy     string `json:"tenancy"     dc:"租期"`
	Verify      string `json:"verify"      dc:"核验"`
	CheckIn     string `json:"checkIn"     dc:"入住"`
	Flag        int    `json:"flag"        dc:"旗帜"`
	Url         string `json:"url"         dc:"网址"`
}

// FangInsertFields 新增租房字段过滤
type FangInsertFields struct {
	Sid         string `json:"sid"         dc:"编号"`
	Village     string `json:"village"     dc:"小区"`
	Title       string `json:"title"       dc:"标题"`
	Price       int64  `json:"price"       dc:"价格"`
	HouseType   string `json:"houseType"   dc:"户型"`
	Area        string `json:"area"        dc:"面积"`
	ToWard      string `json:"toWard"      dc:"朝向"`
	Lift        string `json:"lift"        dc:"电梯"`
	Water       string `json:"water"       dc:"用水"`
	Electricity string `json:"electricity" dc:"用电"`
	Tenancy     string `json:"tenancy"     dc:"租期"`
	Verify      string `json:"verify"      dc:"核验"`
	CheckIn     string `json:"checkIn"     dc:"入住"`
	Flag        int    `json:"flag"        dc:"旗帜"`
	Url         string `json:"url"         dc:"网址"`
}

// FangEditInp 修改/新增租房
type FangEditInp struct {
	entity.Fang
}

func (in *FangEditInp) Filter(ctx context.Context) (err error) {
	// 验证编号
	if err := g.Validator().Rules("required").Data(in.Sid).Messages("编号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证小区
	if err := g.Validator().Rules("required").Data(in.Village).Messages("小区不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证价格
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("价格最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证户型
	if err := g.Validator().Rules("required").Data(in.HouseType).Messages("户型不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证朝向
	if err := g.Validator().Rules("required").Data(in.ToWard).Messages("朝向不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证旗帜
	if err := g.Validator().Rules("required").Data(in.Flag).Messages("旗帜不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1").Data(in.Flag).Messages("旗帜值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FangEditModel struct{}

// FangDeleteInp 删除租房
type FangDeleteInp struct {
	Id interface{} `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *FangDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FangDeleteModel struct{}

// FangViewInp 获取指定租房信息
type FangViewInp struct {
	Id int64 `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *FangViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FangViewModel struct {
	entity.Fang
}

// FangListInp 获取租房列表
type FangListInp struct {
	form.PageReq
	Sid       string `json:"sid"       dc:"编号"`
	Village   string `json:"village"   dc:"小区"`
	Title     string `json:"title"     dc:"标题"`
	HouseType string `json:"houseType" dc:"户型"`
	ToWard    string `json:"toWard"    dc:"朝向"`
	Flag      int    `json:"flag"      dc:"旗帜"`
}

func (in *FangListInp) Filter(ctx context.Context) (err error) {
	return
}

type FangListModel struct {
	Id          int64       `json:"id"          dc:"编号"`
	Village     string      `json:"village"     dc:"小区"`
	Price       int64       `json:"price"       dc:"价格"`
	HouseType   string      `json:"houseType"   dc:"户型"`
	Area        string      `json:"area"        dc:"面积"`
	ToWard      string      `json:"toWard"      dc:"朝向"`
	Lift        string      `json:"lift"        dc:"电梯"`
	Water       string      `json:"water"       dc:"用水"`
	Electricity string      `json:"electricity" dc:"用电"`
	CheckIn     string      `json:"checkIn"     dc:"入住"`
	Flag        int         `json:"flag"        dc:"旗帜"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
}

// FangExportModel 导出租房
type FangExportModel struct {
	Id          int64       `json:"id"          dc:"编号"`
	Sid         string      `json:"sid"         dc:"编号"`
	Village     string      `json:"village"     dc:"小区"`
	Title       string      `json:"title"       dc:"标题"`
	Price       int64       `json:"price"       dc:"价格"`
	HouseType   string      `json:"houseType"   dc:"户型"`
	Area        string      `json:"area"        dc:"面积"`
	ToWard      string      `json:"toWard"      dc:"朝向"`
	Lift        string      `json:"lift"        dc:"电梯"`
	Water       string      `json:"water"       dc:"用水"`
	Electricity string      `json:"electricity" dc:"用电"`
	Tenancy     string      `json:"tenancy"     dc:"租期"`
	Verify      string      `json:"verify"      dc:"核验"`
	CheckIn     string      `json:"checkIn"     dc:"入住"`
	Flag        int         `json:"flag"        dc:"旗帜"`
	Url         string      `json:"url"         dc:"网址"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}
