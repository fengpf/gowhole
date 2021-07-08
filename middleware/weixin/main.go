package main

import (
	"gowhole/module/weixin/config"
	"gowhole/module/weixin/server"
)

func main() {
	config := config.New()
	server.Run(config)
}
