package solution

import (
	"fmt"

	day01 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/01"
	day02 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/02"
	day03 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/03"
	day04 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/04"
	day05 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/05"
	day06 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/06"
	day07 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/07"
	day08 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/08"
	day09 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/09"
	day10 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/10"
	day11 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/11"
	day12 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/12"
	day13 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/13"
	day14 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/14"
	day15 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/15"
	day16 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/16"
	day17 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/17"
	day18 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/18"
	day19 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/19"
	day20 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/20"
	day21 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/21"
	day22 "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/22"

	solutionTypes "github.com/jkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.com/jkondarewicz/aoc2024/internal/benchmark"
)

func PerformAdventOfCode() {
	daysMetadata := []dayMetadata{
		{day: 1, dir: "cmd/aoc2024/days/01/", resolverProvider: day01.Day1ResolverProvide, part1TestSolution: "11", part2TestSolution: "31"},
		{day: 2, dir: "cmd/aoc2024/days/02/", resolverProvider: day02.Day2ResolverProvide, part1TestSolution: "2", part2TestSolution: "4"},
		{day: 3, dir: "cmd/aoc2024/days/03/", resolverProvider: day03.Day3ResolverProvide, part1TestSolution: "161", part2TestSolution: "48"},
		{day: 4, dir: "cmd/aoc2024/days/04/", resolverProvider: day04.Day4ResolverProvide, part1TestSolution: "18", part2TestSolution: "9"},
		{day: 5, dir: "cmd/aoc2024/days/05/", resolverProvider: day05.Day5ResolverProvide, part1TestSolution: "143", part2TestSolution: "123"},
		{day: 6, dir: "cmd/aoc2024/days/06/", resolverProvider: day06.Day6ResolverProvide, part1TestSolution: "41", part2TestSolution: "6"},
		{day: 7, dir: "cmd/aoc2024/days/07/", resolverProvider: day07.Day7ResolverProvide, part1TestSolution: "3749", part2TestSolution: "11387"},
		{day: 8, dir: "cmd/aoc2024/days/08/", resolverProvider: day08.Day8ResolverProvide, part1TestSolution: "14", part2TestSolution: "34"},
		{day: 9, dir: "cmd/aoc2024/days/09/", resolverProvider: day09.Day9ResolverProvide, part1TestSolution: "1928", part2TestSolution: "2858"},
		{day: 10, dir: "cmd/aoc2024/days/10/", resolverProvider: day10.Day10ResolverProvide, part1TestSolution: "36", part2TestSolution: "81"},
		{day: 11, dir: "cmd/aoc2024/days/11/", resolverProvider: day11.Day11ResolverProvide, part1TestSolution: "55312", part2TestSolution: "65601038650482"},
		{day: 12, dir: "cmd/aoc2024/days/12/", resolverProvider: day12.Day12ResolverProvide, part1TestSolution: "1184", part2TestSolution: "368"},
		{day: 13, dir: "cmd/aoc2024/days/13/", resolverProvider: day13.Day13ResolverProvide, part1TestSolution: "480", part2TestSolution: "875318608908"},
		{day: 14, dir: "cmd/aoc2024/days/14/", resolverProvider: day14.Day14ResolverProvide, part1TestSolution: "12", part2TestSolution: "1"},
		{day: 15, dir: "cmd/aoc2024/days/15/", resolverProvider: day15.Day15ResolverProvide, part1TestSolution: "10092", part2TestSolution: "9021"},
		{day: 16, dir: "cmd/aoc2024/days/16/", resolverProvider: day16.Day16ResolverProvide, part1TestSolution: "7036", part2TestSolution: "45"},
		{day: 17, dir: "cmd/aoc2024/days/17/", resolverProvider: day17.Day17ResolverProvide, part1TestSolution: "6,4,6,0,4,5,7,2,7", part2TestSolution: "164541160582845"},
		{day: 18, dir: "cmd/aoc2024/days/18/", resolverProvider: day18.Day18ResolverProvide, part1TestSolution: "22", part2TestSolution: "6,1"},
		{day: 19, dir: "cmd/aoc2024/days/19/", resolverProvider: day19.Day19ResolverProvide, part1TestSolution: "6", part2TestSolution: "16"},
		{day: 20, dir: "cmd/aoc2024/days/20/", resolverProvider: day20.Day20ResolverProvide, part1TestSolution: "0", part2TestSolution: "0"},
		{day: 21, dir: "cmd/aoc2024/days/21/", resolverProvider: day21.Day21ResolverProvide, part1TestSolution: "126384", part2TestSolution: "16"},
		{day: 22, dir: "cmd/aoc2024/days/22/", resolverProvider: day22.Day22ResolverProvide, part1TestSolution: "37990510", part2TestSolution: "23"},
	}
	for _, dayMetadata := range daysMetadata {
		if omitDay[dayMetadata.day] {
			continue
		}
		testDayResolver, error := dayMetadata.resolverProvider(fmt.Sprintf("%s%s", dayMetadata.dir, "test"), true)
		if error != nil {
			printError("Error occured during ProvideDayResolver for test case", error)
			break
		}
		dayResolver, error := dayMetadata.resolverProvider(fmt.Sprintf("%s%s", dayMetadata.dir, "real"), false)
		if error != nil {
			printError("Error occured during ProvideDayResolver for real case", error)
			break
		}

		part1Solution, error := testDayResolver.ResolvePart1Function.Exec()
		if error != nil {
			printError("Error occured during ResolvePart1Function", error)
			break
		}
		if part1Solution != dayMetadata.part1TestSolution {
			fmt.Printf("Day %d part 1. Expected: [%s], got: [%s]", dayMetadata.day, dayMetadata.part1TestSolution, part1Solution)
			break
		}
		fmt.Printf("Day %d part 1 test passed\n", dayMetadata.day)
		benchmarkedPart1, error := benchmark.Benchmark(dayResolver.ResolvePart1Function)
		printSolution(dayMetadata.day, 1, error, benchmarkedPart1)

		part2Solution, error := testDayResolver.ResolvePart2Function.Exec()
		if error != nil {
			printError("Error occured during ResolvePart2Function", error)
			break
		}
		if part2Solution != dayMetadata.part2TestSolution {
			fmt.Printf("Day %d part 2. Expected: [%s], got: [%s]", dayMetadata.day, dayMetadata.part2TestSolution, part2Solution)
			break
		}
		fmt.Printf("Day %d part 2 test passed\n", dayMetadata.day)
		benchmarkedPart2, error := benchmark.Benchmark(dayResolver.ResolvePart2Function)
		printSolution(dayMetadata.day, 2, error, benchmarkedPart2)
	}

}

func printError(message string, error error) {
	fmt.Printf("%s: %v, %T", message, error, error)
}

func printSolution(day int, part int, err error, result benchmark.BenchmarkResult[string]) {
	if err != nil {
		fmt.Printf("Day %d, part %d, error [%s]\n", day, part, err)
	} else {
		fmt.Printf("Day %d, part %d, solution took %dms\n", day, part, result.Time)
		fmt.Printf("%s\n", result.Result)
		fmt.Println()
	}
}

type dayMetadata struct {
	day               int
	dir               string
	resolverProvider  func(filename string, test bool) (solutionTypes.DayResolver, error)
	part1TestSolution string
	part2TestSolution string
}
var omitDay map[int]bool = map[int]bool{
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
	8: true,
	9: true,
	10: true,
	11: true,
	12: true,
	13: true,
	14: true,
	15: true,
	16: true,
	17: true,
	18: true,
	19: true,
	20: false,
	21: true,
	22: true,
	23: true,
	24: true,
	25: true,
}
