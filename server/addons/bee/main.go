// Package bee
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package bee

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	_ "hotgo/addons/bee/crons"
	"hotgo/addons/bee/global"
	_ "hotgo/addons/bee/logic"
	_ "hotgo/addons/bee/queues"
	"hotgo/addons/bee/router"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
	"sync"
)

type module struct {
	skeleton *addons.Skeleton
	ctx      context.Context
	sync.Mutex
}

func init() {
	newModule()
}

func newModule() {
	m := &module{
		skeleton: &addons.Skeleton{
			Label:       `bee`,
			Name:        `bee`,
			Group:       1,
			Logo:        "",
			Brief:       `bee`,
			Description: `bee`,
			Author:      `陈德华`,
			Version:     `v1.0.0`, // 当该版本号高于已安装的版本号时，会提示可以更新
		},
		ctx: gctx.New(),
	}
	addons.RegisterModule(m)
}

// Init 初始化
func (m *module) Init(ctx context.Context) {
	global.Init(ctx, m.skeleton)
	// ...
}

// InitRouter 初始化WEB路由
func (m *module) InitRouter(ctx context.Context, group *ghttp.RouterGroup) {
	m.Init(ctx)
	group.Middleware(service.Middleware().Addon)
	router.Admin(ctx, group)
	router.Api(ctx, group)
	router.Home(ctx, group)
	router.WebSocket(ctx, group)
}

// Ctx 上下文
func (m *module) Ctx() context.Context {
	return m.ctx
}

// GetSkeleton 架子
func (m *module) GetSkeleton() *addons.Skeleton {
	return m.skeleton
}

// Install 安装模块
func (m *module) Install(ctx context.Context) (err error) {
	// ...
	return
}

// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	// ...
	return
}

// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	// ...
	return
}
