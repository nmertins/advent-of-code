package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nmertins/advent-of-code/2022/utils"
)

type Elf struct {
	startID int
	endID   int
}

func (e Elf) AssignmentLength() int {
	return e.endID - e.startID + 1
}

type CleaningPair struct {
	elves []Elf
}

func (c CleaningPair) IndexOfElfWithBiggestAssignment() int {
	if c.elves[0].AssignmentLength() >= c.elves[1].AssignmentLength() {
		return 0
	} else {
		return 1
	}
}

func (c CleaningPair) IndexOfElfWithLowestStartID() int {
	if c.elves[0].startID <= c.elves[1].startID {
		return 0
	} else {
		return 1
	}
}

func (c CleaningPair) AssignmentsCompletelyOverlap() bool {
	indexOfElfWithBiggestAssignment := c.IndexOfElfWithBiggestAssignment()
	indexOfElfWithSmallerAssignment := (indexOfElfWithBiggestAssignment + 1) % 2

	big := c.elves[indexOfElfWithBiggestAssignment]
	small := c.elves[indexOfElfWithSmallerAssignment]

	return big.startID <= small.startID && big.endID >= small.endID
}

func (c CleaningPair) AssignmentsOverlap() bool {
	indexOfElfWithLowestStartID := c.IndexOfElfWithLowestStartID()
	indexOfOtherElf := (indexOfElfWithLowestStartID + 1) % 2

	lowestStartID := c.elves[indexOfElfWithLowestStartID]
	other := c.elves[indexOfOtherElf]

	return lowestStartID.endID >= other.startID
}

func ParseInput(input []string) []CleaningPair {
	pairs := make([]CleaningPair, len(input))
	for i, line := range input {
		elvesString := strings.Split(line, ",")
		elves := make([]Elf, len(elvesString))
		for j, elf := range elvesString {
			ids := strings.Split(elf, "-")
			startID, err := strconv.Atoi(ids[0])
			if err != nil {
				fmt.Errorf("Could not parse start ID on line %d", i)
				return make([]CleaningPair, 0)
			}
			endID, err := strconv.Atoi(ids[1])
			if err != nil {
				fmt.Errorf("Could not parse end ID on line %d", i)
				return make([]CleaningPair, 0)
			}
			elves[j] = Elf{startID, endID}
		}

		pairs[i] = CleaningPair{elves}
	}

	return pairs
}

func main() {
	input := utils.ReadFile("resources/input.txt")
	pairs := ParseInput(input)
	total := 0
	for _, pair := range pairs {
		if pair.AssignmentsOverlap() {
			total++
		}
	}

	fmt.Println(total)
}
