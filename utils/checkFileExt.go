package utils

func CheckFileExt(extension string) bool {
	switch extension {
	case "png", "jpg", "jpeg", "webp", "gif":
		return true
	}
	return false
}
