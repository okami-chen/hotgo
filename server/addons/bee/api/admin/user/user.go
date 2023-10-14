// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package user

import (
	"hotgo/addons/bee/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询用户信息列表
type ListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"用户信息" summary:"获取用户信息列表"`
	sysin.UserListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.UserListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出用户信息列表
type ExportReq struct {
	g.Meta `path:"/user/export" method:"get" tags:"用户信息" summary:"导出用户信息列表"`
	sysin.UserListInp
}

type ExportRes struct{}

// ViewReq 获取用户信息指定信息
type ViewReq struct {
	g.Meta `path:"/user/view" method:"get" tags:"用户信息" summary:"获取用户信息指定信息"`
	sysin.UserViewInp
}

type ViewRes struct {
	*sysin.UserViewModel
}

// EditReq 修改/新增用户信息
type EditReq struct {
	g.Meta `path:"/user/edit" method:"post" tags:"用户信息" summary:"修改/新增用户信息"`
	sysin.UserEditInp
}

type EditRes struct{}

// DeleteReq 删除用户信息
type DeleteReq struct {
	g.Meta `path:"/user/delete" method:"post" tags:"用户信息" summary:"删除用户信息"`
	sysin.UserDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新用户信息状态
type StatusReq struct {
	g.Meta `path:"/user/status" method:"post" tags:"用户信息" summary:"更新用户信息状态"`
	sysin.UserStatusInp
}

type StatusRes struct{}
