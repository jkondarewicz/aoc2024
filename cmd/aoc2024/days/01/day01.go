package day01

import (
	solution "github.comjkondarewicz/aoc2024/cmd/aoc2024/days"
	"github.comjkondarewicz/aoc2024/internal/benchmark"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/internal/solutions"
	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
	"github.comjkondarewicz/aoc2024/pkg/utils"
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
	Nums []int
}
func (parser parser) ReadLine(index int, line string) {
	first := -1; last := -1 
	chars := []rune(line)
	for i := 0; i < len(chars); i++ {
		digit, ok := utils.CharToInt(chars[i])
		if !ok {
			continue
		}
		if first == -1 {
			first = digit; last = digit
		} else {
			last = digit
		}
	}

	parser.Nums[index] = int(aocmath.Max(first * 10 + last, 0)) 
}


func Day01(filename string) solution.Solution {
	file, err := files.Open(filename) 
	if err != nil {
		return solution.ErrorSolution{
			Error: err,
		}
	}
	parser := parser { Nums: make([]int, file.Lines) }
	file.ProcessLineByLine(parser)
	part1 := solutions.Day01Part01 { Nums: parser.Nums }
	part2 := solutions.Day01Part02 {}
	return daySolution{
		part1: &part1,
		part2: &part2,
	}
}

