package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay07(t *testing.T) {
	t.Run("group input by command", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		commands := ParseInput(input)
		expected := 10
		if len(commands) != expected {
			t.Fatalf("expected %d commands, got %d", expected, len(commands))
		}
	})

	t.Run("parse command", func(t *testing.T) {
		testCases := []struct {
			commandString []string
			expected      CommandType
		}{
			{
				commandString: []string{"$ cd /"},
				expected:      ChangeDirectory,
			},
			{
				commandString: []string{"$ ls", "dir a", "14848514 b.txt"},
				expected:      List,
			},
		}

		for _, testCase := range testCases {
			command := ParseCommandString(testCase.commandString)
			actual := command.GetCommandType()
			if actual != testCase.expected {
				t.Fatalf("expected command type %d, got %d", testCase.expected, actual)
			}

			if actual == ChangeDirectory {
				c := command.(ChangeDirectoryCommand)
				if c.Destination != "/" {
					t.Fatalf("expected destination %q, got %q", "/", c.Destination)
				}
			}

			if actual == List {
				c := command.(ListCommand)
				if len(c.Children) != 2 {
					t.Fatalf("expected %d child nodes, got %d", 2, len(c.Children))
				}
			}
		}
	})
}
