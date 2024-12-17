package solutions

import (
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
)

var directions []aocmath.Vertex = []aocmath.Vertex{
	{X: 1, Y: 0},   //right
	{X: -1, Y: 0},  //left
	{X: 0, Y: 1},   //down
	{X: 0, Y: -1},  //up
	{X: 1, Y: -1},  //up-right
	{X: 1, Y: 1},   //down-right
	{X: -1, Y: -1}, //up-left
	{X: -1, Y: 1},  //down-left
}
var xmas map[rune]int = map[rune]int{'X': 0, 'M': 1, 'A': 2, 'S': 3}
var mas map[rune]int = map[rune]int{'M': 1, 'S': -1, 'X': 0, 'A': 0}

func isXmasWord(chars [][]rune, startingPoint aocmath.Vertex, direction aocmath.Vertex) bool {
	currentCharOfWordId := -1
	currentPoint := startingPoint
	for {
		if !(currentPoint.Y < len(chars) && currentPoint.Y >= 0 &&
			currentPoint.X < len(chars[currentPoint.Y]) && currentPoint.X >= 0) {
			//out of bound means return false
			return false
		} 
		nextCharOfWord := chars[currentPoint.Y][currentPoint.X]
		nextCharOfWordId := xmas[nextCharOfWord]
		isNextCharContinueWord := (nextCharOfWordId - currentCharOfWordId) == 1
		if isNextCharContinueWord {
			currentCharOfWordId = nextCharOfWordId
			if currentCharOfWordId == xmas['S'] {
				return true
			}
		} else {
			return false
		}
		currentPoint = currentPoint.Add(direction)
	}
}

func isCrossedWord(chars [][]rune, startingPoint aocmath.Vertex) bool {
	if chars[startingPoint.Y][startingPoint.X] != 'A' {
		return false
	}
	if startingPoint.Y > 0 && startingPoint.Y < len(chars)-1 &&
		startingPoint.X > 0 && startingPoint.X < len(chars[startingPoint.Y])-1 {
		f1 := chars[startingPoint.Y-1][startingPoint.X-1]
		f2 := chars[startingPoint.Y+1][startingPoint.X+1]
		s1 := chars[startingPoint.Y-1][startingPoint.X+1]
		s2 := chars[startingPoint.Y+1][startingPoint.X-1]
		return mas[f1] * mas[f2] == -1 && mas[s1] * mas[s2] == -1
	} else {
		return false
	}
}

type Day04Part01 struct {
	Chars [][]rune
}

type Day04Part02 struct {
	Chars [][]rune
}

func (data *Day04Part01) Exec() (string, error) {
	result := 0
	for y, row := range data.Chars {
		for x := range row {
			for _, direction := range directions {
				if isXmasWord(data.Chars, aocmath.Vertex{X: x, Y: y,}, direction) {
					result++
				}
			}
		}
	}
	return strconv.Itoa(result), nil
}

func (data *Day04Part02) Exec() (string, error) {
	result := 0
	for y, row := range data.Chars {
		for x := range row {
			if isCrossedWord(data.Chars, aocmath.Vertex{X: x, Y: y,}) {
				result++
			}
		}
	}
	return strconv.Itoa(result), nil
}
