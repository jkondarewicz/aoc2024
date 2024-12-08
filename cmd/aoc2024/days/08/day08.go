package day08

import (
	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/files"
	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
	"github.comjkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	width    int
	height   int
	antennas map[rune][]aocmath.Vertex
}

func (parser *parser) ReadLine(index int, line string) {
	parser.width = len(line)
	for x, char := range []rune(line) {
		if char != '.' {
			if _, exists := parser.antennas[char]; !exists {
				parser.antennas[char] = make([]aocmath.Vertex, 0)
			}
			parser.antennas[char] = append(parser.antennas[char], aocmath.Vertex{X: x, Y: index})
		}
	}
}

func (parser *parser) toDay8Part1() solutions.Day08Part01 {
	return solutions.Day08Part01{
		Width: parser.width,
		Height: parser.height,
		Antennas: parser.antennas,
	}
}

func (parser *parser) toDay8Part2() solutions.Day08Part02 {
	return solutions.Day08Part02{
		Width: parser.width,
		Height: parser.height,
		Antennas: parser.antennas,
	}
}

func Day8ResolverProvide(filename string) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		height: file.Lines,
		antennas: make(map[rune][]aocmath.Vertex),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay8Part1()
	part2 := parser.toDay8Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
