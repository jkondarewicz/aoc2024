package day03

import (
	"regexp"
	"strconv"

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
	mulCommands []solutions.MulCommand
	mulCommandsEnabled []solutions.MulCommand
	enabled bool
}

func (parser *parser) ReadLine(index int, line string) {
	commandsRegexp := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := commandsRegexp.FindAllString(line, -1)
	commands := make([]solutions.MulCommand, len(matches))
	commandsEnabled := make([]solutions.MulCommand, 0)
	digitsRegexp := regexp.MustCompile(`\d+`)
	for index, match := range matches {
		if match == "do()" {
			parser.enabled = true
			continue
		}
		if match == "don't()" {
			parser.enabled = false
			continue
		}

		digitMatch := digitsRegexp.FindAllString(match, -1)
		x, error := strconv.ParseInt(digitMatch[0], 10, 64)
		if error != nil {
			panic(error)
		}
		y, error := strconv.ParseInt(digitMatch[1], 10, 64)
		if error != nil {
			panic(error)
		}
		commands[index] = solutions.MulCommand { X: x, Y: y }
		if parser.enabled {
			commandsEnabled = append(commandsEnabled, solutions.MulCommand { X: x, Y: y })
		}
	}
	parser.mulCommands = append(parser.mulCommands, commands...)
	parser.mulCommandsEnabled = append(parser.mulCommandsEnabled, commandsEnabled...)
}

func (parser *parser) toDay3Part1() solutions.Day03Part01 {
	return solutions.Day03Part01{
		MulCommands: parser.mulCommands,
	}
}

func (parser *parser) toDay3Part2() solutions.Day03Part02 {
	return solutions.Day03Part02{
		MulCommands: parser.mulCommandsEnabled,
	}
}


func Day03(filename string) solution.Solution {
	file, err := files.Open(filename) 
	if err != nil {
		return solution.ErrorSolution{
			Error: err,
		}
	}
	parser := parser { 
		enabled: true,
		mulCommands: make([]solutions.MulCommand, 0),
		mulCommandsEnabled: make([]solutions.MulCommand, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay3Part1()
	part2 := parser.toDay3Part2()
	return daySolution{
		part1: &part1,
		part2: &part2,
	}
}

