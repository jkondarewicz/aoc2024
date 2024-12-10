package solutions

import (
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type Day10Part01 struct {
	StartPositions []aocmath.Vertex
	MapTrail       map[aocmath.Vertex]int
	MapWidth       int
	MapHeight      int
}
type Day10Part02 struct {
	StartPositions []aocmath.Vertex
	MapTrail       map[aocmath.Vertex]int
	MapWidth       int
	MapHeight      int
}
type hikingMap struct {
	MapTrail      map[aocmath.Vertex]int
	MapWidth      int
	MapHeight     int
	DistinctTrail bool
}

func (hikingMap hikingMap) howManyReachable(curPos aocmath.Vertex, previousHeight int, visited map[aocmath.Vertex]struct{}) int {
	if curPos.X < 0 || curPos.X >= hikingMap.MapWidth || curPos.Y < 0 || curPos.Y >= hikingMap.MapHeight {
		return 0
	}
	if hikingMap.MapTrail[curPos]-previousHeight != 1 {
		return 0
	}
	if _, exists := visited[curPos]; exists && hikingMap.DistinctTrail {
		return 0
	}
	if hikingMap.MapTrail[curPos] == 9 {
		visited[curPos] = struct{}{}
		return 1
	}
	visited[curPos] = struct{}{}
	directions := []aocmath.Vertex{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}
	result := 0
	for _, direction := range directions {
		result += hikingMap.howManyReachable(curPos.Add(direction), hikingMap.MapTrail[curPos], visited)
	}
	return result
}

func (data *Day10Part01) Exec() (string, error) {
	hikingMap := hikingMap{MapTrail: data.MapTrail, MapWidth: data.MapWidth, MapHeight: data.MapHeight, DistinctTrail: true}
	result := 0
	for _, position := range data.StartPositions {
		t := hikingMap.howManyReachable(position, -1, make(map[aocmath.Vertex]struct{}))
		result += t
	}
	return strconv.Itoa(result), nil
}

func (data *Day10Part02) Exec() (string, error) {
	hikingMap := hikingMap{MapTrail: data.MapTrail, MapWidth: data.MapWidth, MapHeight: data.MapHeight}
	result := 0
	for _, position := range data.StartPositions {
		t := hikingMap.howManyReachable(position, -1, make(map[aocmath.Vertex]struct{}))
		result += t
	}
	return strconv.Itoa(result), nil
}
