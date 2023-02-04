package dao

func FindImage(filename string) string {
	var image ImageTable
	db.Find(&image, "filename = ?", filename)
	if image.ID != 0 {
		return image.Data
	} else {
		return ""
	}
}
