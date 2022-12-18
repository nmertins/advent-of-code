package main

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay09(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		motions := ParseInput(input)
		if len(motions) != 8 {
			t.Errorf("expected %d motions in sample input, got %d", 8, len(motions))
		}

		if motions[0].Direction != Right {
			t.Errorf("expected direction %v, got %v", Right, motions[0].Direction)
		}
		if motions[0].Distance != 4 {
			t.Errorf("expected distance %d, got %d", 4, motions[0].Distance)
		}
	})

	t.Run("motions move the head", func(t *testing.T) {
		testCases := []struct {
			motions  []Motion
			expected Rope
		}{
			{
				motions:  []Motion{{Direction: Right, Distance: 1}},
				expected: Rope{Head: [2]int{1, 0}, Tail: [2]int{0, 0}},
			},
			{
				motions:  []Motion{{Direction: Right, Distance: 3}},
				expected: Rope{Head: [2]int{3, 0}, Tail: [2]int{2, 0}},
			},
			{
				motions:  []Motion{{Direction: Up, Distance: 1}},
				expected: Rope{Head: [2]int{0, 1}, Tail: [2]int{0, 0}},
			},
			{
				motions:  []Motion{{Direction: Left, Distance: 1}},
				expected: Rope{Head: [2]int{-1, 0}, Tail: [2]int{0, 0}},
			},
			{
				motions:  []Motion{{Direction: Down, Distance: 1}},
				expected: Rope{Head: [2]int{0, -1}, Tail: [2]int{0, 0}},
			},
			{
				motions: []Motion{
					{Direction: Right, Distance: 4},
					{Direction: Up, Distance: 4},
				},
				expected: Rope{Head: [2]int{4, 4}, Tail: [2]int{4, 3}},
			},
		}

		for _, testCase := range testCases {
			rope := CreateRope()
			for _, motion := range testCase.motions {
				rope.MoveHead(motion)
			}

			if !rope.Equals(testCase.expected) {
				t.Errorf("expected rope at %v, got %v", testCase.expected, rope)
			}
		}
	})

	t.Run("tail follows head", func(t *testing.T) {
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
			if !testCase.initial.Equals(testCase.expected) {
				t.Errorf("Test case: %q\n", testCase.description)
				t.Errorf("expected rope at %v, got %v\n", testCase.expected, testCase.initial)
			}
		}
	})

	t.Run("track spaces visited", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		motions := ParseInput(input)
		rope := CreateRope()
		for _, motion := range motions {
			rope.MoveHead(motion)
		}

		expected := Rope{Head: [2]int{2, 2}, Tail: [2]int{1, 2}}
		if !rope.Equals(expected) {
			t.Fatalf("expected rope at %v, got %v", expected, rope)
		}

		uniqueVisited := GetUnique(rope.visited)
		expectedLength := 13
		if len(uniqueVisited) != expectedLength {
			t.Errorf("expected tail to visit %d unique spaces, got %d", expectedLength, len(uniqueVisited))
		}
	})
}

func TestUniqueArrays(t *testing.T) {
	points := [][2]int{
		{2, 0},
		{2, 2},
		{1, 0},
		{1, 2},
		{1, 0},
		{2, 1},
		{0, 2},
		{1, 2},
		{0, 2},
		{2, 2},
	}

	expected := [][2]int{
		{2, 0},
		{2, 2},
		{1, 0},
		{1, 2},
		{2, 1},
		{0, 2},
	}

	unique := GetUnique(points)

	if !reflect.DeepEqual(expected, unique) {
		t.Errorf("expected %v, got %v", expected, unique)
	}
}

func createRandomPoints(numberOfPoints int, maxValue int) [][2]int {
	ret := make([][2]int, numberOfPoints)
	for i := 0; i < numberOfPoints; i++ {
		ret[i] = [2]int{rand.Intn(maxValue), rand.Intn(maxValue)}
	}

	return ret
}
