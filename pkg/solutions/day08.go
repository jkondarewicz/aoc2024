package solutions

import (
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
)

type Day08Part01 struct {
	Width    int
	Height   int
	Antennas map[rune][]aocmath.Vertex
}
func isPointInArea(point aocmath.Vertex, width int, height int) bool {
	return point.X >= 0 && point.Y >= 0 && point.X < width && point.Y < height
}

func calculateNextAntinode(
	currentPosition aocmath.Vertex,
	translation aocmath.Vertex,
	width int,
	height int,
) (aocmath.Vertex, bool) {
	nextAntinode := currentPosition.Add(translation)
	if isPointInArea(nextAntinode, width, height) {
		return nextAntinode, true
	} else {
		return nextAntinode, false
	}
}
func findAntinodes(antennas map[rune][]aocmath.Vertex, width int, height int, secondPart bool) map[aocmath.Vertex]bool {
	antinodes := make(map[aocmath.Vertex]bool)
	for _, antennas := range antennas {
		for firstIndex := 0; firstIndex < len(antennas) - 1; firstIndex++ {
			for secondIndex := firstIndex + 1; secondIndex < len(antennas); secondIndex++ {
				firstAntenna := antennas[firstIndex]
				secondAntenna := antennas[secondIndex]
				secondAntennaAntinodeTranslation := firstAntenna.DiffBetweenVertexes(secondAntenna)
				firstAntennaAntinodeTranslation := secondAntennaAntinodeTranslation.Opposite()
				if !secondPart {
					firstAntinode, inArea := calculateNextAntinode(firstAntenna, firstAntennaAntinodeTranslation, width, height)
					if inArea {
						antinodes[firstAntinode] = true
					}
					secondAntinode, inArea := calculateNextAntinode(secondAntenna, secondAntennaAntinodeTranslation, width, height)
					if inArea {
						antinodes[secondAntinode] = true
					}
					continue
				}
				antinodes[firstAntenna] = true
				antinodes[secondAntenna] = true
				var antinode = firstAntenna
				var inArea bool
				for {
					antinode, inArea = calculateNextAntinode(antinode, firstAntennaAntinodeTranslation, width, height)
					if inArea {
						antinodes[antinode] = true
					} else {
						break
					}
				}
				antinode = secondAntenna
				for {
					antinode, inArea = calculateNextAntinode(antinode, secondAntennaAntinodeTranslation, width, height)
					if inArea {
						antinodes[antinode] = true
					} else {
						break
					}
				}

			}
		}
	}
	return antinodes
}


type Day08Part02 struct {
	Width    int
	Height   int
	Antennas map[rune][]aocmath.Vertex
}

func (data *Day08Part01) Exec() (string, error) {
	antinodes := findAntinodes(data.Antennas, data.Width, data.Height, false)
	return strconv.Itoa(len(antinodes)), nil
}

func (data *Day08Part02) Exec() (string, error) {
	antinodes := findAntinodes(data.Antennas, data.Width, data.Height, true)
	return strconv.Itoa(len(antinodes)), nil
}
