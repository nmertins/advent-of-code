package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) (fileLines []string) {
	readFile, err := os.Open(path)
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
		return make([]string, 0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
