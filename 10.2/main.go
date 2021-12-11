package main

import (
	"fmt"
	"sort"

	"../util"
)

func main() {
	scores := []int{}

	for _, line := range util.ReadInputLines("./input.txt") {
		stack := []rune{}

		isCorrupt := false
		for _, r := range line {
			if isOpen(r) {
				stack = append(stack, r)
				continue
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if getClose(top) != r {
				isCorrupt = true
				break
			}
		}

		if len(stack) == 0 || isCorrupt {
			continue
		}

		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + getScore(stack[i])
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}

func isOpen(r rune) bool {
	switch r {
	case '[', '{', '(', '<':
		return true
	default:
		return false
	}
}

func getClose(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		panic("unexpected open symbol")
	}
}

func getScore(r rune) int {
	switch r {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
		panic("unexpected open symbol")
	}
}
