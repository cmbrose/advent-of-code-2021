package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	"../util"
)

/*
#############
#01.2.3.4.56#
###A#B#C#D###
  #A#B#C#D#
  #########
*/

type Board struct {
	Hallway []byte
	Rooms   [][]byte
}

func parseBoard() Board {
	board := Board{
		Hallway: []byte{'.', '.', '.', '.', '.', '.', '.'},
		Rooms: [][]byte{
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
		},
	}

	lines := util.ReadInputLines("./input.txt")

	level3 := strings.Split(lines[2], "#")
	level2 := strings.Split(lines[3], "#")
	level1 := strings.Split(lines[4], "#")
	level0 := strings.Split(lines[5], "#")

	for i := 0; i < 4; i++ {
		board.Rooms[i][3] = level3[i+3][0]
		board.Rooms[i][2] = level2[i+1][0]
		board.Rooms[i][1] = level1[i+1][0]
		board.Rooms[i][0] = level0[i+1][0]
	}

	return board
}

func printBoard(board Board) {
	var out bytes.Buffer

	out.WriteString("#############\n")

	out.WriteString(fmt.Sprintf("#%c%c.%c.%c.%c.%c%c#\n",
		board.Hallway[0], board.Hallway[1], board.Hallway[2],
		board.Hallway[3], board.Hallway[4], board.Hallway[5],
		board.Hallway[6]))

	out.WriteString(fmt.Sprintf("###%c#%c#%c#%c###\n",
		board.Rooms[0][3], board.Rooms[1][3], board.Rooms[2][3], board.Rooms[3][3]))

	out.WriteString(fmt.Sprintf("  #%c#%c#%c#%c#  \n",
		board.Rooms[0][2], board.Rooms[1][2], board.Rooms[2][2], board.Rooms[3][2]))

	out.WriteString(fmt.Sprintf("  #%c#%c#%c#%c#  \n",
		board.Rooms[0][1], board.Rooms[1][1], board.Rooms[2][1], board.Rooms[3][1]))

	out.WriteString(fmt.Sprintf("  #%c#%c#%c#%c#  \n",
		board.Rooms[0][0], board.Rooms[1][0], board.Rooms[2][0], board.Rooms[3][0]))

	out.WriteString("  #########  \n")

	fmt.Print(out.String())
}

func main() {
	board := parseBoard()

	printBoard(board)

	cost := findMinRoute(board, 0, math.MaxInt64)

	fmt.Println(cost)
}

func stepPause(board Board) {
	enabled := false

	if enabled {
		printBoard(board)
		fmt.Println("Press a key to continue")
		fmt.Scanln()
	}
}

func findMinRoute(board Board, currCost, minKnownCost int64) int64 {
	if currCost >= minKnownCost {
		return -1
	}

	stepPause(board)

	completeCnt := 0

	// Try moving into the hallway
	for i := 0; i < 4; i++ {
		if isRoomComplete(board, i) {
			completeCnt++
			continue
		}

		roomPos := topFilledRoomPosition(board, i)
		if roomPos == -1 {
			// Room is empty
			continue
		}

		kind := board.Rooms[i][roomPos]
		targetRoom := kind - 'A'

		if i == int(targetRoom) && isRoomCompleteAtLevel(board, i, roomPos) {
			// Already in place, don't need to check for roomPos = 1
			// because it would have been caught by isRoomComplete
			continue
		}

		hallwayMoves := getValidHallwayPositionsToMoveTo(board, i)

		for _, hallwayPos := range hallwayMoves {
			cost := getMoveCost(kind, i, roomPos, hallwayPos)

			// Make move
			board.Hallway[hallwayPos] = kind
			board.Rooms[i][roomPos] = '.'

			newCost := findMinRoute(board, currCost+cost, minKnownCost)

			// Reset move
			board.Hallway[hallwayPos] = '.'
			board.Rooms[i][roomPos] = kind

			if newCost == -1 {
				continue
			}

			if minKnownCost > newCost {
				minKnownCost = newCost
			}
		}
	}

	if completeCnt == 4 {
		// Done
		return currCost
	}

	// Try moving back to a room
	for hallwayPos, kind := range board.Hallway {
		if kind == '.' {
			continue
		}

		if !canMoveBackToRoom(board, hallwayPos) {
			continue
		}

		targetRoom := int(kind - 'A')

		roomPos := topFilledRoomPosition(board, targetRoom) + 1

		cost := getMoveCost(kind, targetRoom, roomPos, hallwayPos)

		// Make move
		board.Hallway[hallwayPos] = '.'
		board.Rooms[targetRoom][roomPos] = kind

		newCost := findMinRoute(board, currCost+cost, minKnownCost)

		// Reset move
		board.Hallway[hallwayPos] = kind
		board.Rooms[targetRoom][roomPos] = '.'

		if newCost == -1 {
			continue
		}

		if minKnownCost > newCost {
			minKnownCost = newCost
		}
	}

	return minKnownCost
}

func topFilledRoomPosition(board Board, roomNumber int) int {
	for i := 3; i >= 0; i-- {
		if board.Rooms[roomNumber][i] != '.' {
			return i
		}
	}

	return -1
}

func isRoomCompleteAtLevel(board Board, roomNumber, level int) bool {
	targetKind := byte('A' + roomNumber)

	for i := level; i >= 0; i-- {
		if board.Rooms[roomNumber][i] != targetKind {
			return false
		}
	}

	return true
}

func isRoomComplete(board Board, roomNumber int) bool {
	return isRoomCompleteAtLevel(board, roomNumber, 3)
}

func getValidHallwayPositionsToMoveTo(board Board, roomNumber int) []int {
	leftBlockers := getHallwayPositionsToRoom(0, int(roomNumber))
	leftMoves := append([]int{0}, leftBlockers...)

	rightBlockers := getHallwayPositionsToRoom(6, int(roomNumber))
	rightMoves := append(rightBlockers, 6)

	validMoves := []int{}

	for i := len(leftMoves) - 1; i >= 0; i-- {
		pos := leftMoves[i]
		if board.Hallway[pos] == '.' {
			validMoves = append(validMoves, pos)
		} else {
			break
		}
	}

	for _, pos := range rightMoves {
		if board.Hallway[pos] == '.' {
			validMoves = append(validMoves, pos)
		} else {
			break
		}
	}

	return validMoves
}

func canMoveBackToRoom(board Board, hallwayPos int) bool {
	kind := board.Hallway[hallwayPos]
	targetRoom := kind - 'A'

	for i := 0; i < 4; i++ {
		if board.Rooms[targetRoom][i] != '.' && board.Rooms[targetRoom][i] != kind {
			return false
		}
	}

	blockers := getHallwayPositionsToRoom(hallwayPos, int(targetRoom))

	for _, blocker := range blockers {
		if board.Hallway[blocker] != '.' {
			return false
		}
	}

	return true
}

// Gets the hallway positions that are moved through to reach the target position.
// Returns positions from left to right, regardless of what direction the actual move
// is in. Does not include the target position itself.
func getHallwayPositionsToRoom(hallwayPos, roomNumber int) []int {
	switch roomNumber {
	case 0:
		switch hallwayPos {
		case 0:
			return []int{1}
		case 1:
			return []int{}
		case 2:
			return []int{}
		case 3:
			return []int{2}
		case 4:
			return []int{2, 3}
		case 5:
			return []int{2, 3, 4}
		case 6:
			return []int{2, 3, 4, 5}
		}
	case 1:
		switch hallwayPos {
		case 0:
			return []int{1, 2}
		case 1:
			return []int{2}
		case 2:
			return []int{}
		case 3:
			return []int{}
		case 4:
			return []int{3}
		case 5:
			return []int{3, 4}
		case 6:
			return []int{3, 4, 5}
		}
	case 2:
		switch hallwayPos {
		case 0:
			return []int{1, 2, 3}
		case 1:
			return []int{2, 3}
		case 2:
			return []int{3}
		case 3:
			return []int{}
		case 4:
			return []int{}
		case 5:
			return []int{4}
		case 6:
			return []int{4, 5}
		}
	case 3:
		switch hallwayPos {
		case 0:
			return []int{1, 2, 3, 4}
		case 1:
			return []int{2, 3, 4}
		case 2:
			return []int{3, 4}
		case 3:
			return []int{4}
		case 4:
			return []int{}
		case 5:
			return []int{}
		case 6:
			return []int{5}
		}
	}

	panic("Unknown room or hallway position")
}

func getMoveCost(kind byte, roomNumber, roomPos, hallwayPos int) int64 {
	return getMoveDistance(roomNumber, roomPos, hallwayPos) * getEnergyCostOfKind(kind)
}

func getMoveDistance(roomNumber, roomPos, hallwayPos int) int64 {
	var distance int64 = 0

	// Move from room to hallway
	distance += int64(4 - roomPos)

	// Move along the hallway
	distance += getHallwayMoveDistance(roomNumber, hallwayPos)

	return distance
}

func getHallwayMoveDistance(roomNumber, hallwayPos int) int64 {
	switch roomNumber {
	case 0:
		switch hallwayPos {
		case 0:
			return 2
		case 1:
			return 1
		case 2:
			return 1
		case 3:
			return 3
		case 4:
			return 5
		case 5:
			return 7
		case 6:
			return 8
		}
	case 1:
		switch hallwayPos {
		case 0:
			return 4
		case 1:
			return 3
		case 2:
			return 1
		case 3:
			return 1
		case 4:
			return 3
		case 5:
			return 5
		case 6:
			return 6
		}
	case 2:
		switch hallwayPos {
		case 0:
			return 6
		case 1:
			return 5
		case 2:
			return 3
		case 3:
			return 1
		case 4:
			return 1
		case 5:
			return 3
		case 6:
			return 4
		}
	case 3:
		switch hallwayPos {
		case 0:
			return 8
		case 1:
			return 7
		case 2:
			return 5
		case 3:
			return 3
		case 4:
			return 1
		case 5:
			return 1
		case 6:
			return 2
		}
	}

	panic("Unknown room or hallway position")
}

func getEnergyCostOfKind(kind byte) int64 {
	switch kind {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}

	panic("Unknown kind")
}
