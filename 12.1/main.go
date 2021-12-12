package main

import (
	"fmt"
	"strings"
	"unicode"

	"../util"
)

func main() {
	nodes := make(map[string]*Node)

	for _, line := range util.ReadInputLines("./input.txt") {
		pair := strings.Split(line, "-")

		n1, ok := nodes[pair[0]]
		if !ok {
			n1 = newNode(pair[0])
			nodes[pair[0]] = n1
		}

		n2, ok := nodes[pair[1]]
		if !ok {
			n2 = newNode(pair[1])
			nodes[pair[1]] = n2
		}

		n1.AddEdge(n2)
		n2.AddEdge(n1)
	}

	start := nodes["start"]
	end := nodes["end"]

	paths := start.CountPaths(end)

	fmt.Println(paths)
}

type Node struct {
	Name     string
	Edges    []*Node
	IsMarked bool
	IsLarge  bool
}

func newNode(name string) *Node {
	isLarge := unicode.IsUpper(rune(name[0]))

	return &Node{
		Name:     name,
		Edges:    []*Node{},
		IsLarge:  isLarge,
		IsMarked: false,
	}
}

func (n *Node) String() string {
	return n.Name
}

func (n *Node) AddEdge(other *Node) {
	n.Edges = append(n.Edges, other)
}

func (n *Node) CountPaths(end *Node) int {
	if n == end {
		return 1
	}

	if n.IsMarked {
		return 0
	}

	if !n.IsLarge {
		n.IsMarked = true
		defer func() { n.IsMarked = false }()
	}

	sum := 0
	for _, adj := range n.Edges {
		sum += adj.CountPaths(end)
	}

	return sum
}
