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

	t.Run("change directory absolute path", func(t *testing.T) {
		filesystem := NewFilesystem()
		filesystem.currentLocation = "/some/arbitrary/path"
		command := ChangeDirectoryCommand{Destination: "/another/path"}
		filesystem = command.ApplyCommand(filesystem)

		if filesystem.currentLocation != "/another/path" {
			t.Fatal("absolute path should replace existing path")
		}
	})

	t.Run("change directory relative path", func(t *testing.T) {
		filesystem := NewFilesystem()
		command := ChangeDirectoryCommand{Destination: "relative/path"}
		filesystem = command.ApplyCommand(filesystem)

		if filesystem.currentLocation != "/relative/path" {
			t.Fatalf("expected %q, got %q", "/relative/path", filesystem.currentLocation)
		}

		filesystem.currentLocation = "/home"
		filesystem = command.ApplyCommand(filesystem)

		if filesystem.currentLocation != "/home/relative/path" {
			t.Fatalf("expected %q, got %q", "/home/relative/path", filesystem.currentLocation)
		}
	})

	t.Run("get node on filesystem", func(t *testing.T) {
		filesystem := NewFilesystem()

		fileA := File{Name: "a.txt", Size: 1000}
		dirB := Directory{Name: "b", Children: make(map[string]INode)}
		filesystem.root.Children[fileA.Name] = fileA
		filesystem.root.Children[dirB.Name] = dirB
		fileC := File{Name: "c.dat", Size: 200}
		dirB.Children[fileC.Name] = fileC

		if len(filesystem.root.Children) != 2 {
			t.Fatalf("expected %d child nodes in the root, got %d", 2, len(filesystem.root.Children))
		}
	})

	// t.Run("list command", func(t *testing.T) {
	// 	filesystem := NewFilesystem()
	// 	children := []INode{
	// 		Directory{Name: "a"},
	// 		File{Name: "b.txt", Size: 14848514},
	// 		File{Name: "c.dat", Size: 8504156},
	// 		Directory{Name: "d"},
	// 	}
	// 	command := ListCommand{Children: children}

	// 	filesystem = command.ApplyCommand(filesystem)

	// 	if len(filesystem.root.Children) != 4 {
	// 		t.Fatalf("expected %d nodes, got %d", 4, len(filesystem.root.Children))
	// 	}
	// })
}
