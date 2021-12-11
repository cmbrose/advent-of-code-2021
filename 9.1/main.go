package main

import (
	"fmt"

	"../util"
)

func main() {
	sum := 0

	vals := parseMap()

	for y := 0; y < len(vals); y++ {
		for x := 0; x < len(vals[y]); x++ {
			if isLowPoint(x, y, vals) {
				sum += vals[y][x] + 1
			}
		}
	}

	fmt.Println(sum)
}

func parseMap() [][]int {
	mapVals := [][]int{}

	for _, line := range util.ReadInputLines("./input.txt") {
		row := []int{}
		for _, cell := range line {
			row = append(row, int(cell-'0'))
		}

		mapVals = append(mapVals, row)
	}

	return mapVals
}

func isLowPoint(x, y int, vals [][]int) bool {
	pointVal := vals[y][x]

	for i := -1; i <= 1; i++ {
		if y+i < 0 || y+i >= len(vals) {
			continue
		}

		row := vals[y+i]

		for j := -1; j <= 1; j++ {
			if x+j < 0 || x+j >= len(row) {
				continue
			}

			if i == 0 && j == 0 {
				continue
			}

			if pointVal >= row[x+j] {
				return false
			}
		}
	}

	return true
}
