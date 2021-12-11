package main

import (
	"fmt"
	"strings"

	"../util"
)

func main() {
	grid, size := parseGrid()

	fmt.Printf("Before any steps:\n")
	printGrid(grid)
	fmt.Println()

	i := 0
	for {
		i++

		flashes := simulate(grid)

		fmt.Printf("After step %d:\n", i)
		printGrid(grid)
		fmt.Println()

		if flashes == size {
			break
		}
	}

	fmt.Println(i)
}

func parseGrid() ([][]int, int) {
	grid := [][]int{}

	size := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		row := []int{}
		for _, cell := range line {
			row = append(row, int(cell-'0'))
			size++
		}

		grid = append(grid, row)
	}

	return grid, size
}

func printGrid(grid [][]int) {
	rows := make([]string, len(grid))

	for i, row := range grid {
		rows[i] = ""
		for _, cell := range row {
			rows[i] += fmt.Sprintf("%d", cell)
		}
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func simulate(grid [][]int) int {
	flashes := 0

	var doFlash func(x, y int)
	doFlash = func(x, y int) {
		flashes++

		grid[y][x] = -1

		for i := -1; i <= 1; i++ {
			if y+i < 0 || y+i >= len(grid) {
				continue
			}

			row := grid[y+i]

			for j := -1; j <= 1; j++ {
				if x+j < 0 || x+j >= len(row) {
					continue
				}

				if grid[y+i][x+j] == -1 {
					// Already flashed
					continue
				}

				grid[y+i][x+j]++

				if grid[y+i][x+j] == 10 {
					doFlash(x+j, y+i)
				}
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == -1 {
				// Already flashed
				continue
			}

			grid[y][x]++

			if grid[y][x] == 10 {
				doFlash(x, y)
			}
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == -1 {
				grid[y][x] = 0
			}
		}
	}

	return flashes
}
