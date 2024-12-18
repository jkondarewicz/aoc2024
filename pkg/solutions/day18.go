package solutions

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day18Part01 struct {
	Width   int
	Height  int
	Bits    []aocmath.Vertex
	HowMany int
}

type Day18Part02 struct {
	Width   int
	Height  int
	Bits    []aocmath.Vertex
	HowMany int
}

func (data *Day18Part01) Exec() (string, error) {
	start := aocmath.NewVertex(0, 0)
	end := aocmath.NewVertex(data.Width-1, data.Height-1)

	corrupted := utils.NewSet[aocmath.Vertex]()
	for i := 0; i < data.HowMany; i++ {
		corrupted.Add(data.Bits[i])
	}
	return strconv.Itoa(getPath(start, end, data.Width, data.Height, corrupted).cost), nil
}

func (data *Day18Part02) Exec() (string, error) {
	start := aocmath.NewVertex(0, 0)
	end := aocmath.NewVertex(data.Width-1, data.Height-1)
	var x, y int

	for additional := data.HowMany; additional < len(data.Bits); additional++ {
		corrupted := utils.NewSet[aocmath.Vertex]()
		for i := 0; i < data.HowMany; i++ {
			corrupted.Add(data.Bits[i])
		}
		for i := data.HowMany; i <= additional; i++ {
			corrupted.Add(data.Bits[i])
		}
		if fpath := getPath(start, end, data.Width, data.Height, corrupted); fpath == nil {
			d := data.Bits[additional]
			x, y = d.X, d.Y
			break
		}
	}
	return fmt.Sprintf("%d,%d", x, y), nil
}

func getPath(start, end aocmath.Vertex, width, height int, corrupted *utils.Set[aocmath.Vertex]) *path {

	p := &path{cost: 0, position: start}
	pq := &pq{}
	heap.Init(pq)
	heap.Push(pq, p)

	d := []aocmath.Vertex{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
	minCosts := make(map[aocmath.Vertex]int)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			minCosts[aocmath.NewVertex(x, y)] = math.MaxInt64
		}
	}
	var fpath *path = nil
	outOfBoard := func(pos aocmath.Vertex) bool {
		return pos.X < 0 || pos.Y < 0 || pos.X >= width || pos.Y >= height
	}
	farFromEnd := func(pos aocmath.Vertex) int {
		return aocmath.Abs(end.X-pos.X) + aocmath.Abs(end.Y-pos.Y)
	}

	for pq.Len() > 0 {
		endCost := minCosts[end]
		current := heap.Pop(pq).(*path)
		if minCosts[current.position] == current.cost {
			continue
		}
		if current.position == end {
			if current.cost < endCost {
				fpath = current
				minCosts[end] = current.cost
			}
			continue
		}
		minCosts[current.position] = current.cost
		for _, dir := range d {
			nextPosition := dir.Add(current.position)
			nextPath := &path{cost: current.cost + 1, position: nextPosition, farFromEnd: farFromEnd(nextPosition)}
			if nextPath.cost >= minCosts[nextPosition] || corrupted.Exists(nextPosition) || outOfBoard(nextPosition) {
				continue
			}
			heap.Push(pq, nextPath)
		}
	}
	return fpath
}

type path struct {
	position   aocmath.Vertex
	cost       int
	farFromEnd int
}

type pq []*path

func (pq pq) Less(i, j int) bool {
	a := pq[i]
	b := pq[j]
	return a.cost < b.cost
}
func (pq pq) Len() int { return len(pq) }
func (pq pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pq) Push(x any) {
	*pq = append(*pq, x.(*path))
}

func (pq *pq) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
