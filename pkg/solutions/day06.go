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

func (data *Day06Part01) Exec() (string, error) {
	visited := make(map[aocmath.Vertex]bool)
	direction := aocmath.Vertex{X: 0, Y: -1}
	position := data.StartPosition
	for {
		visited[position] = true
		nextPosition := position.Add(direction)
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
	obstaclesContainer := data.toObstacleContainer()
	direction := aocmath.Vertex{X: 0, Y: -1}
	position := data.StartPosition
	visitedPosition := make(map[aocmath.Vertex]bool)
	obstaclesGeneratingLoop := make(map[aocmath.Vertex]bool)
	for {
		nextPosition := position.Add(direction)
		if data.Obstacles[nextPosition] {
			direction = rotateRight(direction)
			continue
		} else if nextPosition.X >= data.XLength || nextPosition.Y >= data.YLength || nextPosition.X < 0 || nextPosition.Y < 0 {
			break
		} else if !visitedPosition[nextPosition] &&
			nextPositionObstacleCreatesLoop(position, nextPosition, rotateRight(direction), obstaclesContainer) {
			obstaclesGeneratingLoop[nextPosition] = true
		}
		visitedPosition[position] = true
		position = nextPosition
	}
	return strconv.Itoa(len(obstaclesGeneratingLoop)), nil
}

func nextPositionObstacleCreatesLoop(
	startPosition aocmath.Vertex,
	additionalObstacle aocmath.Vertex,
	direction aocmath.Vertex,
	obstacleContainer obstaclesContainer,
) bool {
	visitedObstacles := make(map[visitedObstacle]bool)
	position := startPosition
	for {
		var nextObstacle nextObstacle
		if isHorizontalMove(direction) {
			nextObstacle = obstacleContainer.xObstacles[position]
		} else {
			nextObstacle = obstacleContainer.yObstacles[position]
		}
		nextObstaclePosition, isObstacle := nextObstacle.nextObstacle(position, direction, additionalObstacle)
		if !isObstacle {
			return false
		}
		visitedObstacle := visitedObstacle{obstacle: nextObstaclePosition, direction: direction}
		if visitedObstacles[visitedObstacle] {
			return true
		}
		visitedObstacles[visitedObstacle] = true
		position = nextObstaclePosition.Add(direction.Opposite())
		direction = rotateRight(direction)
	}
}

type visitedObstacle struct {
	obstacle  aocmath.Vertex
	direction aocmath.Vertex
}

func rotateRight(direction aocmath.Vertex) aocmath.Vertex {
	return aocmath.Vertex{X: -direction.Y, Y: direction.X}
}
func isHorizontalMove(direction aocmath.Vertex) bool {
	return direction.X != 0
}

type nextObstacle interface {
	nextObstacle(currentPosition aocmath.Vertex, direction aocmath.Vertex, additionalObstacle aocmath.Vertex) (aocmath.Vertex, bool)
}
type xObstacle struct {
	closestLeft  int
	closestRight int
}

func (xObstacle xObstacle) nextObstacle(currentPosition aocmath.Vertex, direction aocmath.Vertex, additionalObstacle aocmath.Vertex) (aocmath.Vertex, bool) {
	sameYAxis := additionalObstacle.Y == currentPosition.Y
	if direction.X == -1 {
		closestLeft := xObstacle.closestLeft
		if sameYAxis && additionalObstacle.X > closestLeft && additionalObstacle.X < currentPosition.X {
			closestLeft = additionalObstacle.X
		}
		return aocmath.Vertex{X: closestLeft, Y: currentPosition.Y}, closestLeft != -1
	} else if direction.X == 1 {
		closestRight := xObstacle.closestRight
		if sameYAxis && additionalObstacle.X > currentPosition.X && (closestRight == -1 || additionalObstacle.X < closestRight) {
			closestRight = additionalObstacle.X
		}
		return aocmath.Vertex{X: closestRight, Y: currentPosition.Y}, closestRight != -1
	} else {
		panic("Should not happen")
	}
}

type yObstacle struct {
	closestUp   int
	closestDown int
}

func (yObstacle yObstacle) nextObstacle(currentPosition aocmath.Vertex, direction aocmath.Vertex, additionalObstacle aocmath.Vertex) (aocmath.Vertex, bool) {
	sameXAxis := additionalObstacle.X == currentPosition.X
	if direction.Y == -1 {
		closestUp := yObstacle.closestUp
		if sameXAxis && additionalObstacle.Y > closestUp && additionalObstacle.Y < currentPosition.Y {
			closestUp = additionalObstacle.Y
		}
		return aocmath.Vertex{X: currentPosition.X, Y: closestUp}, closestUp != -1
	} else if direction.Y == 1 {
		closestDown := yObstacle.closestDown
		if sameXAxis && (closestDown == -1 || additionalObstacle.Y < closestDown) && additionalObstacle.Y > currentPosition.Y {
			closestDown = additionalObstacle.Y
		}
		return aocmath.Vertex{X: currentPosition.X, Y: closestDown}, closestDown != -1
	} else {
		panic("Should not happen")
	}
}

type obstaclesContainer struct {
	xObstacles map[aocmath.Vertex]xObstacle
	yObstacles map[aocmath.Vertex]yObstacle
}

func (data Day06Part02) toObstacleContainer() obstaclesContainer {
	xObstacles := make(map[aocmath.Vertex]xObstacle)
	yObstacles := make(map[aocmath.Vertex]yObstacle)
	for y := 0; y < data.YLength; y++ {
		leftObstacle := -1
		rightObstacle := -1
		for x := 0; x < data.XLength; x++ {
			currentPosition := aocmath.Vertex{X: x, Y: y}
			if data.Obstacles[aocmath.Vertex{X: x, Y: y}] {
				xObstacles[currentPosition] = xObstacle{closestLeft: 0, closestRight: 0}
				leftObstacle = currentPosition.X
			} else {
				xObstacle := xObstacles[currentPosition]
				xObstacle.closestLeft = leftObstacle
				xObstacles[currentPosition] = xObstacle
			}
		}
		for x := data.XLength - 1; x >= 0; x-- {
			currentPosition := aocmath.Vertex{X: x, Y: y}
			if data.Obstacles[aocmath.Vertex{X: x, Y: y}] {
				xObstacles[currentPosition] = xObstacle{closestLeft: 0, closestRight: 0}
				rightObstacle = currentPosition.X
			} else {
				xObstacle := xObstacles[currentPosition]
				xObstacle.closestRight = rightObstacle
				xObstacles[currentPosition] = xObstacle
			}
		}
	}
	for x := 0; x < data.XLength; x++ {
		upObstacle := -1
		downObstacle := -1
		for y := 0; y < data.YLength; y++ {
			currentPosition := aocmath.Vertex{X: x, Y: y}
			if data.Obstacles[aocmath.Vertex{X: x, Y: y}] {
				yObstacles[currentPosition] = yObstacle{closestUp: 0, closestDown: 0}
				upObstacle = currentPosition.Y
			} else {
				yObstacle := yObstacles[currentPosition]
				yObstacle.closestUp = upObstacle
				yObstacles[currentPosition] = yObstacle
			}
		}
		for y := data.YLength - 1; y >= 0; y-- {
			currentPosition := aocmath.Vertex{X: x, Y: y}
			if data.Obstacles[aocmath.Vertex{X: x, Y: y}] {
				yObstacles[currentPosition] = yObstacle{closestUp: 0, closestDown: 0}
				downObstacle = currentPosition.Y
			} else {
				yObstacle := yObstacles[currentPosition]
				yObstacle.closestDown = downObstacle
				yObstacles[currentPosition] = yObstacle
			}
		}
	}
	return obstaclesContainer{
		xObstacles: xObstacles,
		yObstacles: yObstacles,
	}
}
