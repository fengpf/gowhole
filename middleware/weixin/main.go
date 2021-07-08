package main

import (
	"gowhole/middleware/weixin/config"
	"gowhole/middleware/weixin/server"
)

func main() {
	config := config.New()
	server.Run(config)
}
