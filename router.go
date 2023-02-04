package main

import (
	"picgo/controller"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	// 这里我们全部将图片文件放到sqlite数据库中，所以不需要进行静态文件的处理

	apiRouter := r.Group("/tinypicgo")
	apiRouter.GET("/test", controller.Test)

	apiRouter.POST("/upload", controller.UploadImage)
	// 这里假如是 这样的形势  /api/1.png 那么我们就可以使用这个方法
	apiRouter.GET("/image/:filename", controller.GetImage)
	apiRouter.GET("/del/:filename", controller.DeleteImage)
}
