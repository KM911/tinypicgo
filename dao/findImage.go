package dao

import ()

func FindImage(filename string) string {
	var image ImageTable
	db.Find(&image, "filename = ?", filename)
	return image.Data
}
