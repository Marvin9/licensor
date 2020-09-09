package utils

import (
	"fmt"
	"os"
)

func LogError(msg interface{}) {
	fmt.Printf("\nError: %v\n", msg)
	os.Exit(1)
}

func InvalidFlagError(flag string) {
	LogError(fmt.Sprintf("Invalid use of %v flag.", flag))
}
