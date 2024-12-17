package day17

import (
	"regexp"
	"strconv"

	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	a    int
	b    int
	c    int
	nums []int
}

func (parser *parser) ReadLine(index int, line string) {
	numr := regexp.MustCompile(`\d+`)
	nums := numr.FindAllString(line, -1)
	for _, number := range nums {
		num, _ := strconv.Atoi(number)
		switch index {
		case 0:
			parser.a = num
		case 1:
			parser.b = num
		case 2:
			parser.c = num
		case 4:
			parser.nums = append(parser.nums, num)
		}
	}
}

func (parser *parser) toDay17Part1() solutions.Day17Part01 {
	return solutions.Day17Part01{
		Program: solutions.Day17Program{
			RegisterA: parser.a,
			RegisterB: parser.b,
			RegisterC: parser.c,
			Program: parser.nums,
		},
	}
}

func (parser *parser) toDay17Part2() solutions.Day17Part02 {
	return solutions.Day17Part02{
		Program: solutions.Day17Program{
			RegisterA: parser.a,
			RegisterB: parser.b,
			RegisterC: parser.c,
			Program: parser.nums,
		},
	}
}

func Day17ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		nums: make([]int, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay17Part1()
	part2 := parser.toDay17Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
