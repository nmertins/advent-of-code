package main

import(
	"strings"
)

type RockPaperScissorsShape int

const (
	Rock RockPaperScissorsShape = iota
	Paper
	Scissors
)

func ParseInput(strategy string) (RockPaperScissorsShape, RockPaperScissorsShape) {
	shapes := strings.Split(strategy, " ")
	return convertStringToShape(shapes[0]), convertStringToShape(shapes[1])
}

func convertStringToShape(s string) RockPaperScissorsShape {
	switch s {
	case "A":
		return Rock
	case "X":
		return Rock
	case "B":
		return Paper
	case "Y":
		return Paper
	case "C":
		return Scissors
	case "Z":
		return Scissors
	default:
		return -1
	}
}