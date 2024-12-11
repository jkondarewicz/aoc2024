package day11

import (
	"regexp"
	"strconv"

	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	nums []int
}

func (parser *parser) ReadLine(index int, line string) {
	digitsRegex := regexp.MustCompile(`\d+`)
	digitsStr := digitsRegex.FindAllString(line, -1)
	for _, digitStr := range digitsStr {
		digit, _ := strconv.Atoi(digitStr)
		parser.nums = append(parser.nums, digit)
	}

}

func (parser *parser) toDay11Part1() solutions.Day11Part01 {
	return solutions.Day11Part01{
		Nums: parser.nums,
	}
}

func (parser *parser) toDay11Part2() solutions.Day11Part02 {
	return solutions.Day11Part02{
		Nums: parser.nums,
	}
}

func Day11ResolverProvide(filename string) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		nums: make([]int, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay11Part1()
	part2 := parser.toDay11Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
