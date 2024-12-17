package solutions

import (
	"sort"
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
)

type Day01Part01 struct {
	Lefts []int
	Rights []int
}

type Day01Part02 struct {
	Lefts []int
	Rights map[int]int
}

func (data *Day01Part01) Exec() (string, error) {
	sort.Ints(data.Lefts)
	sort.Ints(data.Rights)
	result := 0
	for i := 0; i < len(data.Lefts); i++ {
		diff := aocmath.Abs(data.Rights[i] - data.Lefts[i])
		result += diff
	}
	return strconv.Itoa(result), nil
}

func (data *Day01Part02) Exec() (string, error) {
	result := 0
	for _, num := range data.Lefts {
		result += data.Rights[num] * num
	}
	return strconv.Itoa(result), nil
}
