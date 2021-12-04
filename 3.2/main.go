package main

import (
	"fmt"

	"../util"
)

func main() {
	root := &Node{Weight: 0}

	for _, line := range util.ReadInputLines("./input.txt") {
		root.add(line)
	}

	oxygenS := root.buildPath(
		func(n0, n1 *Node) *Node {
			if n1.Weight >= n0.Weight {
				return n1
			}
			return n0
		})

	co2S := root.buildPath(
		func(n0, n1 *Node) *Node {
			if n1.Weight >= n0.Weight {
				return n0
			}
			return n1
		})

	oxygen := util.ParseBitString(oxygenS)
	co2 := util.ParseBitString(co2S)

	fmt.Printf("%s * %s = %d * %d = %d\n", oxygenS, co2S, oxygen, co2, oxygen*co2)
}

type Node struct {
	Weight int
	Value  rune
	Zero   *Node
	One    *Node
}

func (n *Node) add(str string) {
	n.Weight++

	if len(str) == 0 {
		return
	}

	isOne := str[0] == '1'

	if isOne {
		if n.One == nil {
			n.One = &Node{Weight: 0, Value: '1'}
		}

		n.One.add(str[1:])
	} else {
		if n.Zero == nil {
			n.Zero = &Node{Weight: 0, Value: '0'}
		}

		n.Zero.add(str[1:])
	}
}

func (n *Node) buildPath(selector func(n0, n1 *Node) *Node) string {
	var selected *Node

	if n.Zero != nil && n.One != nil {
		selected = selector(n.Zero, n.One)
	} else if n.Zero != nil {
		selected = n.Zero
	} else if n.One != nil {
		selected = n.One
	} else {
		return fmt.Sprintf("%c", n.Value)
	}

	path := selected.buildPath(selector)
	return fmt.Sprintf("%c%s", n.Value, path)
}
