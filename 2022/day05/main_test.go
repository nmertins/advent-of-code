package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay05(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		stacks, _ := ParseInput(input)
		if len(stacks) != 3 {
			t.Fatalf("expected 3 stacks, got %d", len(stacks))
		}

		assertStackSize(t, stacks[0], 2)
		assertStackSize(t, stacks[1], 3)
		assertStackSize(t, stacks[2], 1)
	})
}

func assertStackSize(t testing.TB, stack Stack, expectedSize int) {
	t.Helper()
	if len(stack) != expectedSize {
		t.Fatalf("expected stack to have %d crates, got %d", expectedSize, len(stack))
	}
}
