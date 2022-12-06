package main

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	StepRegex = `^move ([0-9]*) from ([0-9]*) to ([0-9]*)$`
)

type Stack []Crate

type Crate string

func (c Crate) IsEmpty() bool {
	return len(c) == 0
}

type Procedure []Step

type Step struct {
	cratesToMove int
	fromStack    int
	toStack      int
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
	procedure := Procedure{}
	for _, stepString := range procedureString {
		step := parseStep(stepString)

		procedure = append(procedure, step)
	}

	return procedure
}

func parseStep(stepString string) Step {
	re := regexp.MustCompile(StepRegex)
	matches := re.FindAllString(stepString, -1)
	cratesToMove, err := strconv.Atoi(matches[0])
	if err != nil {
		return Step{0, 0, 0}
	}
	fromStack, err := strconv.Atoi(matches[0])
	if err != nil {
		return Step{0, 0, 0}
	}
	toStack, err := strconv.Atoi(matches[0])
	if err != nil {
		return Step{0, 0, 0}
	}

	return Step{cratesToMove, fromStack, toStack}
}
