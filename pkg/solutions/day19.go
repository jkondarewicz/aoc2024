package solutions

import (
	"container/heap"
	"strconv"

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
		if matching, _ := designMatcher.isDesignMatching(design); matching {
			result++
		}
	}
	return strconv.Itoa(result), nil
}

func (data *Day19Part02) Exec() (string, error) {
	designMatcher := &designMatcher{
		towelPatterns: utils.NewSet[string](),
		patternSizes:  utils.NewSet[int](),
	}
	for _, pattern := range data.PossibleTowelPatterns {
		designMatcher.towelPatterns.Add(pattern)
		designMatcher.patternSizes.Add(len(pattern))
	}
	hm := 0
	for _, design := range data.DesiredDesigns {
		if matching, man := designMatcher.isDesignMatching(design); matching {
			hm += man
		}
	}
	return strconv.Itoa(hm), nil
}

type designMatcher struct {
	towelPatterns *utils.Set[string]
	patternSizes  *utils.Set[int]
}

func (d *designMatcher) isDesignMatching(design string) (bool, int) {
	m := &matcher{currentIndex: 0, visited: utils.NewSet[int](),}

	mq := &mQueue{}
	heap.Init(mq)
	heap.Push(mq, m)
	dl := len(design)

	visited := utils.NewSet[int]()
	all := make([]*matcher, 0)
	found := make([]*matcher, 0)

	for mq.Len() > 0 {
		current := heap.Pop(mq).(*matcher)
		if visited.Exists(current.currentIndex) {
			for _, c := range all {
				if c.visited.Exists(current.currentIndex) {
					c.attached += 1 + current.attached
				}
			}
			continue
		}
		visited.Add(current.currentIndex)
		current.visited.Add(current.currentIndex)
		for _, offset := range d.patternSizes.Get() {
			end := current.currentIndex + offset
			if end > dl {
				continue
			}
			searchingPattern := design[current.currentIndex:end]
			if d.towelPatterns.Exists(searchingPattern) {
				if end == dl { //reached end
					found = append(found, current)
					continue
				}
				next := &matcher{currentIndex: end, visited: current.visited.Copy(), attached: current.attached}
				all = append(all, next)
				heap.Push(mq, next)
			}
		}
	}
	hm := len(found)
	for _, f := range found {
		hm += f.attached
	}
	return len(found) > 0, hm
}

type matcher struct {
	currentIndex int
	visited      *utils.Set[int]
	attached     int
}

type mQueue []*matcher

func (pq mQueue) Less(i, j int) bool {
	a := pq[i]
	b := pq[j]
	return a.currentIndex < b.currentIndex
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
