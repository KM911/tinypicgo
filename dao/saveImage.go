package dao

import (
	"encoding/base64"
)

func SaveImage(filename string, data string) {
	var image ImageTable
	db.Find(&image, "filename = ?", filename)
	if image.ID == 0 {

		nedata, _ := base64.StdEncoding.DecodeString(data)
		db.Create(&ImageTable{Filename: filename, Data: string(nedata)})
	} else {
		db.Model(&image).Update("data", data)
	}
}

