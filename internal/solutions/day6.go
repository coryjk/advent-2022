package solutions

import (
	"container/list"
	"fmt"
	"strconv"
)

func Day6() {
	p := Problem{
		InputPath: "input/day6.txt",
	}
	var input []string = Input(p)
	fmt.Println(day6Part1(&input))
	fmt.Println(day6Part2(&input))
}

func day6Part1(input *[]string) string {
	datastream := (*input)[0]
	return strconv.Itoa(startOfMessageMarker(datastream, 4))
}

func day6Part2(input *[]string) string {
	datastream := (*input)[0]
	return strconv.Itoa(startOfMessageMarker(datastream, 14))
}

func contains(l *list.List, x any) bool {
	for i := l.Front(); i != nil; i = i.Next() {
		if i.Value == x {
			return true
		}
	}
	return false
}

func startOfMessageMarker(datastream string, unique int) int {
	buffer := list.New()
	for i, ch := range datastream {
		// pop from front until duplicate gone
		for contains(buffer, ch) {
			buffer.Remove(buffer.Front())
		}
		// append char
		buffer.PushBack(ch)
		if buffer.Len() >= unique {
			return i + 1
		}
	}
	return -1
}
