// Package index
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package index

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/card/model/input/sysin"
)

// TestReq 测试
type TestReq struct {
	g.Meta `path:"/index/test" method:"get" tags:"card" summary:"测试后台API"`
	sysin.IndexTestInp
}

type TestRes struct {
	*sysin.IndexTestModel
}
