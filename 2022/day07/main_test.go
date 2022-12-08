package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay07(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		filesystem := ParseInput(input)
		if filesystem.tree["/"] == nil {
			t.Fatalf("Filesystem doesn't have a root")
		}
	})

	t.Run("parse command", func(t *testing.T) {
		testCases := []struct {
			CommandString       []string
			ExpectedCommandType int
		}{
			{
				CommandString:       []string{"$ cd /"},
				ExpectedCommandType: ChangeDirectory,
			},
			{
				CommandString:       []string{"$ ls", "dir a", "14848514 b.txt"},
				ExpectedCommandType: List,
			},
		}

		for _, testCase := range testCases {
			command := ParseCommandString(testCase.CommandString)
			actualCommandType := command.GetCommandType()
			if actualCommandType != testCase.ExpectedCommandType {
				t.Fatalf("expected command type %d, got %d", testCase.ExpectedCommandType, actualCommandType)
			}

			if actualCommandType == ChangeDirectory {
				c := command.(ChangeDirectoryCommand)
				if c.Destination != "/" {
					t.Fatalf("expected destination %q, got %q", "/", c.Destination)
				}
			}

			if actualCommandType == List {
				c := command.(ListCommand)
				if len(c.Children) != 2 {
					t.Fatalf("expected %d child nodes, got %d", 2, len(c.Children))
				}
			}
		}
	})
}
