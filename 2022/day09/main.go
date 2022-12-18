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

func (r *Rope) MoveHead(motion Motion) {
	for i := 0; i < motion.Distance; i++ {
		switch motion.Direction {
		case Right:
			r.Head[0]++
		case Up:
			r.Head[1]++
		case Left:
			r.Head[0]--
		case Down:
			r.Head[1]--
		}
		r.UpdateTail()
	}
}

func (r *Rope) UpdateTail() {
	xHead := r.Head[0]
	yHead := r.Head[1]

	xTail := r.Tail[0]
	yTail := r.Tail[1]

	switch xDiff, yDiff := xHead-xTail, yHead-yTail; {
	// if head is 2 spaces to the right,
	// then tail should move right 1 space
	case xDiff > 1 && yDiff == 0:
		r.Tail[0]++
	// if head is 2 spaces below,
	// then tail should move down 1 space
	case yDiff < -1 && xDiff == 0:
		r.Tail[1]--
	// if head is 2 spaces to the left,
	// then tail should move left 1 space
	case xDiff < -1 && yDiff == 0:
		r.Tail[0]--
	// if head is 2 spaces above,
	// then tail should move up 1 space
	case yDiff > 1 && xDiff == 0:
		r.Tail[1]++
	// if head is 2 spaces above and 1 space right,
	// then tail should move diagonally
	case xDiff > 0 && yDiff > 1:
		r.Tail[0]++
		r.Tail[1]++
	// if head is 2 spaces right and 1 space above,
	// then tail should move diagonally
	case xDiff > 1 && yDiff > 0:
		r.Tail[0]++
		r.Tail[1]++
	// if head is 2 spaces above and 1 space left,
	// then tail should move diagonally
	case xDiff < 0 && yDiff > 1:
		r.Tail[0]--
		r.Tail[1]++
	// if head is 2 spaces left and 1 space above,
	// then tail should move diagonally
	case xDiff < -1 && yDiff > 0:
		r.Tail[0]--
		r.Tail[1]++
	// if head is 2 spaces below and 1 space right,
	// then tail should move diagonally
	case xDiff > 0 && yDiff < -1:
		r.Tail[0]++
		r.Tail[1]--
	// if head is 2 spaces right and 1 space below,
	// then tail should move diagonally
	case xDiff > 1 && yDiff < 0:
		r.Tail[0]++
		r.Tail[1]--
	// if head is 2 spaces below and 1 space left,
	// then tail should move diagonally
	case xDiff < 0 && yDiff < -1:
		r.Tail[0]--
		r.Tail[1]--
	// if head is 2 spaces left and 1 space below,
	// then tail should move diagonally
	case xDiff < -1 && yDiff < 0:
		r.Tail[0]--
		r.Tail[1]--
	}
}

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

func GetUnique(points [][2]int) [][2]int {

	visitedPoints := make(map[[2]int]bool)
	ret := make([][2]int, 0)

	for _, point := range points {
		if _, found := visitedPoints[point]; !found {
			visitedPoints[point] = true
			ret = append(ret, point)
		}
	}

	return ret
}
