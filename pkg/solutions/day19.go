package solutions

import (
	"container/heap"
	"fmt"
	"strconv"

	aocmath "github.com/jkondarewicz/aoc2024/pkg/math"
	"github.com/jkondarewicz/aoc2024/pkg/utils"
)

type Day19Part01 struct {
	PossibleTowelPatterns []string
	DesiredDesigns        []string
}

type Day19Part02 struct {
	PossibleTowelPatterns []string
	DesiredDesigns        []string
}

func (data *Day19Part01) Exec() (string, error) {
	designMatcher := &designMatcher{
		towelPatterns: utils.NewSet[string](),
		patternSizes:  utils.NewSet[int](),
	}
	for _, pattern := range data.PossibleTowelPatterns {
		designMatcher.towelPatterns.Add(pattern)
		designMatcher.patternSizes.Add(len(pattern))
	}
	result := 0
	for _, design := range data.DesiredDesigns {
		if matching := designMatcher.isDesignMatching(design); matching {
			result++
		}
	}
	return strconv.Itoa(result), nil
}

func (data *Day19Part02) Exec() (string, error) {
	return "", nil
}

type designMatcher struct {
	towelPatterns *utils.Set[string]
	patternSizes  *utils.Set[int]
}

func (d *designMatcher) isDesignMatching(design string) bool {
	m := &matcher{currentIndex: 0, reachedBy: 1}

	mq := &mQueue{}
	heap.Init(mq)
	heap.Push(mq, m)
	dl := len(design)

	beenHere := make(map[aocmath.Vertex]int)
	reachedEnd := 0

	for mq.Len() > 0 {
		current := heap.Pop(mq).(*matcher)
		for _, offset := range d.patternSizes.Get() {
			end := current.currentIndex + offset
			_, visited := beenHere[aocmath.NewVertex(current.currentIndex, offset)]
			if end > dl || visited {
				continue
			}
			beenHere[aocmath.NewVertex(current.currentIndex, offset)]++
			searchingPattern := design[current.currentIndex:end]
			if d.towelPatterns.Exists(searchingPattern) {
				if end == dl {
					reachedEnd++
				}
				heap.Push(mq, &matcher{currentIndex: end})
			}
		}
	}
	fmt.Println("Reached end", reachedEnd)
	return reachedEnd > 0
}

type matcher struct {
	currentIndex int
	reachedBy    int
}

type mQueue []*matcher

func (pq mQueue) Less(i, j int) bool {
	a := pq[i]
	b := pq[j]
	return a.currentIndex > b.currentIndex
}
func (pq mQueue) Len() int { return len(pq) }
func (pq mQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *mQueue) Push(x any) {
	*pq = append(*pq, x.(*matcher))
}

func (pq *mQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
