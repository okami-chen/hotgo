// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.8.9
package genrouter

import "hotgo/addons/bee/controller/admin/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.User) // 用户信息
}
