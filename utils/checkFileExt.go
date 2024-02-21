package utils

import "fmt"

func CheckFileExt(extension string) bool {
	fmt.Println(extension)
	switch extension {
	case ".png", ".jpg", ".jpeg", ".webp", ".gif":
		return true
	}
	return false
}
