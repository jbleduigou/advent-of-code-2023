package util

import (
	"fmt"
	"os"
)

func GetFileSilently(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}
