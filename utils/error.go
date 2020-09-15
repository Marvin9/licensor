package utils

import (
	"fmt"
	"os"
)

// LogError will log error and exit with code 1
func LogError(msg interface{}) {
	fmt.Printf("Error: %v\n", msg)
	ShowCursor()
	os.Exit(1)
}

// InvalidFlagError will log specific flag error
func InvalidFlagError(flag string) {
	LogError(fmt.Sprintf(`
Invalid use of %v flag.
licensor -help for flags documentation.`, flag))
}
