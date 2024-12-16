package day16

import (
	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
	"github.comjkondarewicz/aoc2024/pkg/utils"
)

type parser struct {
	width         int
	height        int
	startPosition aocmath.Vertex
	endPosition   aocmath.Vertex
	obstacles     *utils.Set[aocmath.Vertex]
}

func (parser *parser) ReadLine(index int, line string) {
	chars := []rune(line)
	parser.width = len(chars)
	for x, char := range chars {
		switch char {
		case '#':
			parser.obstacles.Add(aocmath.NewVertex(x, index))
		case 'S':
			parser.startPosition = aocmath.NewVertex(x, index)
		case 'E':
			parser.endPosition = aocmath.NewVertex(x, index)
		}
	}
}

func (parser *parser) toDay16Part1() solutions.Day16Part01 {
	return solutions.Day16Part01{
		StartPosition: parser.startPosition,
		EndPosition: parser.endPosition,
		Map: solutions.CreateMap(parser.width, parser.height, parser.obstacles),
	}
}

func (parser *parser) toDay16Part2() solutions.Day16Part02 {
	return solutions.Day16Part02{
		StartPosition: parser.startPosition,
		EndPosition: parser.endPosition,
		Map: solutions.CreateMap(parser.width, parser.height, parser.obstacles),
	}
}

func Day16ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		height:    file.Lines,
		obstacles: utils.NewSet[aocmath.Vertex](),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay16Part1()
	part2 := parser.toDay16Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
