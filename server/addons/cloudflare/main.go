// Package cloudflare
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cloudflare

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	_ "hotgo/addons/cloudflare/crons"
	"hotgo/addons/cloudflare/global"
	_ "hotgo/addons/cloudflare/logic"
	_ "hotgo/addons/cloudflare/queues"
	"hotgo/addons/cloudflare/router"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
	"os"
	"strings"
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
			Label:       `cloudflare`,
			Name:        `cloudflare`,
			Group:       1,
			Logo:        "",
			Brief:       `cloudflare`,
			Description: `cloudflare`,
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

	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		fileName := "./addons/cloudflare/sql/install.sql"
		e := m.ImportSql(ctx, fileName)
		if e != nil {
			g.Log().Warningf(ctx, "file: %s Not Found", fileName)
			return nil
		}
		return nil
	})
	return
}

// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		fileName := "./addons/cloudflare/sql/upgrade_" + m.skeleton.Version + ".sql"
		e := m.ImportSql(ctx, fileName)
		if e != nil {
			g.Log().Warningf(ctx, "file: %s Not Found", fileName)
			return nil
		}
		return nil
	})
	return
}

// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	g.DB().Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		fileName := "./addons/cloudflare/sql/uninstall.sql"
		e := m.ImportSql(ctx, fileName)
		if e != nil {
			g.Log().Warningf(ctx, "file: %s Not Found", fileName)
			return nil
		}
		return nil
	})
	return
}

func (m *module) ImportSql(ctx context.Context, path string) error {
	rows, e := os.ReadFile(path)
	if e != nil {
		return e
	}
	sqlArr := strings.Split(string(rows), "\n")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" || strings.HasPrefix(sql, "--") {
			continue
		}
		exec, err := g.DB().Exec(ctx, sql)
		g.Log().Infof(ctx, "views.ImportSql sql:%v, exec:%+v, err:%+v", sql, exec, err)
		if err != nil {
			return err
		}
	}
	return nil
}
