package solutions

import "github.com/coryjk/advent/internal/util"

type Problem struct {
	InputPath string
}

func Input(p Problem) []string {
	input, _ := util.ReadLines(p.InputPath)
	return input
}
