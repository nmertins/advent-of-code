package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay09(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		motions := ParseInput(input)
		if len(motions) != 8 {
			t.Fatalf("expected %d motions in sample input, got %d", 8, len(motions))
		}

		if motions[0].Direction != Right {
			t.Fatalf("expected direction %v, got %v", Right, motions[0].Direction)
		}
		if motions[0].Distance != 4 {
			t.Fatalf("expected distance %d, got %d", 4, motions[0].Distance)
		}
	})

	t.Run("test rope movement", func(t *testing.T) {
		rope := Rope{Head: {0, 0}, Tail: {0, 0}}
	})
}
