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
	"hotgo/internal/consts"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserUpdateFields 修改用户信息字段过滤
type UserUpdateFields struct {
	SiteId         int64   `json:"siteId"         dc:"站点编号"`
	Email          string  `json:"email"          dc:"电子邮箱"`
	Password       string  `json:"password"       dc:"密码"`
	Grade          int     `json:"grade"          dc:"等级"`
	Status         int     `json:"status"         dc:"状态"`
	ConsumeMoney   float64 `json:"consumeMoney"   dc:"消费金额"`
	NonMoney       float64 `json:"nonMoney"       dc:"非货币"`
	FreezeMoney    float64 `json:"freezeMoney"    dc:"冻结金额"`
	Amount         float64 `json:"amount"         dc:"账户余额"`
	Payment        string  `json:"payment"        dc:"支付方式"`
	DefaultPayment string  `json:"defaultPayment" dc:"默认支付方式"`
}

// UserInsertFields 新增用户信息字段过滤
type UserInsertFields struct {
	SiteId         int64   `json:"siteId"         dc:"站点编号"`
	Email          string  `json:"email"          dc:"电子邮箱"`
	Password       string  `json:"password"       dc:"密码"`
	Grade          int     `json:"grade"          dc:"等级"`
	Status         int     `json:"status"         dc:"状态"`
	ConsumeMoney   float64 `json:"consumeMoney"   dc:"消费金额"`
	NonMoney       float64 `json:"nonMoney"       dc:"非货币"`
	FreezeMoney    float64 `json:"freezeMoney"    dc:"冻结金额"`
	Amount         float64 `json:"amount"         dc:"账户余额"`
	Payment        string  `json:"payment"        dc:"支付方式"`
	DefaultPayment string  `json:"defaultPayment" dc:"默认支付方式"`
}

// UserEditInp 修改/新增用户信息
type UserEditInp struct {
	entity.User
}

func (in *UserEditInp) Filter(ctx context.Context) (err error) {
	// 验证站点编号
	if err := g.Validator().Rules("required").Data(in.SiteId).Messages("站点编号不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证电子邮箱
	if err := g.Validator().Rules("email").Data(in.Email).Messages("电子邮箱不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证密码
	if err := g.Validator().Rules("regex:^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$").Data(in.Password).Messages("密码必须包含6-18为字母和数字").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证等级
	if err := g.Validator().Rules("required").Data(in.Grade).Messages("等级不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:0,1").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证消费金额
	if err := g.Validator().Rules("required").Data(in.ConsumeMoney).Messages("消费金额不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证非货币
	if err := g.Validator().Rules("required").Data(in.NonMoney).Messages("非货币不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证冻结金额
	if err := g.Validator().Rules("required").Data(in.FreezeMoney).Messages("冻结金额不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证账户余额
	if err := g.Validator().Rules("required").Data(in.Amount).Messages("账户余额不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type UserEditModel struct{}

// UserDeleteInp 删除用户信息
type UserDeleteInp struct {
	Uid interface{} `json:"uid" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *UserDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserDeleteModel struct{}

// UserViewInp 获取指定用户信息信息
type UserViewInp struct {
	Uid int64 `json:"uid" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *UserViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserViewModel struct {
	entity.User
	Info   *entity.UserInfo     `json:"info"`
	Social []*entity.UserSocial `json:"social"`
}

// UserListInp 获取用户信息列表
type UserListInp struct {
	form.PageReq
	Uid       int64         `json:"uid"       dc:"自动编号"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *UserListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserListModel struct {
	Uid            int64       `json:"uid"            dc:"自动编号"`
	SiteId         int64       `json:"siteId"         dc:"站点编号"`
	ShareUid       int64       `json:"shareUid"       dc:"分享用户"`
	Email          string      `json:"email"          dc:"电子邮箱"`
	Password       string      `json:"password"       dc:"密码"`
	Grade          int         `json:"grade"          dc:"等级"`
	Status         int         `json:"status"         dc:"状态"`
	ConsumeMoney   float64     `json:"consumeMoney"   dc:"消费金额"`
	NonMoney       float64     `json:"nonMoney"       dc:"非货币"`
	FreezeMoney    float64     `json:"freezeMoney"    dc:"冻结金额"`
	Amount         float64     `json:"amount"         dc:"账户余额"`
	Payment        string      `json:"payment"        dc:"支付方式"`
	DefaultPayment string      `json:"defaultPayment" dc:"默认支付方式"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// UserExportModel 导出用户信息
type UserExportModel struct {
	Uid            int64       `json:"uid"            dc:"自动编号"`
	SiteId         int64       `json:"siteId"         dc:"站点编号"`
	ShareUid       int64       `json:"shareUid"       dc:"分享用户"`
	Email          string      `json:"email"          dc:"电子邮箱"`
	Password       string      `json:"password"       dc:"密码"`
	Grade          int         `json:"grade"          dc:"等级"`
	Status         int         `json:"status"         dc:"状态"`
	ConsumeMoney   float64     `json:"consumeMoney"   dc:"消费金额"`
	NonMoney       float64     `json:"nonMoney"       dc:"非货币"`
	FreezeMoney    float64     `json:"freezeMoney"    dc:"冻结金额"`
	Amount         float64     `json:"amount"         dc:"账户余额"`
	Payment        string      `json:"payment"        dc:"支付方式"`
	DefaultPayment string      `json:"defaultPayment" dc:"默认支付方式"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
}

// UserStatusInp 更新用户信息状态
type UserStatusInp struct {
	Uid    int64 `json:"uid" v:"required#自动编号不能为空" dc:"自动编号"`
	Status int   `json:"status" dc:"状态"`
}

func (in *UserStatusInp) Filter(ctx context.Context) (err error) {
	if in.Uid <= 0 {
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

type UserStatusModel struct{}
