package solutions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/coryjk/advent/internal/util"
)

func Day3() {
	p := Problem{
		InputPath: "input/day3.txt",
	}
	input := Input(p)
	fmt.Println(day3Part1(&input))
	fmt.Println(day3Part2(&input))
}

func day3Part1(input *[]string) string {
	var prioritySum int32 = 0
	for _, rucksack := range *input {
		firstHalf, secondHalf := splitCompartments(rucksack)
		for _, ch := range firstHalf {
			if strings.ContainsRune(secondHalf, ch) {
				priority, _ := priority(ch)
				prioritySum += priority
				break
			}
		}
	}
	return strconv.Itoa(int(prioritySum))
}

func day3Part2(input *[]string) string {
	var prioritySum int32 = 0
	lines := *input
	for i := 0; i < len(lines); i += 3 {
		line1, line2, line3 :=
			[]rune(lines[i]),
			[]rune(lines[i+1]),
			[]rune(lines[i+2])
		intersect := util.Intersection(line1, line2)
		intersect = util.Intersection(intersect, line3)

		priority, _ := priority(intersect[0])
		prioritySum += priority
	}
	return strconv.Itoa(int(prioritySum))
}

func splitCompartments(rucksack string) (string, string) {
	half := len(rucksack) / 2
	return rucksack[:half], rucksack[half:]
}

func priority(c rune) (int32, error) {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 1, nil
	} else if c >= 'A' && c <= 'Z' {
		return c - 'A' + 27, nil
	} else {
		return -1, errors.New("Invalid rune input: " + string(c))
	}
}
