package day05

import (
	"regexp"
	"strconv"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/files"
	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/solutions"
)

type parser struct {
	order []aocmath.Vertex
	pages [][]int
}

func (parser *parser) ReadLine(index int, line string) {
	isOrderRegex := regexp.MustCompile(`\d+\|\d+`)
	numRegex := regexp.MustCompile(`\d+`)
	if isOrderRegex.MatchString(line) {
		// fmt.Printf("Matching line %s\n", line)
		orders := numRegex.FindAllString(line, -1)
		x, _ := strconv.Atoi(orders[0])
		y, _ := strconv.Atoi(orders[1])
		parser.order = append(parser.order, aocmath.Vertex{X: x, Y: y})
	} else if len(line) > 0 {
		pagesString := numRegex.FindAllString(line, -1)
		pages := make([]int, len(pagesString))
		for index, page := range pagesString {
			pageInt, _ := strconv.Atoi(page)
			pages[index] = pageInt 
		}
		parser.pages = append(parser.pages, pages)
	}
}

func (parser *parser) toDay5Part1() solutions.Day05Part01 {
	return solutions.Day05Part01{
		Rules: parser.order,
		Pages: parser.pages,
	}
}

func (parser *parser) toDay5Part2() solutions.Day05Part02 {
	return solutions.Day05Part02{
		Rules: parser.order,
		Pages: parser.pages,
	}
}

func Day5ResolverProvide(filename string, test bool) (solutionTypes.DayResolver, error) {
	file, err := files.Open(filename)
	if err != nil {
		return solutionTypes.DayResolver{}, err
	}
	parser := parser{
		order: make([]aocmath.Vertex, 0),
		pages: make([][]int, 0),
	}
	file.ProcessLineByLine(&parser)
	part1 := parser.toDay5Part1()
	part2 := parser.toDay5Part2()
	return solutionTypes.DayResolver{
		ResolvePart1Function: &part1,
		ResolvePart2Function: &part2,
	}, nil
}
