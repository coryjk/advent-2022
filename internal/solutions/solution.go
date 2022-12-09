package solutions

import "github.com/coryjk/advent/internal/util"

type Problem struct {
	InputPath string
	Input     *[]string
}

func Input(p Problem) []string {
	input, _ := util.ReadLines(p.InputPath)
	return input
}

func New(path string) *Problem {
	p := Problem{InputPath: path}
	p.Init()
	return &p
}

func (p *Problem) Init() {
	input, _ := util.ReadLines(p.InputPath)
	p.Input = &input
}
