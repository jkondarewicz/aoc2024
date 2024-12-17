package day09

import (
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	diskBlocks []solutions.Day09DiskBlocks
}

func (parser *parser) ReadLine(index int, line string) {
	currentDiskBlock := solutions.Day09DiskBlocks{}
	id := 0
	for index, char := range []rune(line) {
		num, _ := strconv.Atoi(string(char))
		switch index % 2 {
		case 0:
			currentDiskBlock.Id = id
			currentDiskBlock.FileBlocks = num
			id++
			continue
		case 1:
			currentDiskBlock.EmptyBlock = num
			parser.diskBlocks = append(parser.diskBlocks, currentDiskBlock)
			currentDiskBlock = solutions.Day09DiskBlocks{}
			continue
		}
	}
	if currentDiskBlock.FileBlocks != 0 {
		parser.diskBlocks = append(parser.diskBlocks, currentDiskBlock)
	}
}

func (parser *parser) toDay9Part1() solutions.Day09Part01 {
	return solutions.Day09Part01{
		DiskBlocks: parser.diskBlocks,
	}
}

func (parser *parser) toDay9Part2() solutions.Day09Part02 {
	return solutions.Day09Part02{
		DiskBlocks: parser.diskBlocks,
	}
}

func Day9ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		diskBlocks: make([]solutions.Day09DiskBlocks, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay9Part1()
	part2 := parser.toDay9Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
