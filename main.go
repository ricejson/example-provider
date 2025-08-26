package main

import (
	"github.com/ricejson/rice-rpc-easy/server"
)

func main() {
	// 启动HTTP服务器
	webServer := server.NewWebServer()
	// 监听8080端口
	webServer.DoStart(8080)
}
