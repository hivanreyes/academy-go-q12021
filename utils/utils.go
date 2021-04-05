package utils

import (
	"fmt"
	"os"
)

// RemovingFile file
func RemoveFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
}
