package day10

import (
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	startPositions []aocmath.Vertex
	mapTrail       map[aocmath.Vertex]int
	mapHeight      int
	mapWidth       int
}

func (parser *parser) ReadLine(index int, line string) {
	chars := []rune(line)
	parser.mapWidth = len(chars)
	for x, char := range chars {
		if char == '.' {
			continue
		}
		num, _ := strconv.Atoi(string(char))
		curPos := aocmath.Vertex{X: x, Y: index}
		parser.mapTrail[curPos] = num
		if num == 0 {
			parser.startPositions = append(parser.startPositions, curPos)
		}
	}
}

func (parser *parser) toDay10Part1() solutions.Day10Part01 {
	return solutions.Day10Part01{
		StartPositions: parser.startPositions,
		MapTrail:       parser.mapTrail,
		MapHeight:      parser.mapHeight,
		MapWidth:       parser.mapWidth,
	}
}

func (parser *parser) toDay10Part2() solutions.Day10Part02 {
	return solutions.Day10Part02{
		StartPositions: parser.startPositions,
		MapTrail:       parser.mapTrail,
		MapHeight:      parser.mapHeight,
		MapWidth:       parser.mapWidth,
	}
}

func Day10ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		startPositions: make([]aocmath.Vertex, 0),
		mapTrail:       make(map[aocmath.Vertex]int),
		mapHeight:      file.Lines,
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay10Part1()
	part2 := parser.toDay10Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
