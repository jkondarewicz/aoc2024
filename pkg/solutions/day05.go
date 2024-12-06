package solutions

import (
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type pageRule struct {
	page         int
	shouldBeNext map[int]bool
}

func makePageRules(rules []aocmath.Vertex) map[int]*pageRule {
	pageRules := make(map[int]*pageRule)
	for _, page := range rules {
		if pageRules[page.X] == nil {
			pageRules[page.X] = &pageRule{page: page.X, shouldBeNext: make(map[int]bool)}
		}
		pageRules[page.X].shouldBeNext[page.Y] = true
	}
	return pageRules
}

type Day05Part01 struct {
	Rules []aocmath.Vertex
	Pages [][]int
}
type Day05Part02 struct {
	Rules []aocmath.Vertex
	Pages [][]int
}

func (data *Day05Part01) Exec() (string, error) {
	pageRules := makePageRules(data.Rules)
	result := 0
	for _, pages := range data.Pages {
		visited := make(map[int]bool)
		correct := true
		for _, page := range pages {
			visited[page] = true
			pageRule := pageRules[page]
			if pageRule == nil {
				continue
			}
			for shouldBeBefore := range pageRule.shouldBeNext {
				if visited[shouldBeBefore] {
					correct = false
					break
				}
			}
		}
		if correct {
			result += pages[len(pages)/2]
		}
	}
	return strconv.Itoa(result), nil
}

func (data *Day05Part02) Exec() (string, error) {
	pageRules := makePageRules(data.Rules)
	result := 0
	for _, pages := range data.Pages {
		correctedData, corrected := correctData(pages, pageRules, false)
		if corrected {
			result += correctedData[len(correctedData)/2]
		}
	}
	return strconv.Itoa(result), nil
}

type visitedNumber struct {
	index   int
	visited bool
}

func moveElement(slice []int, fromIndex int, toIndex int) []int {
	if fromIndex < 0 || fromIndex >= len(slice) || toIndex < 0 || toIndex >= len(slice) {
		return slice
	}

	// Get the element to move
	element := slice[fromIndex]

	// Remove the element from the original position
	slice = append(slice[:fromIndex], slice[fromIndex+1:]...)

	// Insert the element at the new position
	slice = append(slice[:toIndex], append([]int{element}, slice[toIndex:]...)...)

	return slice
}

func correctData(pages []int, pageRules map[int]*pageRule, corrected bool) ([]int, bool) {
	visited := make(map[int]visitedNumber)
	newPages := make([]int, 0)
	for index, page := range pages {
		visited[page] = visitedNumber{index: index, visited: true}
		pageRule := pageRules[page]
		if pageRule == nil {
			continue
		}
		for shouldBeBefore := range pageRule.shouldBeNext {
			if visited[shouldBeBefore].visited {
				newPages = moveElement(pages, index, visited[shouldBeBefore].index)
				break
			}
		}
		if len(newPages) > 0 {
			break
		}
	}
	if len(newPages) > 0 {
		return correctData(newPages, pageRules, true)
	}
	return pages, corrected
}
