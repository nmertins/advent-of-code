package main

import (
	"strconv"
	"strings"
)

type Stack []Crate

type Crate string

func (c Crate) IsEmpty() bool {
	return len(c) == 0
}

type Procedure []string

type Step struct {
	numberOfCratesToMove int
}

func ParseInput(input []string) ([]Stack, Procedure) {
	var dividerIndex int
	for i := range input {
		if len(input[i]) == 0 {
			dividerIndex = i
		}
	}

	stacks := parseStacks(input[:dividerIndex])
	procedure := parseProcedure(input[dividerIndex+1:])

	return stacks, procedure
}

func parseStacks(stackString []string) []Stack {
	numberOfStacks := parseNumberOfStacks(stackString[len(stackString)-1])
	stacks := make([]Stack, numberOfStacks)
	for i := 0; i < len(stackString)-1; i++ {
		crates := parseCrates(stackString[i])
		for j := 0; j < numberOfStacks; j++ {
			if !crates[j].IsEmpty() {
				stacks[j] = append(stacks[j], crates[j])
			}
		}
	}

	return stacks
}

func parseNumberOfStacks(stackNumberString string) int {
	trimmed := strings.TrimSpace(stackNumberString)
	indexes := strings.Split(trimmed, "   ")
	intVar, err := strconv.Atoi(indexes[len(indexes)-1])
	if err != nil {
		return 0
	}

	return intVar
}

func parseCrates(input string) []Crate {
	crates := make([]Crate, 0)
	for i := 0; i < len(input); i += 4 {
		crateString := input[i : i+3]
		crateString = strings.Trim(crateString, "[] ")
		var crate Crate = Crate(crateString)
		crates = append(crates, crate)
	}

	return crates
}

func parseProcedure(procedureString []string) Procedure {
	// fmt.Printf("procedure: %v", procedureString)
	return Procedure{}
}
