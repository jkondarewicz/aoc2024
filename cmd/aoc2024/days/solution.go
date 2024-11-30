package solution

import "github.comjkondarewicz/aoc2024/internal/benchmark"

type Solution interface {
	Part1Solution() (benchmark.BenchmarkResult[string], error)
	Part2Solution() (benchmark.BenchmarkResult[string], error)
}

type ErrorSolution struct {
	Error error
}
func (e ErrorSolution) Part1Solution() (benchmark.BenchmarkResult[string], error) {
	return benchmark.BenchmarkResult[string] {}, e.Error
}
func (e ErrorSolution) Part2Solution() (benchmark.BenchmarkResult[string], error) {
	return benchmark.BenchmarkResult[string] {}, e.Error
}

