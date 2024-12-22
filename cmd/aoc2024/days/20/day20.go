package day20

import (
	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type parser struct {
	obstacles  *utils.Set[aocmath.Vertex]
	start, end aocmath.Vertex
}

func (parser *parser) ReadLine(index int, line string) {
	for x, char := range []rune(line) {
		pos := aocmath.NewVertex(x, index)
		switch char {
		case '#':
			parser.obstacles.Add(pos)
		case 'S':
			parser.start = pos
		case 'E':
			parser.end = pos
		}
	}
}

func Day20ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		obstacles: utils.NewSet[aocmath.Vertex](),
	}
	file.ProcessLineByLine(&parser)
	part1, part2 := solutions.CreateDay20(parser.obstacles, parser.start, parser.end)
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
