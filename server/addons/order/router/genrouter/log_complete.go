// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package genrouter

import "hotgo/addons/order/controller/admin/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.LogComplete) // 完成时间
}
