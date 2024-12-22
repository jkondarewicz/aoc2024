package day21

import (
	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	codes *solutions.Codes
}

func (parser *parser) ReadLine(index int, line string) {
	parser.codes.AddCode(line)
}

func (parser *parser) toDay21Part1() solutions.Day21Part01 {
	return solutions.Day21Part01{
		Codes: *parser.codes,
	}
}

func (parser *parser) toDay21Part2() solutions.Day21Part02 {
	return solutions.Day21Part02{
		Codes: *parser.codes,
	}
}

func Day21ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		codes: solutions.CreateCodes(),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay21Part1()
	part2 := parser.toDay21Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
