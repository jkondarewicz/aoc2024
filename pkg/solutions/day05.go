package solutions

import (
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type node struct {
	page int
	nextNodes []*node
	visited bool
}
func (n *node) childVisited() bool {
	for _, node := range n.nextNodes {
		if node.visited {
			return true
		} else if len(node.nextNodes) > 0 {
			for _, node = range node.nextNodes {
				if node.childVisited() {
					return true
				}
			}
		}
	}
	return false
}
func createNode(page int) *node {
	return &node { page: page, nextNodes: make([]*node, 0), visited: false}
}

type nodeHolder struct {
	nodes map[int]*node
}
func (nodeHolder *nodeHolder) reset() {
	for _, value := range nodeHolder.nodes {
		value.visited = false
	}
}
func (nodeHolder *nodeHolder) addNode(before int, after int) *node {
	beforeNode := nodeHolder.nodes[before]
	afterNode := nodeHolder.nodes[after]
	if beforeNode == nil {
		beforeNode = createNode(before)
	}
	if afterNode == nil {
		afterNode = createNode(after)
	}
	beforeNode.nextNodes = append(beforeNode.nextNodes, afterNode)
	nodeHolder.nodes[before] = beforeNode
	nodeHolder.nodes[after] = afterNode
	return beforeNode
}

type Day05Part01 struct {
	Rules []aocmath.Vertex
	Pages [][]int
}
func (nodeHolder *nodeHolder) isCorrect(pages []int) bool {
	for _, page := range pages {
		node := nodeHolder.nodes[page]
		if node == nil {
			continue
		}
		node.visited = true
		if node.childVisited() {
			return false
		}
	}
	return true
}
type Day05Part02 struct {
	Chars [][]rune
}

func generateNodeHolder(rules []aocmath.Vertex) nodeHolder {
	nodeHolder := nodeHolder { nodes: make(map[int]*node)}
	for _, rule := range rules {
		nodeHolder.addNode(rule.X, rule.Y)
	}
	return nodeHolder
}

func (data *Day05Part01) Exec() (string, error) {
	result := 0
	nodeHolder := generateNodeHolder(data.Rules)
	for _, pages := range data.Pages {
		if nodeHolder.isCorrect(pages) {
			result += pages[len(pages) / 2]
		}
		nodeHolder.reset()
	}
	return strconv.Itoa(result), nil
}

func (data *Day05Part02) Exec() (string, error) {
	return "", nil
}
