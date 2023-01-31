package main

import (
	"picgo/config"
	"picgo/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 关闭debug模式

	initRouter(r)
	dao.InitDao()
	println("server start")
	r.Run(config.Port)
}
