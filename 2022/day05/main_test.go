package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay05(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		stacks, procedure := ParseInput(input)
		if len(stacks) != 3 {
			t.Fatalf("expected 3 stacks, got %d", len(stacks))
		}

		assertStackSize(t, stacks[0], 2)
		assertStackSize(t, stacks[1], 3)
		assertStackSize(t, stacks[2], 1)

		if stacks[0][0] != "N" {
			t.Fatalf("expected crate \"N\", got %q", stacks[0][0])
		}

		if stacks[1][2] != "M" {
			t.Fatalf("expected crate \"M\", got %q", stacks[0][0])
		}

		if procedure[0].cratesToMove != 1 {
			t.Fatalf("expected step to move 1 crate, got %d", procedure[0].cratesToMove)
		}
		if procedure[0].fromStack != 2 {
			t.Fatalf("expected step to move crate(s) from stack 2, got %d", procedure[0].fromStack)
		}
		if procedure[0].toStack != 1 {
			t.Fatalf("expected step to move crate(s) to stack 1, got %d", procedure[0].toStack)
		}
	})
}

func assertStackSize(t testing.TB, stack Stack, expectedSize int) {
	t.Helper()
	if len(stack) != expectedSize {
		t.Fatalf("expected stack to have %d crates, got %d", expectedSize, len(stack))
	}
}
