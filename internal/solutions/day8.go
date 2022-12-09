package solutions

import (
	"fmt"
	"strconv"
)

type Day8 struct {
	problem *Problem
}

func SolveDay8() {
	d := Day8{New("input/day8.txt")}
	fmt.Println(d.Part1())
	fmt.Println(d.Part2())
}

func (d *Day8) Part1() string {
	grid := createGrid(d.problem.Input)
	n := len(*grid)

	// treat all edge trees as visible
	visible := 4*n - 4

	// check all trees within the inner edge
	for r := 1; r < n-1; r++ {
		for c := 1; c < n-1; c++ {
			up, down, left, right := rays(r, c, grid)
			if up == 0 || down == n-1 || left == 0 || right == n-1 {
				visible++
			}
		}
	}

	return strconv.Itoa(visible)
}

func (d *Day8) Part2() string {
	grid := createGrid(d.problem.Input)
	n := len(*grid)

	maxScore := -1
	for r := 1; r < n-1; r++ {
		for c := 1; c < n-1; c++ {
			up, down, left, right := rays(r, c, grid)

			// blocking trees are still considered "visible"
			if up > 0 {
				up--
			}
			if down < n-1 {
				down++
			}
			if left > 0 {
				left--
			}
			if right < n-1 {
				right++
			}

			score := (r - up) * (down - r) * (c - left) * (right - c)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return strconv.Itoa(maxScore)
}

func rays(row int, col int, grid *[][]int) (int, int, int, int) {
	_grid := *grid
	height, n := _grid[row][col], len(_grid)

	// rays to trace from source, visible if rays reach edge
	// up:    x--, x >= 0
	// down:  x++, x < n
	// left:  y--, y >= 0
	// right: y++, y < n
	up, down, left, right := row, row, col, col

	// trace rays from [row, col] externally
	i := 0
	for i < n {
		// yet to reach top
		if up > 0 && height > _grid[up-1][col] {
			// move ray up
			up--
		}

		// yet to reach bottom
		if down < n-1 && height > _grid[down+1][col] {
			// move ray down
			down++
		}

		// yet to reach left
		if left > 0 && height > _grid[row][left-1] {
			// move ray left
			left--
		}

		// yet to reach right
		if right < n-1 && height > _grid[row][right+1] {
			// move ray right
			right++
		}

		i++
	}
	return up, down, left, right
}

func createGrid(input *[]string) *[][]int {
	// init empty grid
	n := len((*input)[0])
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// populate grid
	for row, line := range *input {
		for col, ch := range line {
			// row-major order
			grid[row][col] = int(ch - '0')
		}
	}

	return &grid
}
