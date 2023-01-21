package main

import (
	"github.com/gin-gonic/gin"
	"picgo/config"
	"picgo/dao"
)

func main() {
	r := gin.Default()
	initRouter(r)
	dao.InitDao()
	println("server start")
	r.Run(config.Port)
}
