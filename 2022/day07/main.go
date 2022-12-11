package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nmertins/advent-of-code/2022/utils"
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

func (f Filesystem) GetNodeAtPath(path string) (node INode) {
	if path == "/" {
		return f.root
	}

	splitPath := strings.Split(path, "/")[1:]
	node = f.root
	for {
		if len(splitPath) == 0 {
			break
		}
		var next string
		next, splitPath = splitPath[0], splitPath[1:]
		node = node.(Directory).Children[next]
	}

	return node
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
	total := 0
	for _, val := range d.Children {
		total += val.GetSize()
	}
	return total
}

func (d Directory) GetSizeIfLessThanLimit(sizeLimit int) []int {
	size := d.GetSize()
	sizes := make([]int, 0)
	if size <= sizeLimit {
		sizes = append(sizes, size)
	}

	for _, child := range d.Children {
		if child.GetNodeType() == DirectoryNode {
			sizes = append(sizes, child.(Directory).GetSizeIfLessThanLimit(sizeLimit)...)
		}
	}

	return sizes
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
		return filesystem
	}

	if c.Destination == ".." {
		if filesystem.currentLocation == "/" {
			return filesystem
		}

		path := strings.Split(strings.Trim(filesystem.currentLocation, "/"), "/")
		newPath := strings.Join(path[:len(path)-1], "/")
		filesystem.currentLocation = "/" + newPath
		return filesystem
	}

	if strings.LastIndex(filesystem.currentLocation, "/") != len(filesystem.currentLocation)-1 {
		filesystem.currentLocation += "/"
	}
	filesystem.currentLocation += c.Destination

	return filesystem
}

type ListCommand struct {
	Children []INode
}

func (l ListCommand) GetCommandType() CommandType {
	return List
}

func (l ListCommand) ApplyCommand(filesystem Filesystem) Filesystem {
	pathToAddChildren := filesystem.currentLocation
	dir := filesystem.GetNodeAtPath(pathToAddChildren).(Directory)

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
	commandStrings[len(commandStrings)-1] = input[commandIndexes[len(commandIndexes)-1]:]

	return commandStrings
}

func main() {
	filesystem := NewFilesystem()
	input := utils.ReadFile("resources/input.txt")
	commandStrings := ParseInput(input)
	for _, commandString := range commandStrings {
		command := ParseCommandString(commandString)
		filesystem = command.ApplyCommand(filesystem)
	}

	sizeLimit := 100000
	sizes := filesystem.root.GetSizeIfLessThanLimit(sizeLimit)

	total := 0
	for _, dirSize := range sizes {
		total += dirSize
	}

	fmt.Println(total)
}
