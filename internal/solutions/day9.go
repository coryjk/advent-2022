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
	oldPos  position
	visited map[position]bool
}

func SolveDay9() {
	d := Day9{New("input/day9.txt")}
	fmt.Println(d.Part1())
	fmt.Println(d.Part2())
}

func (d *Day9) Part1() string {
	// tracking not required for head
	head := knot{pos(0, 0), pos(0, 0), nil}
	tail := knot{
		pos(0, 0),
		pos(0, 0),
		make(map[position]bool),
	}
	rope := []*knot{&head, &tail}

	// tail has already visited start
	tail.visited[pos(0, 0)] = true

	pullRope(&rope, d.problem.Input)
	return strconv.Itoa(len(tail.visited))
}

func (d *Day9) Part2() string {
	rope := make([]*knot, 10)
	for i := range rope {
		rope[i] = &knot{pos(0, 0), pos(0, 0), nil}
	}

	// only the tail is tracking
	tail := rope[len(rope)-1]
	tail.visited = make(map[position]bool)

	// tail has already visited start
	tail.visited[pos(0, 0)] = true

	pullRope(&rope, d.problem.Input)
	return strconv.Itoa(len(tail.visited))
}

func pullRope(rope *[]*knot, instructions *[]string) {
	head, _rope := (*rope)[0], *rope
	for _, line := range *instructions {
		split := strings.Split(line, " ")
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {
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

			// pull all knots
			for k := 1; k < len(*rope); k++ {
				pullTailIfNeeded(_rope[k-1], _rope[k])
			}
		}
	}
}

// assumption: assume head can only move 1 step at a time
func pullTailIfNeeded(head *knot, tail *knot) {
	headPos := head.pos
	tailPos := tail.pos

	// pull needed
	if abs(headPos.x-tailPos.x) > 1 || abs(headPos.y-tailPos.y) > 1 {
		// set tail to where head once was
		tail.oldPos = tailPos

		// move towards original position
		dx := headPos.x - tailPos.x
		dy := headPos.y - tailPos.y

		// bound by [-1, 1]
		dx, dy = min(max(dx, -1), 1), min(max(dy, -1), 1)
		newPos := position{
			x: tailPos.x + dx,
			y: tailPos.y + dy,
		}
		tail.pos = newPos

		// now mark as visited
		if tail.visited != nil {
			tail.visited[tail.pos] = true
		}
	}
}

func pos(x int, y int) position {
	return position{x, y}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func posString(pos position) string {
	return fmt.Sprintf("(x=%d,y=%d)", pos.x, pos.y)
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
