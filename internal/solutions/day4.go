package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func Day4() {
	p := Problem{
		InputPath: "input/day4.txt",
	}
	var input []string = Input(p)
	fmt.Println(day4Part1(&input))
	fmt.Println(day4Part2(&input))
}

func day4Part1(input *[]string) string {
	fullyContainingPairs := 0

	for _, pair := range *input {
		// split pair of ranges provided by each line
		split := strings.Split(pair, ",")
		range1, range2 := split[0], split[1]

		// see if any of the ranges fully encompasses one another
		if fullyContains(range1, range2) || fullyContains(range2, range1) {
			fullyContainingPairs++
		}
	}

	return strconv.Itoa(fullyContainingPairs)
}

func day4Part2(input *[]string) string {
	overlappingPairs := 0

	for _, pair := range *input {
		// split pair of ranges provided by each line
		split := strings.Split(pair, ",")
		range1, range2 := split[0], split[1]

		left1, right1 := rangeAtoi(range1)
		left2, right2 := rangeAtoi(range2)

		// [left1...right1]
		// [left2...right2]
		if left1 <= right2 && left2 <= right1 {
			overlappingPairs++
		}
	}

	return strconv.Itoa(overlappingPairs)
}

// assumes range string follow format: "x-y",
// only checks if range1 fully contains range2
func fullyContains(range1 string, range2 string) bool {
	left1, right1 := rangeAtoi(range1)
	left2, right2 := rangeAtoi(range2)
	return left1 <= left2 && right1 >= right2
}

func rangeAtoi(range_ string) (int, int) {
	split := strings.Split(range_, "-")
	left, _ := strconv.Atoi(split[0])
	right, _ := strconv.Atoi(split[1])
	return left, right
}
