package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"../util"
)

func main() {
	cnt := 0
	prevValue := math.MaxInt32

	for _, line := range util.ReadInputLines("./input.txt") {
		value, err := strconv.Atoi(strings.TrimSpace(line))
		util.Check(err)

		if prevValue < value {
			cnt++
		}

		prevValue = value
	}

	fmt.Printf("%d", cnt)
}
