package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/nmertins/advent-of-code/2022/utils"
)

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
	input := utils.ReadFile("resources/input.txt")
	totals := ParseInput(input)
	sort.Ints(totals)

	fmt.Println(totals[len(totals)-1])

	combinedTotal := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]

	fmt.Println(combinedTotal)
}
