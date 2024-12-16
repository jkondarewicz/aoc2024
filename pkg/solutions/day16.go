package solutions

import (
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
	"github.comjkondarewicz/aoc2024/pkg/utils"
)

type Day16Part01 struct {
	StartPosition aocmath.Vertex
	EndPosition   aocmath.Vertex
	Map           day16Map
}

type Day16Part02 struct {
	StartPosition aocmath.Vertex
	EndPosition   aocmath.Vertex
	Map           day16Map
}

func CreateMap(width, height int, obstacles *utils.Set[aocmath.Vertex]) day16Map {
	return day16Map{
		width:     width,
		height:    height,
		obstacles: obstacles,
	}
}

type day16Map struct {
	width     int
	height    int
	obstacles *utils.Set[aocmath.Vertex]
}

func (data *Day16Part01) Exec() (string, error) {
	result := data.Map.calculateBestPath(data.StartPosition, data.EndPosition)
	return strconv.Itoa(result), nil
}

func (data *Day16Part02) Exec() (string, error) {
	 // data.Map.calculateBestPath(data.StartPosition, data.EndPosition)
	return strconv.Itoa(0), nil
}

func (m *day16Map) calculateBestPath(
	startPosition aocmath.Vertex,
	endPosition aocmath.Vertex,
) (int) {
	visited := make(map[aocmath.Vertex]int)
	visited[startPosition] = 0
	d := []aocmath.Vertex{
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: 1, Y: 0},
	}
	m.updatePositions(startPosition, endPosition, d[2], d[0], 0, visited)
	m.updatePositions(startPosition, endPosition, d[2], d[2], 0, visited)
	m.updatePositions(startPosition, endPosition, d[2], d[1], 0, visited)
	return visited[endPosition]
}

func (m *day16Map) updatePositions(
	position, endPosition, previousDirection, currentDirection aocmath.Vertex,
	score int,
	visited map[aocmath.Vertex]int,
) {
	currentPosition := position.Add(currentDirection)
	if m.obstacles.Exists(currentPosition) {
		return
	}
	score += 1
	if previousDirection != currentDirection {
		score += 1000
	}
	previousScore, alreadyVisited := visited[currentPosition]
	if alreadyVisited && score > previousScore {
		return
	}
	visited[currentPosition] = score
	if currentPosition == endPosition {
		return
	}
	left := rotateLeft(currentDirection)
	right := rotateRight(currentDirection)
	m.updatePositions(currentPosition, endPosition, currentDirection, currentDirection, score, visited)
	m.updatePositions(currentPosition, endPosition, currentDirection, left, score, visited)
	m.updatePositions(currentPosition, endPosition, currentDirection, right, score, visited)
}
