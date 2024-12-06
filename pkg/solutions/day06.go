package solutions

import (
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type Day06Part01 struct {
	XLength       int
	YLength       int
	StartPosition aocmath.Vertex
	Obstacles     map[aocmath.Vertex]bool
}
type Day06Part02 struct {
	XLength       int
	YLength       int
	StartPosition aocmath.Vertex
	Obstacles     map[aocmath.Vertex]bool
}

func (data Day06Part02) isPositionAnObstacle(position aocmath.Vertex, additionalObstacle aocmath.Vertex) bool {
	return data.Obstacles[position] || additionalObstacle == position
}

func (data *Day06Part01) Exec() (string, error) {
	visited := make(map[aocmath.Vertex]bool)
	direction := aocmath.Vertex{X: 0, Y: -1}
	position := data.StartPosition
	for {
		visited[position] = true
		nextPosition := calculateNextPosition(position, direction)
		if data.Obstacles[nextPosition] {
			direction = rotateRight(direction)
			continue
		} else if nextPosition.X >= data.XLength || nextPosition.Y >= data.YLength || nextPosition.X < 0 || nextPosition.Y < 0 {
			break
		}
		position = nextPosition
	}

	return strconv.Itoa(len(visited)), nil
}

func (data *Day06Part02) Exec() (string, error) {
	possibleLoops := 0
	for y := 0; y < data.YLength; y++ {
		for x := 0; x < data.XLength; x++ {
			if data.Obstacles[aocmath.Vertex{X: x, Y: y}] {
				continue
			}
			direction := aocmath.Vertex{X: 0, Y: -1}
			position := data.StartPosition
			visitedObstacles := make(map[VisitedObstacle]bool)
			additionalObstacle := aocmath.Vertex{X: x, Y: y}
			for {
				nextPosition := calculateNextPosition(position, direction)
				if data.isPositionAnObstacle(nextPosition, additionalObstacle) {
					direction = rotateRight(direction)
					alreadyVisited := visitedObstacles[VisitedObstacle{direction: direction, obstaclePosition: nextPosition}]
					visitedObstacles[VisitedObstacle{direction: direction, obstaclePosition: nextPosition}] = true
					if alreadyVisited {
						possibleLoops++
						break
					}
					continue
				} else if nextPosition.X >= data.XLength || nextPosition.Y >= data.YLength || nextPosition.X < 0 || nextPosition.Y < 0 {
					break
				}
				position = nextPosition

			}

		}
	}

	return strconv.Itoa(possibleLoops), nil
}

type VisitedObstacle struct {
	obstaclePosition aocmath.Vertex
	direction        aocmath.Vertex
}

func calculateNextPosition(position aocmath.Vertex, direction aocmath.Vertex) aocmath.Vertex {
	return position.Add(direction)
}

func rotateRight(direction aocmath.Vertex) aocmath.Vertex {
	switch direction {
	case aocmath.Vertex{X: 1, Y: 0}:
		return aocmath.Vertex{X: 0, Y: 1}
	case aocmath.Vertex{X: 0, Y: 1}:
		return aocmath.Vertex{X: -1, Y: 0}
	case aocmath.Vertex{X: -1, Y: 0}:
		return aocmath.Vertex{X: 0, Y: -1}
	case aocmath.Vertex{X: 0, Y: -1}:
		return aocmath.Vertex{X: 1, Y: 0}
	}
	return direction
}
