package main

import (
	"reflect"
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay08Part1(t *testing.T) {
	t.Run("parse input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		trees := ParseInput(input)
		assertTreeHeight(t, 3, trees[0][0])
		assertTreeHeight(t, 6, trees[2][0])
		assertTreeHeight(t, 5, trees[3][2])
	})

	t.Run("north tree line", func(t *testing.T) {
		trees := createTreePatch()
		north := trees.GetTreesNorthOf(2, 2)
		expectedNorth := []int{3, 6}
		assertTreeLine(t, expectedNorth, north)

		north = trees.GetTreesNorthOf(0, 0)
		assertTreeLine(t, []int{}, north)
	})

	t.Run("east tree line", func(t *testing.T) {
		trees := createTreePatch()
		east := trees.GetTreesEastOf(1, 0)
		expectedEast := []int{5, 6}
		assertTreeLine(t, expectedEast, east)
	})

	t.Run("south tree line", func(t *testing.T) {
		trees := createTreePatch()
		south := trees.GetTreesSouthOf(0, 2)
		expectedSouth := []int{6, 9}
		assertTreeLine(t, expectedSouth, south)
	})

	t.Run("west tree line", func(t *testing.T) {
		trees := createTreePatch()
		west := trees.GetTreesWestOf(0, 2)
		expectedWest := []int{1, 2}
		assertTreeLine(t, expectedWest, west)
	})

	t.Run("is tree visible", func(t *testing.T) {
		trees := TreePatch{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		isVisible := trees.IsTreeVisible(1, 1)
		if isVisible {
			t.Fatalf("expected tree at %d, %d to not be visible", 1, 1)
		}

		isVisible = trees.IsTreeVisible(0, 0)
		if !isVisible {
			t.Fatalf("expected tree at %d, %d to be visible", 0, 0)
		}
	})

	t.Run("test sample input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		trees := ParseInput(input)
		visible := trees.GetNumberOfVisibleTrees()
		assertVisibleTrees(t, 21, visible)
	})
}

func TestDay08Part2(t *testing.T) {
	t.Run("get number of visible trees", func(t *testing.T) {
		trees := createTreePatch()
		visibleTrees, _ := trees.GetNumberOfVisibleTreesFromPoint(1, 1)
		assertVisibleTrees(t, 4, visibleTrees)

		trees = TreePatch{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		visibleTrees, _ = trees.GetNumberOfVisibleTreesFromPoint(2, 2)
		assertVisibleTrees(t, 8, visibleTrees)
	})

	t.Run("test sample input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		trees := ParseInput(input)
		visible, _ := trees.GetNumberOfVisibleTreesFromPoint(1, 2)
		assertVisibleTrees(t, 6, visible)

		visible, _ = trees.GetNumberOfVisibleTreesFromPoint(3, 2)
		assertVisibleTrees(t, 7, visible)
	})
}

func createTreePatch() TreePatch {
	return TreePatch{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
}

func assertTreeHeight(t testing.TB, expected int, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("expected tree with height %d, got %d", expected, actual)
	}
}

func assertTreeLine(t testing.TB, expected []int, actual []int) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, actual %v", expected, actual)
	}
}

func assertVisibleTrees(t testing.TB, expected int, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("expected %d trees to be visible, got %d", expected, actual)
	}
}
