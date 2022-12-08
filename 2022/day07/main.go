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

type INode interface {
	GetName() string
}

type Filesystem struct {
	tree            map[string]INode
	currentLocation []string
}

type Directory struct {
	Name string
}

func (d Directory) GetName() string {
	return d.Name
}

type File struct {
	Name string
	Size int
}

func (f File) GetName() string {
	return f.Name
}

type Command interface {
	GetCommandType() CommandType
}

type ChangeDirectoryCommand struct {
	Destination string
}

func (c ChangeDirectoryCommand) GetCommandType() CommandType {
	return ChangeDirectory
}

type ListCommand struct {
	Children []INode
}

func (l ListCommand) GetCommandType() CommandType {
	return List
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
				node = Directory{Name: nodeSplit[1]}
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
