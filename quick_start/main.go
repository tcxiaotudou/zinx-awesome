package main

import (
	"ZinxawesomeProject/quick_start/router"
	"github.com/aceld/zinx/znet"
)

func main() {
	// 1.创建一个server句柄
	s := znet.NewServer()

	// 2.配置路由
	s.AddRouter(0, &router.PingRouter{})

	// 3.开启服务
	s.Serve()
}
