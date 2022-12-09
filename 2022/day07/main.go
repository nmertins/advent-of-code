package main

import (
	"strconv"
	"strings"
)

type CommandType int

const (
	ChangeDirectory CommandType = iota
	List
)

type INodeType int

const (
	DirectoryNode INodeType = iota
	FileNode
)

type INode interface {
	GetSize() int
	GetName() string
	GetNodeType() INodeType
}

type Filesystem struct {
	root            Directory
	currentLocation string
}

func NewFilesystem() Filesystem {
	return Filesystem{
		root:            Directory{Name: "", Children: make(map[string]INode)},
		currentLocation: "/",
	}
}

type Directory struct {
	Name     string
	Children map[string]INode
}

func (d Directory) GetName() string {
	return d.Name
}

func (d Directory) GetSize() int {
	return 0
}

func (d Directory) GetNodeType() INodeType {
	return DirectoryNode
}

type File struct {
	Name string
	Size int
}

func (f File) GetName() string {
	return f.Name
}

func (f File) GetSize() int {
	return f.Size
}

func (f File) GetNodeType() INodeType {
	return FileNode
}

type Command interface {
	GetCommandType() CommandType
	ApplyCommand(Filesystem) Filesystem
}

type ChangeDirectoryCommand struct {
	Destination string
}

func (c ChangeDirectoryCommand) GetCommandType() CommandType {
	return ChangeDirectory
}

func (c ChangeDirectoryCommand) ApplyCommand(filesystem Filesystem) Filesystem {
	if strings.Index(c.Destination, "/") == 0 {
		filesystem.currentLocation = c.Destination
	} else {
		if strings.LastIndex(filesystem.currentLocation, "/") != len(filesystem.currentLocation)-1 {
			filesystem.currentLocation += "/"
		}
		filesystem.currentLocation += c.Destination
	}

	return filesystem
}

type ListCommand struct {
	Children []INode
}

func (l ListCommand) GetCommandType() CommandType {
	return List
}

func (l ListCommand) ApplyCommand(filesystem Filesystem) Filesystem {
	path := strings.Split(filesystem.currentLocation, "/")[1:]
	dir := filesystem.root

	for {
		if len(path) == 0 {
			break
		}

		dir = dir.Children[path[0]].(Directory)
		_, path = path[0], path[1:]
	}

	for _, child := range l.Children {
		dir.Children[child.GetName()] = child
	}

	return filesystem
}

func ParseCommandString(commandString []string) Command {
	split := strings.Split(commandString[0], " ")
	var command Command
	switch split[1] {
	case "cd":
		command = ChangeDirectoryCommand{
			Destination: split[2],
		}
	case "ls":
		childNodes := make([]INode, 0)
		for _, nodeString := range commandString[1:] {
			nodeSplit := strings.Split(nodeString, " ")
			var node INode
			if nodeSplit[0] == "dir" {
				node = Directory{Name: nodeSplit[1], Children: make(map[string]INode, 0)}
			} else {
				size, _ := strconv.Atoi(nodeSplit[0])
				node = File{Name: nodeSplit[1], Size: size}
			}

			childNodes = append(childNodes, node)
		}
		command = ListCommand{Children: childNodes}
	}

	return command
}

func ParseInput(input []string) [][]string {
	commandIndexes := make([]int, 0)
	for i, line := range input {
		if strings.Index(line, "$") == 0 {
			commandIndexes = append(commandIndexes, i)
		}
	}

	commandStrings := make([][]string, len(commandIndexes))
	for i := 0; i < len(commandIndexes)-1; i++ {
		startIndex := commandIndexes[i]
		endIndex := commandIndexes[i+1]
		commandStrings[i] = input[startIndex:endIndex]
	}

	return commandStrings
}
