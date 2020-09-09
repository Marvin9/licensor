package utils

import (
	"path/filepath"
	"strings"
)

// Exists will check if string exist in array
func Exists(key string, strs []string) bool {
	for _, str := range strs {
		if str == key {
			return true
		}
	}
	return false
}

// IsKeywordCommand will check if given string is command or not.
func IsKeywordCommand(str string) bool {
	return Exists(str, Commands)
}

// IsValidExtension will check if given extension is supported or not.
func IsValidExtension(ext string) bool {
	return Exists(ext, SupportedFileExtensions)
}

// GetExtension will give extension without dot
func GetExtension(file string) string {
	ext := filepath.Ext(file)
	return strings.TrimPrefix(ext, ".")
}
