package solutions

import (
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day15Part01 struct {
	Map   Day15Map
	Robot Day15Robot
}
type Day15Part02 struct {
	Width  int
	Height int
	Map    Day15Warehouse
	Robot  Day15Robot
}

type Day15Map struct {
	Walls *utils.Set[aocmath.Vertex]
	Boxes *utils.Set[aocmath.Vertex]
}
type Day15Warehouse struct {
	Walls *utils.Set[aocmath.Vertex]
	Boxes map[aocmath.Vertex]aocmath.Vertex 
}
type Day15Robot struct {
	Position aocmath.Vertex
	Moves    []Day15Direction
}

type Day15Direction aocmath.Vertex

func (dir *Day15Direction) isVertical() bool {
	return dir.Y != 0
}

func CreateRobotDirection(direction rune) (Day15Direction, bool) {
	switch direction {
	case '^':
		return Day15Direction(aocmath.NewVertex(0, -1)), true
	case 'v':
		return Day15Direction(aocmath.NewVertex(0, 1)), true
	case '<':
		return Day15Direction(aocmath.NewVertex(-1, 0)), true
	case '>':
		return Day15Direction(aocmath.NewVertex(1, 0)), true
	default:
		return Day15Direction{}, false
	}
}

func (data *Day15Part01) Exec() (string, error) {
	cp := data.Robot.Position
	for _, move := range data.Robot.Moves {
		nextFreePosition, exists := data.Map.nextFreeSpace(cp, move)
		if exists {
			data.Map.swapBoxes(nextFreePosition, cp, aocmath.Vertex(move).Opposite())
			cp = cp.Add(aocmath.Vertex(move))
		}
	}
	result := 0
	data.Map.Boxes.ForEach(func(boxPosition aocmath.Vertex) {
		result += (boxPosition.Y*100 + boxPosition.X)
	})
	return strconv.Itoa(result), nil
}

func (data *Day15Part02) Exec() (string, error) {
	data.Robot.Position = data.Robot.Position.Add(aocmath.NewVertex(data.Robot.Position.X, 0))
	cp := data.Robot.Position
	for _, move := range data.Robot.Moves {
		nextMovePossible := data.Map.updateBoxPositions(cp, move)
		if nextMovePossible {
			cp = cp.Add(aocmath.Vertex(move))
		}
	}
	result := 0
	calculated := utils.NewSet[aocmath.Vertex]()
	for _, box := range data.Map.Boxes {
		if !calculated.Exists(box) {
			result += box.Y * 100 + box.X
		}
		calculated.Add(box)
	}
	return strconv.Itoa(result), nil
}

func (warehouse *Day15Warehouse) updateBoxPositions(pos aocmath.Vertex, direction Day15Direction) bool {
	if direction.isVertical() {
		affectedBoxes := utils.NewSet[aocmath.Vertex]()
		positions := make([]aocmath.Vertex, 0)
		positions = append(positions, pos)
		for {
			nPositions := make([]aocmath.Vertex, 0)
			for _, position := range positions {
				np := position.Add(aocmath.Vertex(direction))
				box, boxExists := warehouse.Boxes[np]
				wallExists := warehouse.Walls.Exists(np)
				if !wallExists && !boxExists {
					continue
				} else if wallExists {
					return false
				}
				affectedBoxes.Add(box)
				nPositions = append(nPositions, box)
				nPositions = append(nPositions, box.Add(aocmath.NewVertex(1, 0)))
			}
			positions = nPositions
			if len(positions) == 0 {
				break
			}
		}
		affectedBoxes.ForEach(func(box aocmath.Vertex) {
			delete(warehouse.Boxes, box)
			delete(warehouse.Boxes, box.Add(aocmath.NewVertex(1, 0)))
		})
		affectedBoxes.ForEach(func(box aocmath.Vertex) {
			box = box.Add(aocmath.Vertex(direction))
			warehouse.Boxes[box] = box
			warehouse.Boxes[box.Add(aocmath.NewVertex(1, 0))] = box
		})

		return true
	} else {
		affectedBoxes := utils.NewSet[aocmath.Vertex]()
		for {
			pos = pos.Add(aocmath.Vertex(direction))
			_, boxExists := warehouse.Boxes[pos]
			we := warehouse.Walls.Exists(pos)
			if !we && !boxExists {
				break
			} else if we {
				return false
			}
			affectedBoxes.Add(warehouse.Boxes[pos])
		}
		affectedBoxes.ForEach(func(box aocmath.Vertex) {
			delete(warehouse.Boxes, box)
			delete(warehouse.Boxes, box.Add(aocmath.NewVertex(1, 0)))
		})
		affectedBoxes.ForEach(func(box aocmath.Vertex) {
			box = box.Add(aocmath.Vertex(direction))
			warehouse.Boxes[box] = box
			warehouse.Boxes[box.Add(aocmath.NewVertex(1, 0))] = box
		})
		return true
	}
}
func (warehouse *Day15Map) swapBoxes(freePosition, robotPosition, direction aocmath.Vertex) {
	cp := freePosition
	for {
		np := cp.Add(direction)
		if np == robotPosition {
			return
		}
		warehouse.Boxes.Remove(np)
		warehouse.Boxes.Add(cp)
		cp = np
	}
}

func (warehouse *Day15Map) nextFreeSpace(pos aocmath.Vertex, direction Day15Direction) (aocmath.Vertex, bool) {
	for {
		pos = pos.Add(aocmath.Vertex(direction))
		we, be := warehouse.Walls.Exists(pos), warehouse.Boxes.Exists(pos)
		if !we && !be {
			return pos, true
		} else if we {
			return pos, false
		}
	}
}
