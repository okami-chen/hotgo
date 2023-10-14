// Package device
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package device

import (
	"hotgo/addons/game/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询设备列表
type ListReq struct {
	g.Meta `path:"/device/list" method:"get" tags:"设备" summary:"获取设备列表"`
	sysin.DeviceListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.DeviceListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出设备列表
type ExportReq struct {
	g.Meta `path:"/device/export" method:"get" tags:"设备" summary:"导出设备列表"`
	sysin.DeviceListInp
}

type ExportRes struct{}

// ViewReq 获取设备指定信息
type ViewReq struct {
	g.Meta `path:"/device/view" method:"get" tags:"设备" summary:"获取设备指定信息"`
	sysin.DeviceViewInp
}

type ViewRes struct {
	*sysin.DeviceViewModel
}

// EditReq 修改/新增设备
type EditReq struct {
	g.Meta `path:"/device/edit" method:"post" tags:"设备" summary:"修改/新增设备"`
	sysin.DeviceEditInp
}

type EditRes struct{}

// DeleteReq 删除设备
type DeleteReq struct {
	g.Meta `path:"/device/delete" method:"post" tags:"设备" summary:"删除设备"`
	sysin.DeviceDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取设备最大排序
type MaxSortReq struct {
	g.Meta `path:"/device/maxSort" method:"get" tags:"设备" summary:"获取设备最大排序"`
	sysin.DeviceMaxSortInp
}

type MaxSortRes struct {
	*sysin.DeviceMaxSortModel
}

// StatusReq 更新设备状态
type StatusReq struct {
	g.Meta `path:"/device/status" method:"post" tags:"设备" summary:"更新设备状态"`
	sysin.DeviceStatusInp
}

type StatusRes struct{}
