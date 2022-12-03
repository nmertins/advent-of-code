package main

import(
	"strings"
	"fmt"

	"github.com/nmertins/advent-of-code/2022/utils"
)

type RockPaperScissorsShape int
const (
	Rock RockPaperScissorsShape = iota + 1
	Paper
	Scissors
)

type GameOutcome int
const (
	multiplier = 3
	Win GameOutcome = iota * multiplier
	Lose
	Draw
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

func (s RockPaperScissorsShape) BackoutDesiredShape(desiredOutcome GameOutcome) RockPaperScissorsShape {
	if s == Rock {
		if desiredOutcome == Win {
			return Paper
		}
		if desiredOutcome == Lose {
			return Scissors
		}
		if desiredOutcome == Draw {
			return Rock
		}
	}
	if s == Paper {
		if desiredOutcome == Win {
			return Scissors
		}
		if desiredOutcome == Lose {
			return Rock
		}
		if desiredOutcome == Draw {
			return Paper
		}
	}
	if s == Scissors {
		if desiredOutcome == Win {
			return Rock
		}
		if desiredOutcome == Lose {
			return Paper
		}
		if desiredOutcome == Draw {
			return Scissors
		}
	}

	return -1
}

func ParseInput(strategy string) (RockPaperScissorsShape, RockPaperScissorsShape) {
	shapes := strings.Split(strategy, " ")
	return convertStringToShape(shapes[0]), convertStringToShape(shapes[1])
}

func ParseInput2(strategy string) (RockPaperScissorsShape, RockPaperScissorsShape) {
	shapes := strings.Split(strategy, " ")
	opponentShape := convertStringToShape(shapes[0])
	gameOutcome := convertStringToShapePart2(shapes[1])
	myShape := opponentShape.BackoutDesiredShape(gameOutcome)
	return opponentShape, myShape
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

func convertStringToShapePart2(winLoseDraw string) GameOutcome {
	switch winLoseDraw {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		return -1
	}
}

func Score(opponentShape RockPaperScissorsShape, myShape RockPaperScissorsShape) int {
	return int(myShape) + myShape.Compare(opponentShape)
}

func main() {
	input := utils.ReadFile("resources/input.txt")
	total := 0
	for _, line := range input {
		opponentShape, myShape := ParseInput2(line)
		total += Score(opponentShape, myShape)
	}
	fmt.Println(total)
}