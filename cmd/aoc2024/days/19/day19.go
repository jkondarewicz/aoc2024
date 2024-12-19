package day19

import (
	"strings"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	towelPatterns []string
	designs       []string
}

func (parser *parser) ReadLine(index int, line string) {
	if index == 0 {
		str := strings.Split(line, ", ")
		parser.towelPatterns = append(parser.towelPatterns, str...)
	} else if len(line) > 0 {
		parser.designs = append(parser.designs, line)
	}
}

func (parser *parser) toDay19Part1() solutions.Day19Part01 {
	return solutions.Day19Part01{
		PossibleTowelPatterns: parser.towelPatterns,
		DesiredDesigns: parser.designs,
	}
}

func (parser *parser) toDay19Part2() solutions.Day19Part02 {
	return solutions.Day19Part02{
		PossibleTowelPatterns: parser.towelPatterns,
		DesiredDesigns: parser.designs,
	}
}

func Day19ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		towelPatterns: make([]string, 0),
		designs: make([]string, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay19Part1()
	part2 := parser.toDay19Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
