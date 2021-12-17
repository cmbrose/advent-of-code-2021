package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

type Point struct {
	X int
	Y int
}

func main() {
	lines := util.ReadInputLines("./input.txt")

	maxX := 0
	maxY := 0
	points := []Point{}

	i := 0
	for lines[i] != "" {
		line := lines[i]
		pair := strings.Split(line, ",")

		x, err := strconv.Atoi(pair[0])
		util.Check(err)

		if maxX < x {
			maxX = x
		}

		y, err := strconv.Atoi(pair[1])
		util.Check(err)

		if maxY < y {
			maxY = y
		}

		points = append(points, Point{X: x, Y: y})

		i++
	}

	// Skip blank line
	i++

	grid := initGrid(maxX+1, maxY+1)

	for _, p := range points {
		grid[p.Y][p.X] = true
	}

	//printGrid(grid)

	for i < len(lines) {
		line := lines[i]

		parts := strings.Split(line, " ")
		fold := parts[2]

		pair := strings.Split(fold, "=")

		isXAxis := pair[0] == "x"

		pos, err := strconv.Atoi(pair[1])
		util.Check(err)

		if isXAxis {
			grid = foldLeft(pos, grid)
		} else {
			grid = foldUp(pos, grid)
		}

		i++
	}

	printGrid(grid)
}

func printGrid(grid [][]bool) {
	rows := make([]string, len(grid))

	for i, row := range grid {
		rows[i] = ""
		for _, cell := range row {
			if cell {
				rows[i] += "#"
			} else {
				rows[i] += " "
			}
		}
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func initGrid(x, y int) [][]bool {
	grid := [][]bool{}

	for i := 0; i < y; i++ {
		grid = append(grid, make([]bool, x))
	}

	return grid
}

func foldUp(pos int, grid [][]bool) [][]bool {
	foldTarget := grid[pos+1:]
	grid = grid[:pos]

	for y := 0; y < len(foldTarget); y++ {
		gridRow := len(grid) - y - 1

		for x, cell := range foldTarget[y] {
			grid[gridRow][x] = grid[gridRow][x] || cell
		}
	}

	return grid
}

func foldLeft(pos int, grid [][]bool) [][]bool {
	for y, row := range grid {
		foldTarget := row[pos+1:]
		row = row[:pos]

		for x, cell := range foldTarget {
			rowCell := len(row) - x - 1

			row[rowCell] = row[rowCell] || cell
		}

		grid[y] = row
	}

	return grid
}
