package day02

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
	nums [][]int
}

func (parser *parser) ReadLine(index int, line string) {
	numsString := strings.Split(line, " ")
	nums := make([]int, len(numsString))
	for i := 0; i < len(numsString); i++ {
		num, err := strconv.Atoi(numsString[i])
		if err != nil {
			panic("Cannot parse input")
		}
		nums[i] = num
	}
	parser.nums[index] = nums
}

func (parser *parser) toDay1Part1() solutions.Day02Part01 {
	reports := make([]solutions.Report, len(parser.nums))
	for i := 0; i < len(parser.nums); i++ {
		reports[i].Nums = parser.nums[i]
	}
	return solutions.Day02Part01{
		Reports: reports,
	}
}


func Day02(filename string) solution.Solution {
	file, err := files.Open(filename) 
	if err != nil {
		return solution.ErrorSolution{
			Error: err,
		}
	}
	parser := parser { 
		nums: make([][]int, file.Lines),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay1Part1()
	part2 := solutions.Day02Part02 { Reports: part1.Reports }
	return daySolution{
		part1: &part1,
		part2: &part2,
	}
}

