package day15

import (
	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type parser struct {
	width      int
	height     int
	warehouse  solutions.Day15Map
	warehouse2 solutions.Day15Warehouse
	robot      solutions.Day15Robot
}

func (parser *parser) ReadLine(index int, line string) {
	chars := []rune(line)
	if parser.width == 0 {
		parser.width = len(chars) * 2
	}
	lineWithMap := false
	for x, char := range chars {
		switch char {
		case '#':
			lineWithMap = true
			parser.warehouse.Walls.Add(aocmath.NewVertex(x, index))
			parser.warehouse2.Walls.Add(aocmath.NewVertex(x*2, index))
			parser.warehouse2.Walls.Add(aocmath.NewVertex(x*2+1, index))
		case 'O':
			parser.warehouse.Boxes.Add(aocmath.NewVertex(x, index))
			parser.warehouse2.Boxes[aocmath.NewVertex(x*2, index)] = aocmath.NewVertex(x*2, index)
			parser.warehouse2.Boxes[aocmath.NewVertex(x*2+1, index)] = aocmath.NewVertex(x*2, index)
		case '@':
			parser.robot.Position = aocmath.NewVertex(x, index)
		default:
			direction, exists := solutions.CreateRobotDirection(char)
			if exists {
				parser.robot.Moves = append(parser.robot.Moves, direction)
			}
		}
	}
	if lineWithMap {
		parser.height++
	}
}

func (parser *parser) toDay15Part1() solutions.Day15Part01 {
	return solutions.Day15Part01{
		Map:   parser.warehouse,
		Robot: parser.robot,
	}
}

func (parser *parser) toDay15Part2() solutions.Day15Part02 {
	return solutions.Day15Part02{
		Height: parser.height,
		Width:  parser.width,
		Map:    parser.warehouse2,
		Robot:  parser.robot,
	}
}

func Day15ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		warehouse: solutions.Day15Map{
			Walls: utils.NewSet[aocmath.Vertex](),
			Boxes: utils.NewSet[aocmath.Vertex](),
		},
		warehouse2: solutions.Day15Warehouse{
			Walls: utils.NewSet[aocmath.Vertex](),
			Boxes: make(map[aocmath.Vertex]aocmath.Vertex),
		},
		robot: solutions.Day15Robot{
			Moves: make([]solutions.Day15Direction, 0),
		},
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay15Part1()
	part2 := parser.toDay15Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
