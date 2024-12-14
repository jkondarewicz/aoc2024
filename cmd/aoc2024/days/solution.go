package solution

import (
	"fmt"

	day01 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/01"
	day02 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/02"
	day03 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/03"
	day04 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/04"
	day05 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/05"
	day06 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/06"
	day07 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/07"
	day08 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/08"
	day09 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/09"
	day10 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/10"
	day11 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/11"
	day12 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/12"
	day13 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/13"
	day14 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/14"
	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/benchmark"
)


func PerformAdventOfCode() {
	daysMetadata := []dayMetadata{
		{dir: "cmd/aoc2024/days/01/", resolverProvider: day01.Day1ResolverProvide, part1TestSolution: "11", part2TestSolution: "31"},
		{dir: "cmd/aoc2024/days/02/", resolverProvider: day02.Day2ResolverProvide, part1TestSolution: "2", part2TestSolution: "4"},
		{dir: "cmd/aoc2024/days/03/", resolverProvider: day03.Day3ResolverProvide, part1TestSolution: "161", part2TestSolution: "48"},
		{dir: "cmd/aoc2024/days/04/", resolverProvider: day04.Day4ResolverProvide, part1TestSolution: "18", part2TestSolution: "9"},
		{dir: "cmd/aoc2024/days/05/", resolverProvider: day05.Day5ResolverProvide, part1TestSolution: "143", part2TestSolution: "123"},
		{dir: "cmd/aoc2024/days/06/", resolverProvider: day06.Day6ResolverProvide, part1TestSolution: "41", part2TestSolution: "6"},
		{dir: "cmd/aoc2024/days/07/", resolverProvider: day07.Day7ResolverProvide, part1TestSolution: "3749", part2TestSolution: "11387"},
		{dir: "cmd/aoc2024/days/08/", resolverProvider: day08.Day8ResolverProvide, part1TestSolution: "14", part2TestSolution: "34"},
		{dir: "cmd/aoc2024/days/09/", resolverProvider: day09.Day9ResolverProvide, part1TestSolution: "1928", part2TestSolution: "2858"},
		{dir: "cmd/aoc2024/days/10/", resolverProvider: day10.Day10ResolverProvide, part1TestSolution: "36", part2TestSolution: "81"},
		{dir: "cmd/aoc2024/days/11/", resolverProvider: day11.Day11ResolverProvide, part1TestSolution: "55312", part2TestSolution: "65601038650482"},
		{dir: "cmd/aoc2024/days/12/", resolverProvider: day12.Day12ResolverProvide, part1TestSolution: "1184", part2TestSolution: "368"},
		{dir: "cmd/aoc2024/days/13/", resolverProvider: day13.Day13ResolverProvide, part1TestSolution: "480", part2TestSolution: "875318608908"},
		{dir: "cmd/aoc2024/days/14/", resolverProvider: day14.Day14ResolverProvide, part1TestSolution: "12", part2TestSolution: "1"},
	}
	for index, dayMetadata := range daysMetadata {
		day := index + 1
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
			fmt.Printf("Day %d part 1. Expected: [%s], got: [%s]", day, dayMetadata.part1TestSolution, part1Solution)
			break
		}
		benchmarkedPart1, error := benchmark.Benchmark(dayResolver.ResolvePart1Function)
		printSolution(day, 1, error, benchmarkedPart1)

		part2Solution, error := testDayResolver.ResolvePart2Function.Exec()
		if error != nil {
			printError("Error occured during ResolvePart2Function", error)
			break
		}
		if part2Solution != dayMetadata.part2TestSolution {
			fmt.Printf("Day %d part 2. Expected: [%s], got: [%s]", day, dayMetadata.part2TestSolution, part2Solution)
			break
		}
		benchmarkedPart2, error := benchmark.Benchmark(dayResolver.ResolvePart2Function)
		printSolution(day, 2, error, benchmarkedPart2)
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
	dir               string
	resolverProvider  func(filename string, test bool) (solutionTypes.DayResolver, error)
	part1TestSolution string
	part2TestSolution string
}
