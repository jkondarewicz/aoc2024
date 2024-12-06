package day03

import (
	"regexp"
	"strconv"

	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	mulCommands []solutions.MulSystemCommand
}

func (parser *parser) ReadLine(index int, line string) {
	commandsRegexp := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := commandsRegexp.FindAllString(line, -1)
	commands := make([]solutions.MulSystemCommand, 0)
	digitsRegexp := regexp.MustCompile(`\d+`)
	for _, match := range matches {
		if match == "do()" {
			commands = append(commands, solutions.ActivateMul{})
			continue
		}
		if match == "don't()" {
			commands = append(commands, solutions.DeactivateMul{})
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
		commands = append(commands, solutions.PerformCalcMul{ X: x, Y: y })
	}
	parser.mulCommands = append(parser.mulCommands, commands...)
}

func (parser *parser) toDay3Part1() solutions.Day03Part01 {
	return solutions.Day03Part01{
		MulCommands: parser.mulCommands,
	}
}

func (parser *parser) toDay3Part2() solutions.Day03Part02 {
	return solutions.Day03Part02{
		MulCommands: parser.mulCommands,
	}
}

func Day3ResolverProvide(filename string) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename) 
	if err != nil {
		return solutionTypes.DayResolver {}, err
	}
	parser := parser { 
		mulCommands: make([]solutions.MulSystemCommand, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay3Part1()
	part2 := parser.toDay3Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
