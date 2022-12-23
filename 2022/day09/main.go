package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nmertins/advent-of-code/2022/utils"
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
	head    [2]int
	tail    [2]int
	visited [][2]int
}

type IRope interface {
	MoveHead(Motion)
	GetVisitedSpaces() [][2]int
	GetHead() [2]int
	GetTail() [2]int
}

func CreateRope() IRope {
	return &Rope{head: [2]int{0, 0}, tail: [2]int{0, 0}, visited: make([][2]int, 0)}
}

func (r *Rope) MoveHead(motion Motion) {
	for i := 0; i < motion.Distance; i++ {
		switch motion.Direction {
		case Right:
			r.head[0]++
		case Up:
			r.head[1]++
		case Left:
			r.head[0]--
		case Down:
			r.head[1]--
		}
		r.UpdateTail()
		r.visited = append(r.visited, r.tail)
	}
}

func (r Rope) GetVisitedSpaces() [][2]int {
	return r.visited
}

func (r Rope) GetHead() [2]int {
	return r.head
}

func (r Rope) GetTail() [2]int {
	return r.tail
}

type Knot struct {
	location [2]int
	child    *Knot
}

type RopeN struct {
	head *Knot
	tail *Knot
}

func (r *RopeN) AddKnot() {
	knot := &Knot{
		location: [2]int{0, 0},
		child:    nil,
	}

	if r.head == nil {
		r.head = knot
	}

	if r.tail == nil {
		r.tail = knot
	}

	r.tail.child = knot
	r.tail = knot
}

func (k *Knot) UpdateChild(parentLocation [2]int) {
	k.location = GetNewPosition(k.location, parentLocation)

	if k.child != nil {
		k.child.UpdateChild(k.location)
	}
}

func CreateRopeN(length int) *RopeN {
	ret := &RopeN{}
	for i := 0; i < length; i++ {
		ret.AddKnot()
	}

	return ret
}

func GetNewPosition(current [2]int, parent [2]int) [2]int {
	xHead := parent[0]
	yHead := parent[1]

	xTail := current[0]
	yTail := current[1]

	ret := [2]int{xTail, yTail}

	switch xDiff, yDiff := xHead-xTail, yHead-yTail; {
	// if head is 2 spaces to the right,
	// then tail should move right 1 space
	case xDiff > 1 && yDiff == 0:
		ret[0]++
	// if head is 2 spaces below,
	// then tail should move down 1 space
	case yDiff < -1 && xDiff == 0:
		ret[1]--
	// if head is 2 spaces to the left,
	// then tail should move left 1 space
	case xDiff < -1 && yDiff == 0:
		ret[0]--
	// if head is 2 spaces above,
	// then tail should move up 1 space
	case yDiff > 1 && xDiff == 0:
		ret[1]++
	// if head is 2 spaces above and 1 space right,
	// then tail should move diagonally
	case xDiff > 0 && yDiff > 1:
		ret[0]++
		ret[1]++
	// if head is 2 spaces right and 1 space above,
	// then tail should move diagonally
	case xDiff > 1 && yDiff > 0:
		ret[0]++
		ret[1]++
	// if head is 2 spaces above and 1 space left,
	// then tail should move diagonally
	case xDiff < 0 && yDiff > 1:
		ret[0]--
		ret[1]++
	// if head is 2 spaces left and 1 space above,
	// then tail should move diagonally
	case xDiff < -1 && yDiff > 0:
		ret[0]--
		ret[1]++
	// if head is 2 spaces below and 1 space right,
	// then tail should move diagonally
	case xDiff > 0 && yDiff < -1:
		ret[0]++
		ret[1]--
	// if head is 2 spaces right and 1 space below,
	// then tail should move diagonally
	case xDiff > 1 && yDiff < 0:
		ret[0]++
		ret[1]--
	// if head is 2 spaces below and 1 space left,
	// then tail should move diagonally
	case xDiff < 0 && yDiff < -1:
		ret[0]--
		ret[1]--
	// if head is 2 spaces left and 1 space below,
	// then tail should move diagonally
	case xDiff < -1 && yDiff < 0:
		ret[0]--
		ret[1]--
	}

	return ret
}

func (r *Rope) UpdateTail() {
	xHead := r.head[0]
	yHead := r.head[1]

	xTail := r.tail[0]
	yTail := r.tail[1]

	switch xDiff, yDiff := xHead-xTail, yHead-yTail; {
	// if head is 2 spaces to the right,
	// then tail should move right 1 space
	case xDiff > 1 && yDiff == 0:
		r.tail[0]++
	// if head is 2 spaces below,
	// then tail should move down 1 space
	case yDiff < -1 && xDiff == 0:
		r.tail[1]--
	// if head is 2 spaces to the left,
	// then tail should move left 1 space
	case xDiff < -1 && yDiff == 0:
		r.tail[0]--
	// if head is 2 spaces above,
	// then tail should move up 1 space
	case yDiff > 1 && xDiff == 0:
		r.tail[1]++
	// if head is 2 spaces above and 1 space right,
	// then tail should move diagonally
	case xDiff > 0 && yDiff > 1:
		r.tail[0]++
		r.tail[1]++
	// if head is 2 spaces right and 1 space above,
	// then tail should move diagonally
	case xDiff > 1 && yDiff > 0:
		r.tail[0]++
		r.tail[1]++
	// if head is 2 spaces above and 1 space left,
	// then tail should move diagonally
	case xDiff < 0 && yDiff > 1:
		r.tail[0]--
		r.tail[1]++
	// if head is 2 spaces left and 1 space above,
	// then tail should move diagonally
	case xDiff < -1 && yDiff > 0:
		r.tail[0]--
		r.tail[1]++
	// if head is 2 spaces below and 1 space right,
	// then tail should move diagonally
	case xDiff > 0 && yDiff < -1:
		r.tail[0]++
		r.tail[1]--
	// if head is 2 spaces right and 1 space below,
	// then tail should move diagonally
	case xDiff > 1 && yDiff < 0:
		r.tail[0]++
		r.tail[1]--
	// if head is 2 spaces below and 1 space left,
	// then tail should move diagonally
	case xDiff < 0 && yDiff < -1:
		r.tail[0]--
		r.tail[1]--
	// if head is 2 spaces left and 1 space below,
	// then tail should move diagonally
	case xDiff < -1 && yDiff < 0:
		r.tail[0]--
		r.tail[1]--
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

func main() {
	input := utils.ReadFile("resources/input.txt")
	motions := ParseInput(input)
	rope := CreateRope()
	for _, motion := range motions {
		rope.MoveHead(motion)
	}

	uniqueVisited := GetUnique(rope.GetVisitedSpaces())
	fmt.Println(len(uniqueVisited))
}
