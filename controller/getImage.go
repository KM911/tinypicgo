package controller

import (
	"github.com/gin-gonic/gin"
	"os"
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
	println(data[:10])
	//c.Data(200, "image/"+string(filetype), []byte(data)

	// 3. 返回文件数据
	c.Data(200, "image/"+filetype, []byte(data))

	// 写入文件进行测试
	os.WriteFile(filename, []byte(data), 0666)
}
