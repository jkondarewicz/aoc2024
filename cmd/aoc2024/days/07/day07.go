package day07

import (
	"regexp"
	"strconv"

	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	equations []solutions.Day07Equation
}

func (parser *parser) ReadLine(index int, line string) {
	numDetect := regexp.MustCompile(`\d+`)
	nums := numDetect.FindAllString(line, -1)
	equation := solutions.Day07Equation {
		Result: 0,
		Nums: make([]int, 0),
	}
	for index, numString := range nums {
		num, _ := strconv.Atoi(numString)
		if index == 0 {
			equation.Result = num
		} else {
			equation.Nums = append(equation.Nums, num)
		}
	}
	parser.equations = append(parser.equations, equation)
}

func (parser *parser) toDay7Part1() solutions.Day07Part01 {
	return solutions.Day07Part01{
		Equations: parser.equations,
	}
}

func (parser *parser) toDay7Part2() solutions.Day07Part02 {
	return solutions.Day07Part02{
		Equations: parser.equations,
	}
}

func Day7ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		equations: make([]solutions.Day07Equation, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay7Part1()
	part2 := parser.toDay7Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
