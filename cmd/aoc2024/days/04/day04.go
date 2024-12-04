package day04

import (
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
	chars [][]rune
}

func (parser *parser) ReadLine(index int, line string) {
	chars := []rune(line)
	parser.chars[index] = chars
}

func (parser *parser) toDay4Part1() solutions.Day04Part01 {
	return solutions.Day04Part01{
		Chars: parser.chars,
	}
}

func (parser *parser) toDay4Part2() solutions.Day04Part02 {
	return solutions.Day04Part02{
		Chars: parser.chars,
	}
}


func Day04(filename string) solution.Solution {
	file, err := files.Open(filename) 
	if err != nil {
		return solution.ErrorSolution{
			Error: err,
		}
	}
	parser := parser { 
		chars: make([][]rune, file.Lines),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay4Part1()
	part2 := parser.toDay4Part2()
	return daySolution{
		part1: &part1,
		part2: &part2,
	}
}

