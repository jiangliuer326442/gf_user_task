package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	// /home/john/temp 是一个目录，当然也可以指定文件
	path := "/tmp/gfsnotify"
	ctx := gctx.New()
	_, err := gfsnotify.Add(path, func(event *gfsnotify.Event) {
		if event.IsCreate() {
			glog.Debug(ctx, "创建文件 : ", event.Path)
		}
		if event.IsWrite() {
			glog.Debug(ctx, "写入文件 : ", event.Path)
		}
		if event.IsRemove() {
			glog.Debug(ctx, "删除文件 : ", event.Path)
		}
		if event.IsRename() {
			glog.Debug(ctx, "重命名文件 : ", event.Path)
		}
		if event.IsChmod() {
			glog.Debug(ctx, "修改权限 : ", event.Path)
		}
		glog.Debug(ctx, event)
	}, true)

	// 移除对该path的监听
	// gfsnotify.Remove(path)

	if err != nil {
		glog.Fatal(ctx, err)
	} else {
		select {}
	}
}
