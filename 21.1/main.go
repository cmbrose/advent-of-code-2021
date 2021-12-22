package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

func main() {
	lines := util.ReadInputLines("./input.txt")

	dieVal := 1
	rollNumber := 0

	roll := func() int {
		rollNumber += 3
		val := dieVal*3 + 3
		dieVal = (dieVal+2)%100 + 1
		return val
	}

	player := 0
	positions := []int{getStartPos(lines[0]), getStartPos(lines[1])}
	scores := []int{0, 0}

	takeTurn := func() bool {
		rollVal := roll()
		positions[player] = (positions[player]+rollVal-1)%10 + 1
		scores[player] += positions[player]

		if scores[player] >= 1000 {
			return false
		}

		player = 1 - player
		return true
	}

	for takeTurn() {
	}

	loser := 1 - player

	loserScore := scores[loser]
	fmt.Println(loserScore, " * ", rollNumber, " => ", loserScore*rollNumber)
}

func getStartPos(line string) int {
	parts := strings.Split(line, " ")
	startStr := parts[len(parts)-1]

	start, err := strconv.Atoi(startStr)
	util.Check(err)

	return start
}
