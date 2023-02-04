package dao

import (
	"encoding/base64"
)

func SaveImage(filename string, data string) {
	var image ImageTable
	nedata, _ := base64.StdEncoding.DecodeString(data)
	db.Find(&image, "filename = ?", filename)

	if image.ID == 0 {
		db.Create(&ImageTable{Filename: filename, Data: string(nedata)})
	} else {
		db.Create(&ImageTable{Filename: "A" + filename, Data: string(nedata)})
	}
}
