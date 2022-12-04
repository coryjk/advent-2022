package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/coryjk/advent/internal/util"
)

func Day2() {
	p := Problem{
		InputPath: "input/day2.txt",
	}
	var input []string = Input(p)
	fmt.Println(day2Part1(&input))
	fmt.Println(day2Part2(&input))
}

const scoreWin int = 6
const scoreDraw int = 3
const scoreLoss int = 0

const (
	rock     = iota
	paper    = iota
	scissors = iota
)

var moves = map[string]int{
	// their moves
	"A": rock,
	"B": paper,
	"C": scissors,
	// your moves
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var shapeScore = map[int]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

func day2Part1(input *[]string) string {
	totalScore := 0

	for _, round := range *input {
		parsed := strings.Split(round, " ")
		theirMove := moves[parsed[0]]
		yourMove := moves[parsed[1]]

		// check for draw
		if theirMove == yourMove {
			totalScore += scoreDraw
		} else {
			// no draw, check for win or loss
			result := won(theirMove, yourMove)

			// increment points for win, no points for losing
			if result {
				totalScore += scoreWin
			}
		}

		// always increment with shape score
		totalScore += shapeScore[yourMove]
	}

	return strconv.Itoa(totalScore)
}

func day2Part2(input *[]string) string {
	totalScore := 0
	result := map[string]int{
		"X": scoreLoss,
		"Y": scoreDraw,
		"Z": scoreWin,
	}

	for _, round := range *input {
		parsed := strings.Split(round, " ")
		theirMove := moves[parsed[0]]
		desiredOutcome := result[parsed[1]]

		// add to score based on desired outcome
		totalScore += desiredOutcome

		// evaluate necessary move based on outcome
		var yourMove int
		switch desiredOutcome {
		case scoreLoss:
			yourMove = util.EuclideanMod(theirMove-1, 3)
		case scoreDraw:
			yourMove = theirMove
		case scoreWin:
			yourMove = util.EuclideanMod(theirMove+1, 3)
		}

		// increment with your shape score
		totalScore += shapeScore[yourMove]
	}

	return strconv.Itoa(totalScore)
}

func won(theirMove int, yourMove int) bool {
	if theirMove == rock {
		return yourMove == paper
	} else if theirMove == paper {
		return yourMove == scissors
	} else {
		return yourMove == rock
	}
}
