package day12

import (
	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	width  int
	height int
	plants map[aocmath.Vertex]rune
}

func (parser *parser) ReadLine(index int, line string) {
	parser.width = len(line)
	chars := []rune(line)
	for x, plant := range chars {
		position := aocmath.Vertex{X: x * 2, Y: index * 2}
		parser.plants[position] = plant
	}
}

func (parser *parser) toDay12Part1() solutions.Day12Part01 {
	return solutions.Day12Part01{
		Width: parser.width,
		Height: parser.height,
		Plants: parser.plants,
	}
}

func (parser *parser) toDay12Part2() solutions.Day12Part02 {
	return solutions.Day12Part02{
		Width: parser.width,
		Height: parser.height,
		Plants: parser.plants,
	}
}

func Day12ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		height: file.Lines,
		plants: make(map[aocmath.Vertex]rune),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay12Part1()
	part2 := parser.toDay12Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
