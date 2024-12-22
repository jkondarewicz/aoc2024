package solutions

import (
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day20Part01 struct {
	race race
}

type Day20Part02 struct {
	race race
}

type race struct {
	obstacles  *utils.Set[aocmath.Vertex]
	start, end aocmath.Vertex
}
type racePath struct {
	position aocmath.Vertex
	at       int
}

func (data *Day20Part01) Exec() (string, error) {
	return strconv.Itoa(data.race.solvePart1(data)), nil
}

func (data *Day20Part02) Exec() (string, error) {
	return strconv.Itoa(data.race.solvePart2(data)), nil
}

func (r *race) solvePart1(data *Day20Part01) int {
	visited := make(map[aocmath.Vertex]racePath)
	cheats := make([]racePath, 0)
	pos := data.race.start
	dirs := []aocmath.Vertex{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}
	for {
		visited[pos] = racePath{position: pos, at: len(visited)}
		if pos == data.race.end {
			break
		}
		nextPosition := pos
		for _, dir := range dirs {
			np := pos.Add(dir)
			npp := np.Add(dir)
			nppp := npp.Add(dir)
			npo := data.race.obstacles.Exists(np)
			_, alreadyVisited := visited[np]
			if !npo && !alreadyVisited {
				nextPosition = np
			}
			if npo && !data.race.obstacles.Exists(npp) {
				cheats = append(cheats, racePath{position: npp, at: len(visited) + 1})
			}
			if npo && data.race.obstacles.Exists(npp) && !data.race.obstacles.Exists(nppp) {
				cheats = append(cheats, racePath{position: nppp, at: len(visited) + 2})
			}
		}
		pos = nextPosition
	}
	hm := 0
	for _, cheat := range cheats {
		if visited, e := visited[cheat.position]; e {
			cheatSavedSeconds := visited.at - cheat.at
			if cheatSavedSeconds >= 100 {
				hm++
			}
		}
	}
	return hm
}
func (r *race) solvePart2(data *Day20Part02) int {
	visited := make(map[aocmath.Vertex]racePath)
	cheats := make([]racePath, 0)
	pos := data.race.start
	dirs := []aocmath.Vertex{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}
	for {
		visited[pos] = racePath{position: pos, at: len(visited)}
		if pos == data.race.end {
			break
		}
		nextPosition := pos
		for x := pos.X - 20; x <= pos.X+20; x++ {
			for y := pos.Y - 20; y <= pos.Y+20; y++ {
				diff := aocmath.Abs(pos.X-x) + aocmath.Abs(pos.Y-y)
				if diff <= 20 {
					np := aocmath.NewVertex(x, y)
					if !data.race.obstacles.Exists(np) {
						cheats = append(cheats, racePath{position: np, at: len(visited) + diff - 1})
					}
				}
			}
		}
		for _, dir := range dirs {
			np := pos.Add(dir)
			npo := data.race.obstacles.Exists(np)
			_, alreadyVisited := visited[np]
			if !npo && !alreadyVisited {
				nextPosition = np
				break
			}
		}
		pos = nextPosition
	}
	hm := 0
	for _, cheat := range cheats {
		if visited, e := visited[cheat.position]; e {
			cheatSavedSeconds := visited.at - cheat.at
			if cheatSavedSeconds >= 100 {
				hm++
			}
		}
	}
	return hm
}

func CreateDay20(obstacles *utils.Set[aocmath.Vertex], start, end aocmath.Vertex) (Day20Part01, Day20Part02) {
	r := race{obstacles: obstacles, start: start, end: end}
	return Day20Part01{
			race: r,
		},
		Day20Part02{
			race: r,
		}
}
