package util

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SumInts(nums *[]int) int {
	sum := 0
	for _, n := range *nums {
		sum += n
	}
	return sum
}

// Golang performs truncated modulos by default. This implementation
// handles negative modulos similarly to Python.
// E.g., Golang: -1 % 3 = -1, Python: -1 % 3 = 2
func EuclideanMod(a, b int) int {
	return (a%b + b) % b
}

func Intersection[V comparable](set1 []V, set2 []V) []V {
	present := make(map[V]bool)
	intersection := make([]V, 0)

	for _, element := range set1 {
		present[element] = true
	}

	for _, element := range set2 {
		if present[element] {
			intersection = append(intersection, element)
		}
	}

	return intersection
}
