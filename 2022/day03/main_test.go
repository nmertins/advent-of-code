package main

import (
	"testing"

	"github.com/nmertins/advent-of-code/2022/utils"
)

func TestDay03(t *testing.T) {
	t.Run("parse sample input", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		rucksacks := ParseInput(input)
		if len(rucksacks) != 6 {
			t.Fatalf("expected 6 rucksacks, got %d", len(rucksacks))
		}
		for i, rucksack := range rucksacks {
			if len(rucksack) != 2 {
				t.Fatalf("rucksack %d from sample_input.txt has %d compartments, expected 2", i, len(rucksack))
			}
		}
		if rucksacks[0][0] != "vJrwpWtwJgWr" {
			t.Fatalf("first compartment of rucksack 0 has the wrong items. expected 'vJrwpWtwJgWr', got %q", rucksacks[0][0])
		}
		if rucksacks[0][1] != "hcsFMMfFFhFp" {
			t.Fatalf("second compartment of rucksack 0 has the wrong items. expected 'vJrwpWtwJgWr', got %q", rucksacks[0][1])
		}
	})

	t.Run("find matching items", func(t *testing.T) {
		firstCompartment := "vJrwpWtwJgWr"
		secondCompartment := "hcsFMMfFFhFp"

		items := FindMatchingItems(firstCompartment, secondCompartment)

		if items != "p" {
			t.Fatalf("expected 'p', got %q", items)
		}
	})

	t.Run("get item priority", func(t *testing.T) {
		priority := GetPriority("a")
		if priority != 1 {
			t.Fatalf("expected 1, got %d", priority)
		}

		priority = GetPriority("z")
		if priority != 26 {
			t.Fatalf("expected 1, got %d", priority)
		}

		priority = GetPriority("A")
		if priority != 27 {
			t.Fatalf("expected 1, got %d", priority)
		}

		priority = GetPriority("Z")
		if priority != 52 {
			t.Fatalf("expected 1, got %d", priority)
		}
	})

	t.Run("test sample input for part 1", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		rucksacks := ParseInput(input)
		total := 0
		for _, rucksack := range rucksacks {
			item := FindMatchingItems(rucksack[0], rucksack[1])
			total += GetPriority(item)
		}

		if total != 157 {
			t.Fatalf("expected total priority 157, got %d", total)
		}
	})

	t.Run("find group badge", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		badges := GetGroupBadges(input)
		if badges[0] != "r" {
			t.Fatalf("expected badge to be 'r', got %q", badges[0])
		}
		if badges[1] != "Z" {
			t.Fatalf("expected badge to be 'Z', got %q", badges[0])
		}
	})

	t.Run("test sample input for part 2", func(t *testing.T) {
		input := utils.ReadFile("resources/sample_input.txt")
		badges := GetGroupBadges(input)
		total := 0
		for _, badge := range badges {
			total += GetPriority(badge)
		}

		if total != 70 {
			t.Fatalf("expected 70, got %d", total)
		}
	})
}
