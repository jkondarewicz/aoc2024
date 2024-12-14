package day01

import (
	"strconv"
	"strings"

	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)


type parser struct {
	lefts []int
	rights []int
	rightsPart2 map[int]int
}

func (parser *parser) ReadLine(index int, line string) {
	nums := strings.Split(line, "   ")
	left, error := strconv.Atoi(nums[0])
	if error != nil {
		panic("cannot parse input text")
	}
	right, error := strconv.Atoi(nums[1])
	if error != nil {
		panic("cannot parse input text")
	}
	parser.lefts[index] = left
	parser.rights[index] = right
	parser.rightsPart2[right] = parser.rightsPart2[right] + 1
}
func Day1ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename) 
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser { 
		lefts: make([]int, file.Lines),
		rights: make([]int, file.Lines),
		rightsPart2: make(map[int]int),
	}
	file.ProcessLineByLine(&parser)
	part1 := solutions.Day01Part01 { Lefts: parser.lefts, Rights: parser.rights }
	part2 := solutions.Day01Part02 { Lefts: parser.lefts, Rights: parser.rightsPart2}
	return solutionTypes.DayResolver {
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}

