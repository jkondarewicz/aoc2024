package day18

import (
	"regexp"
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	bits []aocmath.Vertex
}

func (parser *parser) ReadLine(index int, line string) {
	numR := regexp.MustCompile(`\d+`)
	numsStr := numR.FindAllString(line, -1)
	x, _ := strconv.Atoi(numsStr[0])
	y, _ := strconv.Atoi(numsStr[1])
	parser.bits = append(parser.bits, aocmath.NewVertex(x, y))
}

func (parser *parser) toDay18Part1(width, height, bits int) solutions.Day18Part01 {
	return solutions.Day18Part01{
		Bits: parser.bits,
		Width: width,
		Height: height,
		HowMany: bits,
	}
}

func (parser *parser) toDay18Part2(width, height, bits int) solutions.Day18Part02 {
	return solutions.Day18Part02{
		Bits: parser.bits,
		Width: width,
		Height: height,
		HowMany: bits,
	}
}

func Day18ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	width, height, howMany := 7, 7, 12
	if !test {
		width, height, howMany = 71, 71, 1024
	}
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		bits: make([]aocmath.Vertex, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay18Part1(width, height, howMany)
	part2 := parser.toDay18Part2(width, height, howMany)
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
