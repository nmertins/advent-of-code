package main

import(
	"strings"
)

type RockPaperScissorsShape int

const (
	Rock RockPaperScissorsShape = iota + 1
	Paper
	Scissors
)

func (s RockPaperScissorsShape) Compare(other RockPaperScissorsShape) int {
	if s == Rock {
		if other == Rock {
			return 3
		}
		if other == Paper {
			return 0
		}
		if other == Scissors {
			return 6
		}
	}
	if s == Paper {
		if other == Rock {
			return 6
		}
		if other == Paper {
			return 3
		}
		if other == Scissors {
			return 0
		}
	}
	if s == Scissors {
		if other == Rock {
			return 0
		}
		if other == Paper {
			return 6
		}
		if other == Scissors {
			return 3
		}
	}

	return -1
}

func ParseInput(strategy string) (RockPaperScissorsShape, RockPaperScissorsShape) {
	shapes := strings.Split(strategy, " ")
	return convertStringToShape(shapes[0]), convertStringToShape(shapes[1])
}

func convertStringToShape(s string) RockPaperScissorsShape {
	switch s {
	case "A":
		fallthrough
	case "X":
		return Rock
	case "B":
		fallthrough
	case "Y":
		return Paper
	case "C":
		fallthrough
	case "Z":
		return Scissors
	default:
		return -1
	}
}

func Score(opponentShape RockPaperScissorsShape, myShape RockPaperScissorsShape) int {
	return int(myShape) + myShape.Compare(opponentShape)
}