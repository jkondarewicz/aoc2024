package day02

import (
	"strconv"
	"strings"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	nums [][]int
}

func (parser *parser) ReadLine(index int, line string) {
	numsString := strings.Split(line, " ")
	nums := make([]int, len(numsString))
	for i := 0; i < len(numsString); i++ {
		num, err := strconv.Atoi(numsString[i])
		if err != nil {
			panic("Cannot parse input")
		}
		nums[i] = num
	}
	parser.nums[index] = nums
}

func (parser *parser) toDay1Part1() solutions.Day02Part01 {
	reports := make([]solutions.Report, len(parser.nums))
	for i := 0; i < len(parser.nums); i++ {
		reports[i].Nums = parser.nums[i]
	}
	return solutions.Day02Part01{
		Reports: reports,
	}
}


func Day2ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename) 
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser { 
		nums: make([][]int, file.Lines),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay1Part1()
	part2 := solutions.Day02Part02 { Reports: part1.Reports }
	return solutionTypes.DayResolver {
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
