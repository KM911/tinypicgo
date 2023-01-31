package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"picgo/dao"
)

func main() {
	var iamgeList []dao.ImageTable
	db, err := gorm.Open(sqlite.Open("dao/image.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Find(&iamgeList)
	os.Mkdir("image", 0666)
	for _, v := range iamgeList {
		// 将其写入到文件中
		os.WriteFile("image/"+v.Filename, []byte(v.Data), 0666)
	}

}
