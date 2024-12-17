package solutions

import (
	"container/heap"
	"math"
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
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
	paths := data.Map.findPath(data.StartPosition, data.EndPosition)
	return strconv.Itoa(paths[0].cost), nil
}

func (data *Day16Part02) Exec() (string, error) {
	paths := data.Map.findPath(data.StartPosition, data.EndPosition)
	uniquePaths := utils.NewSet[aocmath.Vertex]()
	for _, path := range paths {
		path.visited.ForEach(func(value aocmath.Vertex) {
			uniquePaths.Add(value)
		})
	}
	return strconv.Itoa(uniquePaths.Size()), nil
}

var mapDirections = []aocmath.Vertex{
	aocmath.NewVertex(1, 0),
	aocmath.NewVertex(-1, 0),
	aocmath.NewVertex(0, 1),
	aocmath.NewVertex(0, -1),
}

func (m *day16Map) findPath(start, end aocmath.Vertex) []state {
	pq := &PriorityQueue{}
	heap.Init(pq)
	currentState := &state{position: start, direction: aocmath.NewVertex(1, 0), cost: 0, visited: &utils.Set[aocmath.Vertex]{}}
	currentState.visited = utils.NewSet[aocmath.Vertex]()
	heap.Push(pq, currentState)
	minCosts := make(map[aocmath.Vertex]map[aocmath.Vertex]int)
	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			minCosts[aocmath.NewVertex(x, y)] = make(map[aocmath.Vertex]int)
			for _, dir := range mapDirections {
				minCosts[aocmath.NewVertex(x, y)][dir] = math.MaxInt64
			}
		}
	}
	foundPaths := make([]state, 0)
	minCost := math.MaxInt64
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*state)
		if current.position == end {
			minCost = aocmath.Min(minCost, current.cost)
			current.visited.Add(current.position)
			foundPaths = append(foundPaths, *current)
			continue
		}
		if current.visited.Exists(current.position) || current.cost > minCost {
			continue
		}
		current.visited.Add(current.position)
		for i := 0; i < 3; i++ {
			direction := current.direction
			if i == 1 {
				direction = rotateLeft(direction)
			} else if i == 2 {
				direction = rotateRight(direction)
			}
			nextPosition := direction.Add(current.position)
			if m.obstacles.Exists(nextPosition) {
				continue
			}
			newCost := current.cost + 1
			if direction != current.direction {
				newCost += 1000
			}
			if newCost > minCosts[nextPosition][direction] {
				continue
			}
			minCosts[nextPosition][direction] = newCost
			newVisted := utils.NewSet[aocmath.Vertex]()
			current.visited.ForEach(func(v aocmath.Vertex) {
				newVisted.Add(v)
			})
			heap.Push(pq, &state{position: nextPosition, direction: direction, cost: newCost, visited: newVisted})
		}
	}
	p := make([]state, 0)
	for _, s := range foundPaths {
		if s.cost == minCost {
			p = append(p, s)
		}
	}
	return p
}

type state struct {
	position, direction aocmath.Vertex
	cost                int
	visited             *utils.Set[aocmath.Vertex]
}

type PriorityQueue []*state

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*state))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
