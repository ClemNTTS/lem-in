package utils

import (
	"fmt"
	"os"
)

func GetFileContent(fileName string) string {

	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	content := string(b)
	return content
}
