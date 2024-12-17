package day13

import (
	"regexp"
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	machines []solutions.Day13Machine
	machine  solutions.Day13Machine
}

func (parser *parser) ReadLine(index int, line string) {
	which := index % 4
	numberRegex := regexp.MustCompile(`\d+`)
	stringNums := numberRegex.FindAllString(line, -1)
	nums := make([]int, 0)
	for _, num := range stringNums {
		cNum, _ := strconv.Atoi(num)
		nums = append(nums, cNum)
	}
	switch which {
	case 0:
		parser.machine.ButtonA = aocmath.NewVertex(nums[0], nums[1])
	case 1:
		parser.machine.ButtonB = aocmath.NewVertex(nums[0], nums[1])
	case 2:
		parser.machine.WinningPosition = aocmath.NewVertex(nums[0], nums[1])
	case 3:
		parser.machines = append(parser.machines, parser.machine)
		parser.machine = solutions.Day13Machine{}
	}
}

func (parser *parser) toDay13Part1() solutions.Day13Part01 {
	return solutions.Day13Part01{
		Machines: parser.machines,
	}
}

func (parser *parser) toDay13Part2() solutions.Day13Part02 {
	return solutions.Day13Part02{
		Machines: parser.machines,
	}
}

func Day13ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		machines: make([]solutions.Day13Machine, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay13Part1()
	part2 := parser.toDay13Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
