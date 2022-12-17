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
		testCases := []struct {
			rope        Rope
			description string
		}{
			{Rope{Head: [2]int{0, 0}, Tail: [2]int{0, 0}}, "head covers tail, no movement"},
			{Rope{Head: [2]int{1, 0}, Tail: [2]int{0, 0}}, "head one space right of tail, no movement"},
		}

		for _, testCase := range testCases {
			testCase.rope.UpdateTail()
			if testCase.rope.Head != [2]int{0, 0} && testCase.rope.Tail != [2]int{0, 0} {
				t.Errorf("Test case: %q\n", testCase.description)
				t.Errorf("expected head at %v, got %v\n", [2]int{0, 0}, testCase.rope.Head)
				t.Errorf("expected tail at %v, got %v\n", [2]int{0, 0}, testCase.rope.Tail)
			}
		}
	})
}
