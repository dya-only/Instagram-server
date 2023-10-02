package utils

import "path/filepath"

func GetContentType(filename string) string {
	switch filepath.Ext(filename) {
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	default:
		return "image/jpeg"
	}
}
