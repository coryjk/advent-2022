package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day9 struct {
	problem *Problem
}

type position struct {
	x int
	y int
}

type knot struct {
	pos     position
	visited map[position]bool
}

func SolveDay9() {
	d := Day9{New("input/day9.txt")}
	fmt.Println(d.Part1())
	fmt.Println(d.Part2())
}

func (d *Day9) Part1() string {
	// tracking not required for head
	head := knot{position{0, 0}, nil}
	tail := knot{
		position{0, 0},
		make(map[position]bool),
	}

	// tail has already visited start
	tail.visited[pos(0, 0)] = true

	for _, line := range *d.problem.Input {
		split := strings.Split(line, " ")
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {
			oldPos := head.pos
			switch direction {
			case "R":
				// move right
				head.pos.x++
			case "L":
				// move left
				head.pos.x--
			case "U":
				// move up
				head.pos.y++
			case "D":
				// move down
				head.pos.y--
			}
			pullTailIfNeeded(&head, &tail, oldPos)
		}
	}

	return strconv.Itoa(len(tail.visited))
}

func (d *Day9) Part2() string {
	return ""
}

func pos(x int, y int) position {
	return position{x, y}
}

// assumption: assume head can only move 1 step at a time
func pullTailIfNeeded(head *knot, tail *knot, originalPos position) {
	headPos := head.pos
	tailPos := tail.pos

	// pull needed
	if abs(headPos.x-tailPos.x) > 1 || abs(headPos.y-tailPos.y) > 1 {
		// set tail to where head once was
		tail.pos = originalPos

		// now mark as visited
		tail.visited[originalPos] = true
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
