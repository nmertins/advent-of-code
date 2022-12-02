package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func ParseInput(input []string) (totals []int) {
	total := 0
	for _, line := range input {
		intVar, err := strconv.Atoi(line)
		if err != nil {
			totals = append(totals, total)
			total = 0
		} else {
			total += intVar
		}

	}

	return append(totals, total)
}

func Max(values []int) int {
	sort.Ints(values)
	return values[len(values)-1]
}

func main() {
	input := ReadFile("resources/input.txt")
	totals := ParseInput(input)
	sort.Ints(totals)

	fmt.Println(totals[len(totals)-1])

	combinedTotal := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]

	fmt.Println(combinedTotal)
}
