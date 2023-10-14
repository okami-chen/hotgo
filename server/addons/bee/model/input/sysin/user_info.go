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

// UserInfoUpdateFields 修改客户信息扩展字段过滤
type UserInfoUpdateFields struct {
	Uid         int64  `json:"uid"         dc:"用户"`
	Username    string `json:"username"    dc:"用户名"`
	FirstName   string `json:"firstName"   dc:"用户名"`
	LastName    string `json:"lastName"    dc:"用户姓"`
	Mobile      string `json:"mobile"      dc:"手机"`
	Country     string `json:"country"     dc:"国家"`
	Avatar      string `json:"avatar"      dc:"头像"`
	Account     string `json:"account"     dc:"账号"`
	Gender      string `json:"gender"      dc:"性别"`
	SocialMedia string `json:"socialMedia" dc:"联系"`
}

// UserInfoInsertFields 新增客户信息扩展字段过滤
type UserInfoInsertFields struct {
	Uid         int64  `json:"uid"         dc:"用户"`
	Username    string `json:"username"    dc:"用户名"`
	FirstName   string `json:"firstName"   dc:"用户名"`
	LastName    string `json:"lastName"    dc:"用户姓"`
	Mobile      string `json:"mobile"      dc:"手机"`
	Country     string `json:"country"     dc:"国家"`
	Avatar      string `json:"avatar"      dc:"头像"`
	Account     string `json:"account"     dc:"账号"`
	Gender      string `json:"gender"      dc:"性别"`
	SocialMedia string `json:"socialMedia" dc:"联系"`
}

// UserInfoEditInp 修改/新增客户信息扩展
type UserInfoEditInp struct {
	entity.UserInfo
}

func (in *UserInfoEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户
	if err := g.Validator().Rules("required").Data(in.Uid).Messages("用户不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证手机
	if err := g.Validator().Rules("telephone").Data(in.Mobile).Messages("手机不是座机号码").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type UserInfoEditModel struct{}

// UserInfoDeleteInp 删除客户信息扩展
type UserInfoDeleteInp struct {
	Id interface{} `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *UserInfoDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserInfoDeleteModel struct{}

// UserInfoViewInp 获取指定客户信息扩展信息
type UserInfoViewInp struct {
	Id int64 `json:"id" v:"required#自动编号不能为空" dc:"自动编号"`
}

func (in *UserInfoViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserInfoViewModel struct {
	entity.UserInfo
}

// UserInfoListInp 获取客户信息扩展列表
type UserInfoListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"自动编号"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *UserInfoListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserInfoListModel struct {
	Id          int64       `json:"id"          dc:"自动编号"`
	Uid         int64       `json:"uid"         dc:"用户"`
	Username    string      `json:"username"    dc:"用户名"`
	FirstName   string      `json:"firstName"   dc:"用户名"`
	LastName    string      `json:"lastName"    dc:"用户姓"`
	Mobile      string      `json:"mobile"      dc:"手机"`
	Country     string      `json:"country"     dc:"国家"`
	Avatar      string      `json:"avatar"      dc:"头像"`
	Account     string      `json:"account"     dc:"账号"`
	Gender      string      `json:"gender"      dc:"性别"`
	SocialMedia string      `json:"socialMedia" dc:"联系"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}

// UserInfoExportModel 导出客户信息扩展
type UserInfoExportModel struct {
	Id          int64       `json:"id"          dc:"自动编号"`
	Uid         int64       `json:"uid"         dc:"用户"`
	Username    string      `json:"username"    dc:"用户名"`
	FirstName   string      `json:"firstName"   dc:"用户名"`
	LastName    string      `json:"lastName"    dc:"用户姓"`
	Mobile      string      `json:"mobile"      dc:"手机"`
	Country     string      `json:"country"     dc:"国家"`
	Avatar      string      `json:"avatar"      dc:"头像"`
	Account     string      `json:"account"     dc:"账号"`
	Gender      string      `json:"gender"      dc:"性别"`
	SocialMedia string      `json:"socialMedia" dc:"联系"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}
