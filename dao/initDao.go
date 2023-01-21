// 这里是我们使用sqlite的部分

package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ImageTable struct {
	gorm.Model
	Filename string
	Data     string
}

var (
	db  *gorm.DB
	err error
)
func InitDao() {
	// 1. 连接数据库
	// 2. 初始化表
	// 3. 插入数据

	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 2. 初始化表
	db.AutoMigrate(&ImageTable{})
	// 这里我希望db是一个全局变量 该怎么做呢？
	// 
}
