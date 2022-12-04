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

func GetGroupBadges(input []string) []string {
	var badges []string
	for i := 0; i < len(input); i += 3 {
		first := input[i]
		second := input[i+1]
		third := input[i+2]

		for _, rune := range first {
			s := string(rune)
			if strings.Contains(second, s) && strings.Contains(third, s) {
				badges = append(badges, s)
				break
			}
		}
	}
	return badges
}

func main() {
	input := utils.ReadFile("resources/input.txt")
	// rucksacks := ParseInput(input)
	// total := 0
	// for _, rucksack := range rucksacks {
	// 	item := FindMatchingItems(rucksack[0], rucksack[1])
	// 	total += GetPriority(item)
	// }

	// fmt.Println(total)
	badges := GetGroupBadges(input)
	total := 0
	for _, badge := range badges {
		total += GetPriority(badge)
	}
	fmt.Println(total)
}
