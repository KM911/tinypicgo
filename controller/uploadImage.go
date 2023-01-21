package controller

import (
	"github.com/gin-gonic/gin"
	"picgo/config"
	"picgo/dao"
)

type Image struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func UploadImage(c *gin.Context) {
	println("upload image")
	// 解析json
	var image Image
	c.BindJSON(&image)
	// 保存图片
	dao.SaveImage(image.Name, image.Data)
	// 返回网站地址
	c.String(200, config.Host+"api/image/"+image.Name)

}
