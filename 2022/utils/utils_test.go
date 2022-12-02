package utils

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("day01 sample input", func(t *testing.T) {
		input := ReadFile("resources/day01_sample_input.txt")

		expectedLength := 14
		actualLength := len(input)
		if len(input) != 14 {
			t.Fatalf("day01_sample_input.txt should have %d lines, instead found %d", expectedLength, actualLength)
		}
		if input[7] != "6000" {
			t.Fatalf("expected value of '6000' on line 8, got %s", input[7])
		}
	})

	t.Run("day02 sample input", func(t *testing.T) {
		input := ReadFile("resources/day02_sample_input.txt")

		expectedLength := 3
		actualLength := len(input)
		if actualLength != expectedLength {
			t.Fatalf("day02_sample_input.txt should have %d lines, instead found %d", expectedLength, actualLength)
		}

		expectedLine := "B X"
		actualLine := input[1]
		if actualLine != expectedLine {
			t.Fatalf("expected value of %q, but got %q", expectedLine, actualLine)
		}
	})
}
