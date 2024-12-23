package solutionTypes

import "github.com/jkondarewicz/aoc2024/internal/benchmark"

type DayResolver struct {
	ResolvePart1Function benchmark.BenchmarkFunction[string]
	ResolvePart2Function benchmark.BenchmarkFunction[string]
}
type DayResolverProvider interface {
	ProvideDayResolver(filename string) (DayResolver, error)
}
