package day04

import (
	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

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

func Day4ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		chars: make([][]rune, file.Lines),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay4Part1()
	part2 := parser.toDay4Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
