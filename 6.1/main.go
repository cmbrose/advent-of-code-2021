package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

func main() {
	line := util.ReadInputLines("./input.txt")[0]

	cntByAge := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	simulationDays := 80

	for _, ageStr := range strings.Split(line, ",") {
		age, err := strconv.Atoi(ageStr)
		util.Check(err)
		cntByAge[age]++
	}

	for i := 0; i < simulationDays; i++ {
		spawnedCnt := cntByAge[0]

		for j := 0; j < len(cntByAge)-1; j++ {
			cntByAge[j] = cntByAge[j+1]
		}

		cntByAge[6] += spawnedCnt
		cntByAge[8] = spawnedCnt
	}

	sum := 0
	for _, cnt := range cntByAge {
		sum += cnt
	}

	fmt.Println(sum)
}
