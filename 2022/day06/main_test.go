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
			{
				Input:          "mjqj",
				ExpectedOutput: false,
			},
			{
				Input:          "jpqm",
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

	t.Run("find unique characters", func(t *testing.T) {
		testCases := []struct {
			Input    string
			Expected string
		}{
			{
				Input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				Expected: "jpqm",
			},
			{
				Input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
				Expected: "vwbj",
			},
			{
				Input:    "nppdvjthqldpwncqszvftbrmjlhg",
				Expected: "pdvj",
			},
		}

		for _, testCase := range testCases {
			actual := FindMarker(testCase.Input)
			if testCase.Expected != actual {
				t.Fatalf("expected %q got %q", testCase.Expected, actual)
			}
		}
	})

	t.Run("find marker index", func(t *testing.T) {
		testCases := []struct {
			Input    string
			Expected int
		}{
			{
				Input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				Expected: 7,
			},
			{
				Input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
				Expected: 5,
			},
			{
				Input:    "nppdvjthqldpwncqszvftbrmjlhg",
				Expected: 6,
			},
		}

		for _, testCase := range testCases {
			marker := FindMarker(testCase.Input)
			index := GetMarkerIndex(testCase.Input, marker)
			if testCase.Expected != index {
				t.Fatalf("case %q: expected %d got %d", testCase.Input, testCase.Expected, index)
			}
		}
	})
}
