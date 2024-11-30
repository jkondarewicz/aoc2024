package benchmark

import "time"

type BenchmarkExec[T any] interface {
	Exec() (T, error)
}

type BenchmarkResult[T any] struct {
	Result T
	Time int64
}

func Benchmark[T any](exec BenchmarkExec[T]) (BenchmarkResult[T], error) {
	startTime := time.Now().UnixMilli()
	executed, err := exec.Exec()
	endTime := time.Now().UnixMilli()
	if err != nil {
		return BenchmarkResult[T] {}, err
	}
	return BenchmarkResult[T] {
		Result: executed,
		Time: endTime - startTime,
	}, nil
}
