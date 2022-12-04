package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay04(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		pairs := ParseInput(input)
		if len(pairs) != 6 {
			t.Fatalf("expected 6 elements, got %d", len(pairs))
		}
		if pairs[0].elves[0].startID != 2 {
			t.Fatalf("expected start ID to be 2, got %d", pairs[0].elves[0].startID)
		}
		if pairs[0].elves[0].endID != 4 {
			t.Fatalf("expected end ID to be 4, got %d", pairs[0].elves[0].startID)
		}
	})

	t.Run("test if cleaning assignments overlap", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		pairs := ParseInput(input)

		if pairs[0].AssignmentsCompletelyOverlap() {
			t.Fatalf("expected no overlap")
		}
		if !pairs[3].AssignmentsCompletelyOverlap() {
			t.Fatalf("expected assignments to overlap")
		}
	})

	t.Run("test sample input: part 1", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		pairs := ParseInput(input)
		total := 0
		for _, pair := range pairs {
			if pair.AssignmentsCompletelyOverlap() {
				total++
			}
		}

		if total != 2 {
			t.Fatalf("expected 2 overlapping assignments, got %d", total)
		}
	})

	t.Run("test sample input: part 2", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		pairs := ParseInput(input)
		total := 0
		for _, pair := range pairs {
			if pair.AssignmentsOverlap() {
				total++
			}
		}

		if total != 4 {
			t.Fatalf("expected 4 overlapping assignments, got %d", total)
		}
	})
}
