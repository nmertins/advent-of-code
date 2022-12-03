package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay02(t *testing.T) {
	t.Run("parse sample input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		opponentShape, myShape := ParseInput(input[0])
		assertShapesEqual(t, Rock, opponentShape)
		assertShapesEqual(t, Paper, myShape)

		opponentShape, myShape = ParseInput(input[1])
		assertShapesEqual(t, Paper, opponentShape)
		assertShapesEqual(t, Rock, myShape)
	})

	t.Run("compare shapes", func(t *testing.T) {
		if Rock.Compare(Paper) != 0 {
			t.Fatal("Paper beats Rock")
		}
	})

	t.Run("calculate score", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		total := 0
		for _, line := range input {
			opponentShape, myShape := ParseInput(line)
			total += Score(opponentShape, myShape)
		}

		if total != 15 {
			t.Fatalf("expected final score 15, got %d", total)
		}
	})
}

func assertShapesEqual(t testing.TB, expected RockPaperScissorsShape, actual RockPaperScissorsShape) {
	t.Helper()
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}