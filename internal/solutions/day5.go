package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/coryjk/advent/internal/util/collections"
)

func Day5() {
	p := Problem{
		InputPath: "input/day5.txt",
	}
	input := Input(p)
	fmt.Println(day5Part1(&input))
	fmt.Println(day5Part2(&input))
}

func day5Part1(input *[]string) string {
	crateStacks, instructions := parseInput(input)

	// perform move instructions on stacks
	for _, instruction := range *instructions {
		moveCount, fromStack, toStack := parseInstruction(instruction)

		// fetch stacks
		sourceStack := (*crateStacks)[fromStack]
		targetStack := (*crateStacks)[toStack]

		// transfer crates `moveCount` times
		for i := 0; i < int(moveCount); i++ {
			a := sourceStack.Pop()
			targetStack.Push(a)
		}
	}
	return peekAllCrateStacks(crateStacks)
}

func day5Part2(input *[]string) string {
	crateStacks, instructions := parseInput(input)
	printStacks(crateStacks)

	// perform move instructions on stacks
	for _, instruction := range *instructions {
		moveCount, fromStack, toStack := parseInstruction(instruction)

		// fetch stacks
		sourceStack := (*crateStacks)[fromStack]
		targetStack := (*crateStacks)[toStack]

		// transfer crates `moveCount` times, conserving order in batches
		var batch []interface{}
		for i := 0; i < int(moveCount); i++ {
			batch = append(batch, sourceStack.Pop())
		}

		// push backwards, conserve order from source stack
		for i := len(batch) - 1; i >= 0; i-- {
			targetStack.Push(batch[i].(rune))
		}
	}
	return peekAllCrateStacks(crateStacks)
}

func parseInput(input *[]string) (*[]*collections.Stack, *[]string) {
	// find where instructions begin
	instructionsStart := -1
	for i := 0; i < len(*input); i++ {
		line := (*input)[i]
		if strings.Contains(line, "move") {
			instructionsStart = i
			break
		}
	}

	// first populate all crate stacks
	crateStacks := []*collections.Stack{}
	indexes := (*input)[instructionsStart-2] // stack input ends 2 before instrunctionss

	// maps actual string pos to corresponding stack index
	indexToStackIndex := make(map[int]int)
	for i, ch := range indexes {
		if '1' <= ch && ch <= '9' {
			indexToStackIndex[i] = int(ch)

			// init new stack
			crateStacks = append(crateStacks, collections.NewStack())
		}
	}

	// populate stacks starting from "bottom", 1 before indexes pos
	for i := instructionsStart - 3; i >= 0; i-- {
		line := (*input)[i]

		// add crates, creates represented as alphabetical runes
		for j, ch := range line {
			if 'A' <= ch && ch <= 'Z' {
				// should have the same index as some existing stack,
				// needs to be converted from rune val to int -> offset
				// by rune('1')
				stackIndex := indexToStackIndex[j] - '1'
				stack := crateStacks[stackIndex]
				stack.Push(ch)
			}
		}
	}

	instructions := (*input)[instructionsStart:]
	return &crateStacks, &instructions
}

// Formatted as: "move {} from {} to {}"
func parseInstruction(instruction string) (int, int, int) {
	split := strings.Split(instruction, " ")

	moveCount, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])

	// zero-indexing for indexes that will be used directly
	return moveCount, from - 1, to - 1
}

// Just for debugging
func printStacks(stacks *[]*collections.Stack) {
	asString := func(x interface{}) string {
		return string(x.(rune))
	}
	for i, stack := range *stacks {
		fmt.Println(i, stack.String(asString))
	}
}

func peekAllCrateStacks(stacks *[]*collections.Stack) string {
	// peek all stacks for answer
	topCrates := strings.Builder{}
	for _, stack := range *stacks {
		topCrate := stack.Peek()

		// validate rune type then append as string
		if topCrate != nil {
			topCrates.WriteString(
				string(topCrate.(rune)))
		}
	}
	return topCrates.String()
}
