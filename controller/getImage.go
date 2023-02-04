package controller

import (
	"github.com/gin-gonic/gin"
	"path"
	"picgo/dao"
)

func GetImage(c *gin.Context) {
	// 1. 获取文件名
	filename := c.Param("filename")
	// 2. 从数据库中获取文件数据
	println(filename)
	data := dao.FindImage(filename)
	filetype := path.Ext(filename)[1:]

	// 3. 返回文件数据
	if data == "" {
		c.String(200, "not found")
		return
	}
	c.Data(200, "image/"+filetype, []byte(data))
}
