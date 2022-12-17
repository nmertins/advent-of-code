package main

import (
	"strconv"
	"strings"
)

type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
)

type Motion struct {
	Direction Direction
	Distance  int
}

type Rope struct {
	Head [2]int
	Tail [2]int
}

func (r *Rope) UpdateTail() {}

func ParseInput(input []string) []Motion {
	motions := make([]Motion, len(input))
	for i, motionString := range input {
		split := strings.Split(motionString, " ")
		var direction Direction
		switch split[0] {
		case "R":
			direction = Right
		case "L":
			direction = Left
		case "U":
			direction = Up
		case "D":
			direction = Down
		}
		distance, _ := strconv.Atoi(split[1])
		motions[i] = Motion{Direction: direction, Distance: distance}
	}
	return motions
}
