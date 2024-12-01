package day01

import (
	"strconv"
	"strings"

	solution "github.comjkondarewicz/aoc2024/cmd/aoc2024/days"
	"github.comjkondarewicz/aoc2024/internal/benchmark"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/internal/solutions"
)

type daySolution struct {
	part1 benchmark.BenchmarkExec[string]
	part2 benchmark.BenchmarkExec[string]
}

func (s daySolution) Part1Solution() (benchmark.BenchmarkResult[string], error) {
	return benchmark.Benchmark(s.part1)
}

func (s daySolution) Part2Solution() (benchmark.BenchmarkResult[string], error) {
	return benchmark.Benchmark(s.part2)
}

type parser struct {
	lefts []int
	rights []int
	rightsPart2 map[int]int
}

func (parser *parser) ReadLine(index int, line string) {
	nums := strings.Split(line, "   ")
	left, error := strconv.Atoi(nums[0])
	if error != nil {
		panic("cannot parse input text")
	}
	right, error := strconv.Atoi(nums[1])
	if error != nil {
		panic("cannot parse input text")
	}
	parser.lefts[index] = left
	parser.rights[index] = right
	parser.rightsPart2[right] = parser.rightsPart2[right] + 1
}


func Day01(filename string) solution.Solution {
	file, err := files.Open(filename) 
	if err != nil {
		return solution.ErrorSolution{
			Error: err,
		}
	}
	parser := parser { 
		lefts: make([]int, file.Lines),
		rights: make([]int, file.Lines),
		rightsPart2: make(map[int]int),
	}
	file.ProcessLineByLine(&parser)
	part1 := solutions.Day01Part01 { Lefts: parser.lefts, Rights: parser.rights }
	part2 := solutions.Day01Part02 { Lefts: parser.lefts, Rights: parser.rightsPart2}
	return daySolution{
		part1: &part1,
		part2: &part2,
	}
}

