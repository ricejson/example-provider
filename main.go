package main

import (
	"github.com/ricejson/example-consumer/service"
	"github.com/ricejson/rice-rpc-easy/registry"
	"github.com/ricejson/rice-rpc-easy/server"
)

func main() {
	// 注册用户服务到本地注册器
	localRegistry := registry.NewLocalRegistry()
	var userService = service.NewUserServiceImpl()
	localRegistry.Register("UserService", userService)
	// 启动HTTP服务器
	webServer := server.NewWebServer()
	// 监听8080端口
	webServer.DoStart(8080)
}
