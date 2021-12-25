package main

import (
	"bytes"
	"fmt"

	"../util"
)

type Grid [][]rune
type MoveCheck [][]bool

func parseGrid() (Grid, MoveCheck) {
	grid := Grid{}
	check := MoveCheck{}
	lines := util.ReadInputLines("./input.txt")

	for _, line := range lines {
		row := []rune{}
		checkRow := []bool{}
		for _, cell := range line {
			row = append(row, cell)
			checkRow = append(checkRow, false)
		}
		grid = append(grid, row)
		check = append(check, checkRow)
	}

	return grid, check
}

func printGrid(grid Grid) {
	var out bytes.Buffer

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			cell := grid[y][x]
			out.WriteRune(rune(cell))
		}

		out.WriteRune('\n')
	}

	fmt.Println(out.String())
}

func doMove(grid Grid, check MoveCheck) {
	for y, row := range check {
		for x := range row {
			if row[x] {
				kind := grid[y][x]

				if kind == '>' {
					newX := (x + 1) % len(row)
					grid[y][newX] = '>'
					grid[y][x] = '.'
				} else {
					newY := (y + 1) % len(grid)
					grid[newY][x] = 'v'
					grid[y][x] = '.'
				}
			}

			row[x] = false
		}
	}
}

func checkMove(x, y int, targetKind rune, grid Grid) bool {
	kind := grid[y][x]

	if kind != targetKind {
		return false
	}

	newX := x
	newY := y

	if kind == '>' {
		newX = (x + 1) % len(grid[0])
	} else {
		newY = (y + 1) % len(grid)
	}

	if grid[newY][newX] != '.' {
		return false
	}

	return true
}

func main() {
	grid, check := parseGrid()

	anyMoved := true
	moveNumber := 0
	for anyMoved {
		anyMoved = false
		moveNumber++

		stepPause(grid)

		for x := len(grid[0]) - 1; x >= 0; x-- {
			for y := len(grid) - 1; y >= 0; y-- {
				ok := checkMove(x, y, '>', grid)
				check[y][x] = ok
				anyMoved = anyMoved || ok
			}
		}

		doMove(grid, check)
		stepPause(grid)

		for y := len(grid) - 1; y >= 0; y-- {
			for x := len(grid[0]) - 1; x >= 0; x-- {
				ok := checkMove(x, y, 'v', grid)
				check[y][x] = ok
				anyMoved = anyMoved || ok
			}
		}

		doMove(grid, check)
	}

	printGrid(grid)
	fmt.Println(moveNumber)
}

func stepPause(grid Grid) {
	enabled := false

	if enabled {
		printGrid(grid)
		fmt.Println("Press a key to continue")
		fmt.Scanln()
	}
}
