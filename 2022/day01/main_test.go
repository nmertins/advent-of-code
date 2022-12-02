package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay01(t *testing.T) {
	t.Run("parse input file", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		if len(input) != 14 {
			t.Fatalf("sample_input.txt should have 14 lines, instead found %d", len(input))
		}
		if input[7] != "6000" {
			t.Fatalf("expected value of '6000' on line 8, got %s", input[7])
		}

		totals := ParseInput(input)
		if totals[0] != 6000 {
			t.Fatalf("expected first total to be 6000, got %d", totals[0])
		}

		max := Max(totals)
		if max != 24000 {
			t.Fatalf("expected max value to be 24000, got %d", max)
		}
	})
}

func ReadFile(s string) {
	panic("unimplemented")
}
