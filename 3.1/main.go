package main

import (
	"fmt"

	"../util"
)

func main() {
	bits := 12

	setBitCounts := make([]int, bits)

	lineCnt := 0
	for _, line := range util.ReadInputLines("./input.txt") {
		for i, c := range line {
			if i >= bits {
				panic("More bits than expected")
			}

			if c == '1' {
				setBitCounts[i] += 1
			}
		}

		lineCnt++
	}

	gamma := 0
	epsilon := 0

	for i := 0; i < bits; i++ {
		gamma <<= 1
		epsilon <<= 1

		// More than half the bits are set
		if setBitCounts[i]*2 > lineCnt {
			gamma++
		} else {
			epsilon++
		}
	}

	fmt.Printf("%d * %d = %d\n", gamma, epsilon, gamma*epsilon)
}
