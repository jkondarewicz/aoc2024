package solution

import (
	"fmt"

	day01 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/01"
	day02 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/02"
	day03 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/03"
	day04 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/04"
	day05 "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/05"
	solutionTypes "github.comjkondarewicz/aoc2024/cmd/aoc2024/days/model"
	"github.comjkondarewicz/aoc2024/internal/benchmark"
)


func PerformAdventOfCode() {
	daysMetadata := []dayMetadata{
		{dir: "cmd/aoc2024/days/01/", resolverProvider: day01.Day1Resolver {}, part1TestSolution: "11", part2TestSolution: "31"},
		{dir: "cmd/aoc2024/days/02/", resolverProvider: day02.Day2Resolver {}, part1TestSolution: "2", part2TestSolution: "4"},
		{dir: "cmd/aoc2024/days/03/", resolverProvider: day03.Day3Resolver {}, part1TestSolution: "161", part2TestSolution: "48"},
		{dir: "cmd/aoc2024/days/04/", resolverProvider: day04.Day4Resolver {}, part1TestSolution: "18", part2TestSolution: "9"},
		{dir: "cmd/aoc2024/days/05/", resolverProvider: day05.Day5Resolver {}, part1TestSolution: "143", part2TestSolution: ""},
	}
	for index, dayMetadata := range daysMetadata {
		day := index + 1
		testDayResolver, error := dayMetadata.resolverProvider.ProvideDayResolver(fmt.Sprintf("%s%s", dayMetadata.dir, "test"))
		if error != nil {
			printError("Error occured during ProvideDayResolver for test case", error)
			break
		}
		dayResolver, error := dayMetadata.resolverProvider.ProvideDayResolver(fmt.Sprintf("%s%s", dayMetadata.dir, "real"))
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
	resolverProvider  solutionTypes.DayResolverProvider
	part1TestSolution string
	part2TestSolution string
}
