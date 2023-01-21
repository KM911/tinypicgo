package main

import (
	"github.com/gin-gonic/gin"
	"picgo/config"
	"picgo/dao"
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
