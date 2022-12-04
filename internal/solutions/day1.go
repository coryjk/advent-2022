package solutions

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/coryjk/advent/internal/util"
	"github.com/coryjk/advent/internal/util/functional"
)

func Day1() {
	p := Problem{
		InputPath: "input/day1.txt",
	}
	var input []string = Input(p)
	fmt.Println(day1Part1(&input))
	fmt.Println(day1Part2(&input))
}

func day1Part1(input *[]string) string {
	calories := calories(input)

	// basic predicate
	isGreaterThan := func(left int, right int) bool {
		return left > right
	}

	// return max calorie count as string
	return strconv.Itoa(functional.Max(calories, isGreaterThan))
}

func day1Part2(input *[]string) string {
	calories := calories(input)

	// sort in ascending order
	sort.Ints(calories)

	// sum of last 3 elements (3 highest values)
	top3 := calories[len(calories)-3:]
	sum := util.SumInts(&top3)

	return strconv.Itoa(sum)
}

func calories(input *[]string) []int {
	var calories []int

	// determine list of total calorie counts per elf
	sum := 0
	for _, element := range *input {
		// we move onto the next elf during line break
		if len(element) == 0 {
			calories = append(calories, sum)
			sum = 0
		} else {
			// convert to int, accumulate calorie count
			calorie, _ := strconv.Atoi(element)
			sum += calorie
		}
	}

	return calories
}
