package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

// P1 position
// P1 score
// P2 position
// P2 score
// Wins counts
type ResultCache [][][][][][]int64

func makeCache() ResultCache {
	cache := make([][][][][][]int64, 10)

	for pos1 := 0; pos1 < 10; pos1++ {
		cache[pos1] = make([][][][][]int64, 21)

		for score1 := 0; score1 < 21; score1++ {
			cache[pos1][score1] = make([][][][]int64, 10)

			for pos2 := 0; pos2 < 10; pos2++ {
				cache[pos1][score1][pos2] = make([][][]int64, 21)

				for score2 := 0; score2 < 21; score2++ {
					cache[pos1][score1][pos2][score2] = make([][]int64, 6)
				}
			}
		}
	}

	return cache
}

func (cache ResultCache) getCacheValue(pos1, score1, pos2, score2, roll int) []int64 {
	return cache[pos1-1][score1][pos2-1][score2][roll]
}

func (cache ResultCache) setCacheValue(pos1, score1, pos2, score2, roll int, scores []int64) {
	cache[pos1-1][score1][pos2-1][score2][roll] = scores
}

func main() {
	lines := util.ReadInputLines("./input.txt")

	cache := makeCache()

	pos1 := getStartPos(lines[0])
	pos2 := getStartPos(lines[1])

	p1, p2 := getWins(pos1, 0, pos2, 0, 0, cache)

	fmt.Println(p1, " to ", p2)
}

func getStartPos(line string) int {
	parts := strings.Split(line, " ")
	startStr := parts[len(parts)-1]

	start, err := strconv.Atoi(startStr)
	util.Check(err)

	return start
}

func getWins(pos1, score1, pos2, score2, rollNumber int, cache ResultCache) (int64, int64) {
	if score1 >= 21 {
		return 1, 0
	}
	if score2 >= 21 {
		return 0, 1
	}

	if cached := cache.getCacheValue(pos1, score1, pos2, score2, rollNumber); cached != nil {
		return cached[0], cached[1]
	}

	var p1Wins int64 = 0
	var p2Wins int64 = 0
	for roll := 1; roll <= 3; roll++ {
		newPos1 := pos1
		newScore1 := score1
		newPos2 := pos2
		newScore2 := score2

		if rollNumber < 3 {
			newPos1 = getNextPosition(pos1, roll)

			if rollNumber == 2 {
				newScore1 = score1 + newPos1
			}
		} else {
			newPos2 = getNextPosition(pos2, roll)

			if rollNumber == 5 {
				newScore2 = score2 + newPos2
			}
		}

		addtP1Wins, addtP2Wins := getWins(newPos1, newScore1, newPos2, newScore2, (rollNumber+1)%6, cache)
		p1Wins += addtP1Wins
		p2Wins += addtP2Wins
	}

	cache.setCacheValue(pos1, score1, pos2, score2, rollNumber, []int64{p1Wins, p2Wins})

	return p1Wins, p2Wins
}

func getNextPosition(pos, movement int) int {
	return (pos+movement-1)%10 + 1
}
