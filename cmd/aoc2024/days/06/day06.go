package day06

import (
	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	startPosition aocmath.Vertex
	obstacles     map[aocmath.Vertex]bool
	xLength       int
}

func (parser *parser) ReadLine(index int, line string) {
	parser.xLength = len(line)
	for x, char := range []rune(line) {
		switch char {
		case '#':
			parser.obstacles[aocmath.Vertex{X: x, Y: index}] = true
		case '^':
			parser.startPosition = aocmath.Vertex{X: x, Y: index}
		}
	}
}

func (parser *parser) toDay6Part1(height int) solutions.Day06Part01 {
	return solutions.Day06Part01{
		StartPosition: parser.startPosition,
		Obstacles:     parser.obstacles,
		YLength:       height,
		XLength:       parser.xLength,
	}
}

func (parser *parser) toDay6Part2(height int) solutions.Day06Part02 {
	return solutions.Day06Part02{
		StartPosition: parser.startPosition,
		Obstacles:     parser.obstacles,
		YLength:       height,
		XLength:       parser.xLength,
	}
}

func Day6ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		obstacles: make(map[aocmath.Vertex]bool),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay6Part1(file.Lines)
	part2 := parser.toDay6Part2(file.Lines)
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
