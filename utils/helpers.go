package utils

import (
	"path/filepath"
	"strings"
)

func Exists(key string, strs []string) bool {
	for _, str := range strs {
		if str == key {
			return true
		}
	}
	return false
}

func IsKeywordCommand(str string) bool {
	return Exists(str, Commands)
}

func IsValidExtension(ext string) bool {
	return Exists(ext, SupportedFileExtensions)
}

func GetExtension(file string) string {
	ext := filepath.Ext(file)
	return strings.TrimPrefix(ext, ".")
}
