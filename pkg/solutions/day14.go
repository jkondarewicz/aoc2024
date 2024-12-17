package solutions

import (
	"fmt"
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day14Part01 struct {
	Width int
	Height int
	Robots []Day14Robot

}
type Day14Part02 struct {
	Width int
	Height int
	Robots []Day14Robot
}

type Day14Robot struct {
	StartPosition aocmath.Vertex
	Transition aocmath.Vertex
}

func (data *Day14Part01) Exec() (string, error) {
	quadrants := make(map[aocmath.Vertex]int)
	for _, robot := range data.Robots {
		quadrants[robot.positionRelativeToCenterNormalized(100, data.Width, data.Height)]++
	}
	tl, tr, br, bl := 
		quadrants[aocmath.NewVertex(-1, -1)],
		quadrants[aocmath.NewVertex(1, -1)],
		quadrants[aocmath.NewVertex(1, 1,)],
		quadrants[aocmath.NewVertex(-1, 1)]
	return strconv.Itoa(tl * tr * br * bl), nil
}

func (data *Day14Part02) Exec() (string, error) {
	seconds := 1
	for {
		robotsPositions := utils.NewSet[aocmath.Vertex]()
		for _, robot := range data.Robots {
			robotsPositions.Add(robot.positionAfterMoves(seconds, data.Width, data.Height))
		}
		if robotsPositions.Size() == len(data.Robots) {
			printArea(robotsPositions, data.Width, data.Height)
			break
		}
		seconds++
	}
	return strconv.Itoa(seconds), nil
}

func printArea(robotsPositions *utils.Set[aocmath.Vertex], width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rune := '.'
			if robotsPositions.Exists(aocmath.NewVertex(x, y)) {
				rune = '#'
			}
			fmt.Printf("%c", rune)
		}
		fmt.Println()
	}
}

func (robot *Day14Robot) positionRelativeToCenterNormalized(moves int, width int, height int) aocmath.Vertex {
	cp := robot.positionAfterMoves(moves, width, height)
	cx, cy := width / 2, height / 2
	quadrantPlace := aocmath.NewVertex(cp.X - cx, cp.Y - cy).Normalize()
	return quadrantPlace
}

func (robot *Day14Robot) positionAfterMoves(moves int, width int, height int) aocmath.Vertex {
	tx, ty := (robot.Transition.X * moves) % width, (robot.Transition.Y * moves) % height
	cp := robot.StartPosition.Add(aocmath.NewVertex(tx, ty))
	cp = aocmath.NewVertex(cp.X % width, cp.Y % height)
	if cp.X < 0 {
		cp.X = width + cp.X
	}
	if cp.Y < 0 {
		cp.Y = height + cp.Y
	}
	return cp
}
