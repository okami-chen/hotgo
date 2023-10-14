// Package userinfo
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package userinfo

import (
	"hotgo/addons/bee/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询客户信息扩展列表
type ListReq struct {
	g.Meta `path:"/userInfo/list" method:"get" tags:"客户信息扩展" summary:"获取客户信息扩展列表"`
	sysin.UserInfoListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.UserInfoListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出客户信息扩展列表
type ExportReq struct {
	g.Meta `path:"/userInfo/export" method:"get" tags:"客户信息扩展" summary:"导出客户信息扩展列表"`
	sysin.UserInfoListInp
}

type ExportRes struct{}

// ViewReq 获取客户信息扩展指定信息
type ViewReq struct {
	g.Meta `path:"/userInfo/view" method:"get" tags:"客户信息扩展" summary:"获取客户信息扩展指定信息"`
	sysin.UserInfoViewInp
}

type ViewRes struct {
	*sysin.UserInfoViewModel
}

// EditReq 修改/新增客户信息扩展
type EditReq struct {
	g.Meta `path:"/userInfo/edit" method:"post" tags:"客户信息扩展" summary:"修改/新增客户信息扩展"`
	sysin.UserInfoEditInp
}

type EditRes struct{}

// DeleteReq 删除客户信息扩展
type DeleteReq struct {
	g.Meta `path:"/userInfo/delete" method:"post" tags:"客户信息扩展" summary:"删除客户信息扩展"`
	sysin.UserInfoDeleteInp
}

type DeleteRes struct{}
