// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/addons/order/global"
	"hotgo/addons/order/model/input/sysin"
	"hotgo/addons/order/service"
	"hotgo/internal/library/contexts"
)

type sSysIndex struct{}

func NewSysIndex() *sSysIndex {
	return &sSysIndex{}
}

func init() {
	service.RegisterSysIndex(NewSysIndex())
}

// Test 测试
func (s *sSysIndex) Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error) {
	res = new(sysin.IndexTestModel)
	res.Name = in.Name
	res.Module = fmt.Sprintf("当前插件模块是：%s，当前应用模块是：%s", global.GetSkeleton().Name, contexts.Get(ctx).Module)
	res.Time = gtime.Now()
	return
}
