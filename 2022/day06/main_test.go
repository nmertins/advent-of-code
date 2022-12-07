package main

import (
	"testing"
)

func TestDay06(t *testing.T) {
	t.Run("characters are unique", func(t *testing.T) {
		testCases := []struct {
			Input          string
			ExpectedOutput bool
		}{
			{
				Input:          "aaaa",
				ExpectedOutput: false,
			},
			{
				Input:          "abcd",
				ExpectedOutput: true,
			},
		}

		for _, testCase := range testCases {
			actual := AllCharactersUnique(testCase.Input)
			if actual != testCase.ExpectedOutput {
				t.Fatalf("unique characters in %q, expected %t got %t", testCase.Input, testCase.ExpectedOutput, actual)
			}
		}
	})
}
