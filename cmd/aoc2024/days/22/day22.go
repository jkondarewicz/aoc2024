package day22

import (
	"regexp"
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	secrets []int64
}

func (parser *parser) ReadLine(index int, line string) {
	numRegex := regexp.MustCompile(`\d+`)
	num, _ := strconv.ParseInt(numRegex.FindAllString(line, -1)[0], 10, 64)
	parser.secrets = append(parser.secrets, num)
}

func (parser *parser) toDay22Part1() solutions.Day22Part01 {
	return solutions.Day22Part01{
		Secrets: parser.secrets,
	}
}

func (parser *parser) toDay22Part2() solutions.Day22Part02 {
	return solutions.Day22Part02{
		Secrets: parser.secrets,
	}
}

func Day22ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		secrets: make([]int64, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay22Part1()
	part2 := parser.toDay22Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
