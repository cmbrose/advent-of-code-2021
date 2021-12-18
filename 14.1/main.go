package main

import (
	"fmt"
	"math"
	"strings"

	"../util"
)

type Data struct {
	Counts []int
	Next   byte
}

func main() {
	lines := util.ReadInputLines("./input.txt")

	start := lines[0]

	pairs := make(map[string]*Data)

	for _, line := range lines[2:] {
		rule := strings.Split(line, " -> ")

		pair := rule[0]
		next := rule[1]

		pairs[pair] = &Data{Counts: []int{0, 0}, Next: next[0]}
	}

	prev := start[0]
	for _, curr := range start[1:] {
		pair := fmt.Sprintf("%c%c", prev, curr)

		pairs[pair].Counts[0] += 1

		prev = byte(curr)
	}

	on := 0
	off := 1

	for i := 0; i < 10; i++ {
		// printState(pairs, on)

		for _, data := range pairs {
			data.Counts[off] = 0
		}

		for pair, data := range pairs {
			pair1 := fmt.Sprintf("%c%c", pair[0], data.Next)

			pairs[pair1].Counts[off] += data.Counts[on]

			pair2 := fmt.Sprintf("%c%c", data.Next, pair[1])

			pairs[pair2].Counts[off] += data.Counts[on]
		}

		on = 1 - on
		off = 1 - off
	}

	// printState(pairs, on)

	counts := make(map[byte]int)

	for pair, data := range pairs {
		cnt, ok := counts[pair[0]]
		if !ok {
			counts[pair[0]] = data.Counts[on]
		} else {
			counts[pair[0]] = cnt + data.Counts[on]
		}

		cnt, ok = counts[pair[1]]
		if !ok {
			counts[pair[1]] = data.Counts[on]
		} else {
			counts[pair[1]] = cnt + data.Counts[on]
		}
	}

	minCnt := math.MaxInt
	var minChar byte = 0

	maxCnt := 0
	var maxChar byte = 0

	for char, cnt := range counts {
		// Chars get double counted as the same char appears in 2 pairs
		cnt = cnt / 2

		// ... except for the first or last char, those don't get double counted one time
		if char == start[0] || char == start[len(start)-1] {
			cnt += 1
		}

		if minCnt > cnt {
			minCnt = cnt
			minChar = char
		}

		if maxCnt < cnt {
			maxCnt = cnt
			maxChar = char
		}
	}

	fmt.Printf("%c = %d, %c = %d => %d\n", maxChar, maxCnt, minChar, minCnt, maxCnt-minCnt)
}

func printState(pairs map[string]*Data, on int) {
	str := "> "
	for pair, data := range pairs {
		str += fmt.Sprintf("%s = %d, ", pair, data.Counts[on])
	}

	fmt.Println(str)
}
