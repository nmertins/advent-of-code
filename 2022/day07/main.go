package main

import (
	"strconv"
	"strings"
)

const (
	ChangeDirectory = iota
	List            = iota
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
	GetCommandType() int
}

type ChangeDirectoryCommand struct {
	Destination string
}

func (c ChangeDirectoryCommand) GetCommandType() int {
	return ChangeDirectory
}

type ListCommand struct {
	Children []INode
}

func (l ListCommand) GetCommandType() int {
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
		command = ListCommand{}
		for _, nodeString := range commandString[1:] {
			nodeSplit := strings.Split(nodeString, " ")
			var node INode
			if nodeSplit[0] == "dir" {
				node = Directory{Name: nodeSplit[1]}
			} else {
				size, _ := strconv.Atoi(nodeSplit[0])
				node = File{Name: nodeSplit[1], Size: size}
			}

			//TODO add node to ListCommand
		}
	}

	return command
}

func ParseInput(input []string) Filesystem {
	filesystem := Filesystem{
		tree:            map[string]INode{"/": Directory{}},
		currentLocation: []string{"/"},
	}

	filesystem.tree["/"] = Directory{}

	return filesystem
}
