package main

import (
	"fmt"
	"sort"

	"github.com/nmertins/advent-of-code/2022/utils"
)

type TreePatch [][]int

func (t TreePatch) GetTreesNorthOf(row int, column int) []int {
	ret := make([]int, row)
	for i := 0; i < row; i++ {
		ret[i] = t[i][column]
	}
	return ret
}

func (t TreePatch) GetTreesEastOf(row int, column int) []int {
	ret := make([]int, len(t[row])-column-1)
	for i := 0; i < len(t[row])-column-1; i++ {
		ret[i] = t[row][column+i+1]
	}
	return ret
}

func (t TreePatch) GetTreesSouthOf(row int, column int) []int {
	ret := make([]int, len(t)-row-1)
	for i := 0; i < len(t)-row-1; i++ {
		ret[i] = t[row+i+1][column]
	}
	return ret
}

func (t TreePatch) GetTreesWestOf(row int, column int) []int {
	ret := make([]int, column)
	for i := 0; i < column; i++ {
		ret[i] = t[row][i]
	}
	return ret
}

func (t TreePatch) IsTreeVisible(row int, column int) bool {
	if t.onBoundary(row, column) {
		return true
	}

	size := t[row][column]
	north := t.GetTreesNorthOf(row, column)
	sort.Ints(north)
	visibleFromNorth := size > north[len(north)-1]
	east := t.GetTreesEastOf(row, column)
	sort.Ints(east)
	visibleFromEast := size > east[len(east)-1]
	south := t.GetTreesSouthOf(row, column)
	sort.Ints(south)
	visiableFromSouth := size > south[len(south)-1]
	west := t.GetTreesWestOf(row, column)
	sort.Ints(west)
	visibleFromWest := size > west[len(west)-1]

	return visibleFromNorth || visibleFromEast || visiableFromSouth || visibleFromWest
}

func (t TreePatch) onBoundary(row int, column int) bool {
	northBoundary := row == 0
	eastBoundary := column == len(t[row])-1
	southBoundary := row == len(t)-1
	westBoundary := column == 0
	return northBoundary || eastBoundary || southBoundary || westBoundary
}

func (t TreePatch) GetNumberOfVisibleTrees() int {
	total := 0

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			if t.IsTreeVisible(i, j) {
				total++
			}
		}
	}

	return total
}

func (t TreePatch) GetNumberOfVisibleTreesFromPoint(row int, column int) int {
	return 0
}

func ParseInput(input []string) TreePatch {
	trees := make([][]int, len(input))
	for i := range input {
		trees[i] = make([]int, len(input[i]))
		for j, r := range input[i] {
			trees[i][j] = int(r - '0')
		}
	}

	return trees
}

func main() {
	input := utils.ReadFile("resources/input.txt")
	trees := ParseInput(input)
	total := trees.GetNumberOfVisibleTrees()
	fmt.Println(total)
}
