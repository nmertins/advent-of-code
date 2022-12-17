package main

import (
	"reflect"
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
			description string
			initial     Rope
			expected    Rope
		}{
			// no movement
			{
				"head covers tail, no movement",
				Rope{Head: [2]int{0, 0}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{0, 0}, Tail: [2]int{0, 0}},
			},
			{
				"head one space right of tail, no movement",
				Rope{Head: [2]int{1, 0}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{1, 0}, Tail: [2]int{0, 0}},
			},
			{
				"head one space left of tail, no movement",
				Rope{Head: [2]int{-1, 0}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-1, 0}, Tail: [2]int{0, 0}},
			},
			{
				"head one space above tail, no movement",
				Rope{Head: [2]int{0, 1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{0, 1}, Tail: [2]int{0, 0}},
			},
			{
				"head one space below tail, no movement",
				Rope{Head: [2]int{0, -1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{0, -1}, Tail: [2]int{0, 0}},
			},
			{
				"head upper right diagonal from tail, no movement",
				Rope{Head: [2]int{1, 1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{1, 1}, Tail: [2]int{0, 0}},
			},
			{
				"head upper left diagonal from tail, no movement",
				Rope{Head: [2]int{-1, 1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-1, 1}, Tail: [2]int{0, 0}},
			},
			{
				"head lower right diagonal from tail, no movement",
				Rope{Head: [2]int{1, -1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{1, -1}, Tail: [2]int{0, 0}},
			},
			{
				"head lower left diagonal from tail, no movement",
				Rope{Head: [2]int{-1, -1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-1, -1}, Tail: [2]int{0, 0}},
			},
			// linear movement
			{
				"head two spaces right of tail, tail moves right one space",
				Rope{Head: [2]int{2, 0}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{2, 0}, Tail: [2]int{1, 0}},
			},
			{
				"head two spaces below tail, tail moves down one space",
				Rope{Head: [2]int{0, -2}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{0, -2}, Tail: [2]int{0, -1}},
			},
			{
				"head two spaces left of tail, tail moves left one space",
				Rope{Head: [2]int{-2, 0}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-2, 0}, Tail: [2]int{-1, 0}},
			},
			{
				"head two spaces above tail, tail moves up one space",
				Rope{Head: [2]int{0, 2}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{0, 2}, Tail: [2]int{0, 1}},
			},
			// diagonal movement
			// upper right
			{
				"head two spaces above and one space right of tail, tail moves diagonally",
				Rope{Head: [2]int{1, 2}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{1, 2}, Tail: [2]int{1, 1}},
			},
			{
				"head two spaces right and one space above tail, tail moves diagonally",
				Rope{Head: [2]int{2, 1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{2, 1}, Tail: [2]int{1, 1}},
			},
			// upper left
			{
				"head two spaces above and one space left of tail, tail moves diagonally",
				Rope{Head: [2]int{-1, 2}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-1, 2}, Tail: [2]int{-1, 1}},
			},
			{
				"head two spaces left and one space above tail, tail moves diagonally",
				Rope{Head: [2]int{-2, 1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-2, 1}, Tail: [2]int{-1, 1}},
			},
			// lower right
			{
				"head two spaces below and one space right of tail, tail moves diagonally",
				Rope{Head: [2]int{1, -2}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{1, -2}, Tail: [2]int{1, -1}},
			},
			{
				"head two spaces right and one space below tail, tail moves diagonally",
				Rope{Head: [2]int{2, -1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{2, -1}, Tail: [2]int{1, -1}},
			},
			// lower left
			{
				"head two spaces below and one space left of tail, tail moves diagonally",
				Rope{Head: [2]int{-1, -2}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-1, -2}, Tail: [2]int{-1, -1}},
			},
			{
				"head two spaces left and one space below tail, tail moves diagonally",
				Rope{Head: [2]int{-2, -1}, Tail: [2]int{0, 0}},
				Rope{Head: [2]int{-2, -1}, Tail: [2]int{-1, -1}},
			},
		}

		for _, testCase := range testCases {
			testCase.initial.UpdateTail()
			if !reflect.DeepEqual(testCase.initial, testCase.expected) {
				t.Errorf("Test case: %q\n", testCase.description)
				t.Errorf("expected rope at %v, got %v\n", testCase.expected, testCase.initial)
			}
		}
	})
}
