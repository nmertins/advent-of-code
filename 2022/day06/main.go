package main

import (
	"fmt"
	"strings"

	"github.com/nmertins/advent-of-code/2022/utils"
)

const (
	StartOfPacketMarkerLength = 4
)

func FindMarker(input string) string {
	for i := 0; i < len(input)-StartOfPacketMarkerLength; i++ {
		marker := input[i : i+StartOfPacketMarkerLength]
		if AllCharactersUnique(marker) {
			return marker
		}
	}

	return ""
}

func GetMarkerIndex(signal string, marker string) int {
	return strings.Index(signal, marker) + StartOfPacketMarkerLength
}

func AllCharactersUnique(input string) bool {
	unique := true
	for i := 0; i < len(input); i++ {
		s, other := string(input[i]), strings.Join([]string{input[:i], input[i+1:]}, "")
		unique = unique && !strings.Contains(other, s)
	}
	return unique
}

func main() {
	signal := utils.ReadFile("resources/input.txt")[0]
	marker := FindMarker(signal)
	index := GetMarkerIndex(signal, marker)
	fmt.Println(index)
}
