package utils

import (
	"fmt"
	"os"
)

func RemoveFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
}
