package utils

import (
	"fmt"
	"os"
)

// Removing file
func RemoveFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
}
