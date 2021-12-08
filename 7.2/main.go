package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

func main() {
	line := util.ReadInputLines("./input.txt")[0]

	positions := []int{}
	sumPosition := 0

	for _, positionStr := range strings.Split(line, ",") {
		position, err := strconv.Atoi(positionStr)
		util.Check(err)
		positions = append(positions, position)
		sumPosition += position
	}

	avgPosition := int(float64(sumPosition) / float64(len(positions)))

	// Using the floor'ed avg
	totalDistance1 := 0
	for _, position := range positions {
		steps := util.AbsInt(position - avgPosition)

		totalDistance1 += (steps * (steps + 1)) / 2
	}

	// Using the ceil'ed avg
	totalDistance2 := 0
	for _, position := range positions {
		steps := util.AbsInt(position - (avgPosition + 1))

		totalDistance2 += (steps * (steps + 1)) / 2
	}

	fmt.Println(util.MinInt(totalDistance1, totalDistance2))
}
