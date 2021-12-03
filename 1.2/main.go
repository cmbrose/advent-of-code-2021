package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

func main() {
	windowSize := 3

	ring := make([]int, windowSize)

	cnt := 0
	for i, line := range util.ReadInputLines("./input.txt") {
		value, err := strconv.Atoi(strings.TrimSpace(line))
		util.Check(err)

		oldValue := ring[i%windowSize]

		ring[i%windowSize] = value

		if i < windowSize {
			continue
		}

		if oldValue < value {
			cnt++
		}
	}

	fmt.Printf("%d\n", cnt)
}
