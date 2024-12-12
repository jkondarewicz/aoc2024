package solutions

import (
	"fmt"
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type Day12Part01 struct {
	Width  int
	Height int
	Plants map[aocmath.Vertex]rune //position is 2x wider as need to find perimeters
}
type Day12Part02 struct {
	Width  int
	Height int
	Plants map[aocmath.Vertex]rune
}

type regionFence struct {
	name      rune
	area      int
	perimeter int
}

func (regionFence regionFence) String() string {
	return fmt.Sprintf("name: %c, area: %d, perimeter: %d, region: %d", regionFence.name, regionFence.area, regionFence.perimeter, regionFence.calculateRegion())
}
func (regionFence *regionFence) calculateRegion() int64 {
	return int64(regionFence.area) * int64(regionFence.perimeter)
}

func calculateRegionFences(plants map[aocmath.Vertex]rune, width int, height int) []regionFence {
	regionFences := make([]regionFence, 0)
	visited := make(map[aocmath.Vertex]struct{})
	for x := 0; x < width*2; x += 2 {
		for y := 0; y < height*2; y += 2 {
			position := aocmath.Vertex{X: x, Y: y}
			if _, alreadyVisited := visited[position]; alreadyVisited {
				continue
			}
			regionFences = append(regionFences, generateRegionFenceForPosition(position, plants, visited))
		}
	}

	return regionFences
}

func countPerimeter(position aocmath.Vertex, plant rune, plants map[aocmath.Vertex]rune) int {
	areaDirections := []aocmath.Vertex{{X: -2, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: -2}, {X: 0, Y: 2}}
	neighboors := 0
	for _, dir := range areaDirections {
		if plants[position.Add(dir)] == plant {
			neighboors++
		}
	}
	return 4 - neighboors
}

func generateRegionFenceForPosition(
	position aocmath.Vertex,
	plants map[aocmath.Vertex]rune,
	visited map[aocmath.Vertex]struct{},
) regionFence {
	plant := plants[position]
	areaDirections := []aocmath.Vertex{{X: -2, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: -2}, {X: 0, Y: 2}}
	area := 0
	perimeter := 0
	positions := []aocmath.Vertex{position}
	for {
		nPositions := make([]aocmath.Vertex, 0)
		for _, currentPosition := range positions {
			_, alreadyVisited := visited[currentPosition]
			newPlant, newPlantExists := plants[currentPosition]
			if !newPlantExists || newPlant != plant || alreadyVisited {
				continue
			}
			visited[currentPosition] = struct{}{}
			area++
			perimeter += countPerimeter(currentPosition, plant, plants)
			for _, areaDirection := range areaDirections {
				nPositions = append(nPositions, areaDirection.Add(currentPosition))
			}
		}
		positions = nPositions
		if len(positions) == 0 {
			break
		}
	}
	regionFence := regionFence{name: plant, area: area, perimeter: perimeter}
	return regionFence
}

func (data *Day12Part01) Exec() (string, error) {
	regionFences := calculateRegionFences(data.Plants, data.Width, data.Height)
	var result int64 = 0
	for _, regionFence := range regionFences {
		result += regionFence.calculateRegion()
	}
	return strconv.FormatInt(result, 10), nil
}

func (data *Day12Part02) Exec() (string, error) {
	regionFences := calculateRegionSiteFences(data.Plants, data.Width, data.Height)
	var result int64 = 0
	for _, regionFence := range regionFences {
		result += regionFence.calculateRegion()
	}
	return strconv.FormatInt(result, 10), nil
}

func calculateRegionSiteFences(plants map[aocmath.Vertex]rune, width int, height int) []regionFence {
	regionFences := make([]regionFence, 0)
	visited := make(map[aocmath.Vertex]struct{})

	for x := 0; x < width*2; x += 2 {
		for y := 0; y < height*2; y += 2 {
			currentPosition := aocmath.NewVertex(x, y)
			if _, visited := visited[currentPosition]; visited {
				continue
			}
			regionFenceArea := getRegionFence(currentPosition, plants, visited)
			regionFences = append(regionFences, calculateRegionFence(plants[currentPosition], regionFenceArea))
		}
	}

	return regionFences
}

func calculateRegionFence(plant rune, regionFenceArea map[aocmath.Vertex]struct{}) regionFence {
	edges := make(map[aocmath.Vertex]int)
	detectEdgesDirections := [][]aocmath.Vertex{
		{{X: -2, Y: 0}, {X: -2, Y: -2}, {X: 0, Y: -2}},
		{{X: 0, Y: -2}, {X: 2, Y: -2}, {X: 2, Y: 0}},
		{{X: 2, Y: 0}, {X: 2, Y: 2}, {X: 0, Y: 2}},
		{{X: 0, Y: 2}, {X: -2, Y: 2}, {X: -2, Y: 0}},
	}
	detectedEdgeDirection := []aocmath.Vertex{
		{X: -1, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 1},
		{X: -1, Y: 1},
	}
	for plantPosition := range regionFenceArea {
		// fmt.Println(plantPosition)
		for whichEdge, directions := range detectEdgesDirections {
			edge := ""
			for _, direction := range directions {
				_, exists := regionFenceArea[plantPosition.Add(direction)]
				if exists {
					edge += "1"
				} else {
					edge += "0"
				}
			}
			if edge == "000" || edge == "101" {
				// fmt.Println("Found edge at", whichEdge)
				edges[plantPosition.Add(detectedEdgeDirection[whichEdge])] = 1
			} else if edge == "010" {
				// fmt.Println("Found edge at", whichEdge)
				edges[plantPosition.Add(detectedEdgeDirection[whichEdge])] = 2
			}
		}
	}
	perimeter := 0
	for _, p := range edges {
		perimeter += p
	}
	r := regionFence{ name: plant, area: len(regionFenceArea), perimeter: perimeter}
	// fmt.Println(edges)
	// fmt.Println(r)
	return r
}

func getRegionFence(position aocmath.Vertex, plants map[aocmath.Vertex]rune, visited map[aocmath.Vertex]struct{}) map[aocmath.Vertex]struct{} {
	searchingPlant := plants[position]
	visited[position] = struct{}{}
	directions := []aocmath.Vertex{aocmath.NewVertex(-2, 0), aocmath.NewVertex(2, 0), aocmath.NewVertex(0, -2), aocmath.NewVertex(0, 2)}
	regionFences := make(map[aocmath.Vertex]struct{})
	regionFences[position] = struct{}{}
	positions := []aocmath.Vertex {position}
	for {
	
		nPositions := make([]aocmath.Vertex, 0)
		for _, position := range positions {
			for _, dir := range directions {
				newPosition := dir.Add(position)
				plant := plants[newPosition]
				_, alreadyVisited := visited[newPosition]
				if plant != searchingPlant || alreadyVisited {
					continue
				}
				visited[newPosition] = struct{}{}
				regionFences[newPosition] = struct{}{}
				nPositions = append(nPositions, newPosition)
			}
		}
		if len(nPositions) <= 0 {
			break
		}
		positions = nPositions
	}
	return regionFences
}
