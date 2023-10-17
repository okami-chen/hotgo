// Package crons
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package crons

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/internal/library/cron"
)

// 定时任务.
// 插件中的定时任务可以统一在这里注册和处理

func init() {
	cron.Register(CleanLog)
}

var CleanLog = &CLogClean{name: "log_clean"}

type CLogClean struct {
	name string
}

func (c *CLogClean) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *CLogClean) Execute(ctx context.Context) {

	arr := garray.NewArray()
	arr.Append("logs/cron")
	arr.Append("logs/logger")
	arr.Append("logs/queue")
	arr.Append("logs/server")
	arr.Append("logs/sql")
	arr.Append("logs/tcpClient")
	arr.Append("logs/tcpServer")

	// 遍历数组
	arr.Iterator(func(index int, value interface{}) bool {
		//fmt.Printf("索引 %d 的值为 %s\n", index, value)
		//cron.Logger().Infof(ctx, "Clean Log :%v", value)
		c.Clean(ctx, value.(string))
		return true // 返回 true 继续遍历，返回 false 终止遍历
	})
}

func (c *CLogClean) Clean(ctx context.Context, dirPath string) {
	// 获取今天的日期
	today := gtime.Now().Format("Y-m-d")

	// 获取目录下的所有文件列表
	files, _ := gfile.ScanDirFile(dirPath, "*.log")

	for _, file := range files {
		// 取出文件名中的日期部分
		fileDate := gfile.Basename(file)[:10]

		// 比较日期，如果文件日期早于今天，则删除文件
		if fileDate < today {
			err := gfile.Remove(gfile.Join(dirPath, file))
			if err != nil {
				cron.Logger().Infof(ctx, "删除失败 :%v", gfile.Join(dirPath, file))
			} else {
				cron.Logger().Infof(ctx, "删除成功 :%v", gfile.Join(dirPath, file))
			}
		} else {
			cron.Logger().Infof(ctx, "没有日志 :%v", dirPath)
		}
	}
}
