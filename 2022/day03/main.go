package main

import (
	"fmt"
	"strings"

	"github.com/nmertins/advent-of-code/2022/utils"
)

const (
	priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func ParseInput(input []string) [][]string {
	rucksacks := make([][]string, len(input))

	for i, line := range input {
		rucksacks[i] = make([]string, 2)
		split := len(line) / 2
		rucksacks[i][0] = line[:split]
		rucksacks[i][1] = line[split:]
	}

	return rucksacks
}

func FindMatchingItems(firstCompartment string, secondCompartment string) string {
	for _, item := range firstCompartment {
		if strings.Contains(secondCompartment, string(item)) {
			return string(item)
		}
	}

	return ""
}

func GetPriority(item string) int {
	return strings.Index(priority, item) + 1
}

func main() {
	input := utils.ReadFile("resources/input.txt")
	rucksacks := ParseInput(input)
	total := 0
	for _, rucksack := range rucksacks {
		item := FindMatchingItems(rucksack[0], rucksack[1])
		total += GetPriority(item)
	}

	fmt.Println(total)
}
