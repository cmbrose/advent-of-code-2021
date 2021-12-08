package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../util"
)

func main() {
	line := util.ReadInputLines("./input.txt")[0]

	positions := []int{}

	for _, positionStr := range strings.Split(line, ",") {
		position, err := strconv.Atoi(positionStr)
		util.Check(err)
		positions = append(positions, position)
	}

	sort.Ints(positions)

	var median int
	if len(positions)%2 == 0 {
		median = (positions[len(positions)/2] + positions[len(positions)/2-1]) / 2
	} else {
		median = positions[len(positions)/2]
	}

	totalDistance := 0
	for _, position := range positions {
		totalDistance += util.AbsInt(position - median)
	}

	fmt.Println(totalDistance)
}
