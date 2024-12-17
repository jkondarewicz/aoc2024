package day14

import (
	"regexp"
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	width int
	height int
	robots []solutions.Day14Robot
}

func (parser *parser) ReadLine(index int, line string) {
	numRegex := regexp.MustCompile(`-?\d+`)
	numStr := numRegex.FindAllString(line, -1)
	sx, sy, tx, ty := getNum(numStr[0]), getNum(numStr[1]), getNum(numStr[2]), getNum(numStr[3])
	tx, ty = tx % parser.width, ty % parser.height
	parser.robots = append(parser.robots, solutions.Day14Robot{
		StartPosition: aocmath.NewVertex(sx, sy),
		Transition: aocmath.NewVertex(tx, ty),
	})
}

func getNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func (parser *parser) toDay14Part1() solutions.Day14Part01 {
	return solutions.Day14Part01{
		Width: parser.width,
		Height: parser.height,
		Robots: parser.robots,
	}
}

func (parser *parser) toDay14Part2() solutions.Day14Part02 {
	return solutions.Day14Part02{
		Width: parser.width,
		Height: parser.height,
		Robots: parser.robots,
	}
}

func Day14ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	width, height := 101, 103
	if test {
		width, height = 11, 7
	}
	parser := parser{
		width: width,
		height: height,
		robots: make([]solutions.Day14Robot, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay14Part1()
	part2 := parser.toDay14Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
