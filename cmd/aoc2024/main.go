package main

import (
	"fmt"

	solution "github.comjkondarewicz/aoc2024/cmd/aoc2024/days"
	day01 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/01"
	"github.comjkondarewicz/aoc2024/internal/benchmark"
)


func main() {
	test := false

	var filename string
	if test {
		filename = "test"
	} else {
		filename = "real"
	}
	days := [] solution.Solution {
		day01.Day01("cmd/aoc2024/days/01/" + filename),
	}
	for i := 0; i < len(days); i++ {
		res, err := days[i].Part1Solution()
		printSolution(i + 1, 1, err, res)
		res, err = days[i].Part2Solution()
		printSolution(i + 1, 2, err, res)
	}
}

func printSolution(day int, part int, err error, result benchmark.BenchmarkResult[string]) {
	if err != nil {
		fmt.Printf("Day %d, part %d, error [%s]\n", day, part, err)
	} else {
		fmt.Printf("Day %d, part %d, solution took %dms\n", day, part, result.Time)
		fmt.Printf("%s\n", result.Result)
		fmt.Println()
	}
}
