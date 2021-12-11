package main

import (
	"fmt"

	"../util"
)

func main() {
	sum := 0
	for _, line := range util.ReadInputLines("./input.txt") {
		stack := []rune{}

		for _, r := range line {
			if isOpen(r) {
				stack = append(stack, r)
				continue
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if isMatchingClose(top, r) {
				continue
			}

			sum += scoreClose(r)
			break
		}
	}

	fmt.Println(sum)
}

func isOpen(r rune) bool {
	switch r {
	case '[', '{', '(', '<':
		return true
	default:
		return false
	}
}

func isMatchingClose(open, r rune) bool {
	switch r {
	case ')':
		return open == '('
	case ']':
		return open == '['
	case '}':
		return open == '{'
	case '>':
		return open == '<'
	default:
		panic("unexpected close symbol")
	}
}

func scoreClose(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic("unexpected close symbol")
	}
}
