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
}

func assertShapesEqual(t testing.TB, expected RockPaperScissorsShape, actual RockPaperScissorsShape) {
	t.Helper()
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}