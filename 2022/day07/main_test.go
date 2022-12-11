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

	t.Run("change directory up one", func(t *testing.T) {
		filesystem := NewFilesystem()
		filesystem.currentLocation = "/some/path"
		command := ChangeDirectoryCommand{Destination: ".."}
		filesystem = command.ApplyCommand(filesystem)

		if filesystem.currentLocation != "/some" {
			t.Fatalf("expected %q, got %q", "/some", filesystem.currentLocation)
		}

		filesystem = command.ApplyCommand(filesystem)

		if filesystem.currentLocation != "/" {
			t.Fatalf("expected %q, got %q", "/", filesystem.currentLocation)
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

		dir := filesystem.GetNodeAtPath("/b")
		if dir.GetNodeType() != DirectoryNode {
			t.Fatalf("expected node type %v, got %v", DirectoryNode, dir.GetNodeType())
		}

		file := filesystem.GetNodeAtPath("/b/c.dat")
		if file.GetNodeType() != FileNode {
			t.Fatalf("expected node type %v, got %v", FileNode, file.GetNodeType())
		}
		if file.GetSize() != 200 {
			t.Fatalf("expected file to have size %d, got %d", 200, file.GetSize())
		}
	})

	t.Run("list command", func(t *testing.T) {
		filesystem := NewFilesystem()
		children := []INode{
			Directory{Name: "a", Children: make(map[string]INode)},
			File{Name: "b.txt", Size: 14848514},
			File{Name: "c.dat", Size: 8504156},
			Directory{Name: "d", Children: make(map[string]INode)},
		}
		command := ListCommand{Children: children}

		filesystem = command.ApplyCommand(filesystem)

		if len(filesystem.root.Children) != 4 {
			t.Fatalf("expected %d nodes, got %d", 4, len(filesystem.root.Children))
		}
	})

	t.Run("get size of directory", func(t *testing.T) {
		dir := Directory{Name: "test", Children: map[string]INode{
			"a.txt": File{Name: "a.txt", Size: 123},
			"b.dat": File{Name: "b.dat", Size: 456},
		}}

		if dir.GetSize() != (123 + 456) {
			t.Fatalf("expected directory size to be %d, got %d", 123+456, dir.GetSize())
		}
	})

	t.Run("test sample input", func(t *testing.T) {
		filesystem := NewFilesystem()
		input := utils.ReadFile("resources/sample_input.txt")
		commandStrings := ParseInput(input)
		for _, commandString := range commandStrings {
			command := ParseCommandString(commandString)
			filesystem = command.ApplyCommand(filesystem)
		}

		e := filesystem.GetNodeAtPath("/a/e")
		if e.GetSize() != 584 {
			t.Fatalf("expected directory e to have size %d, got %d", 584, e.GetSize())
		}

		root := filesystem.root
		if root.GetSize() != 48381165 {
			t.Fatalf("expected filesystem to have total size %d, got %d", 48381165, root.GetSize())
		}

		sizes := root.GetSizeIfLessThanLimit(100000)

		total := 0
		for _, dirSize := range sizes {
			total += dirSize
		}

		if total != 95437 {
			t.Fatalf("expected %d, got %d", 95437, total)
		}
	})
}
