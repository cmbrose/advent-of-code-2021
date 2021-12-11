package main

import (
	"fmt"
	"strings"

	"../util"
)

func main() {
	walls := parseMap()

	max1 := 0
	max2 := 0
	max3 := 0

	handleResult := func(size int) {
		if size > max1 {
			max3 = max2
			max2 = max1
			max1 = size
		} else if size > max2 {
			max3 = max2
			max2 = size
		} else if size > max3 {
			max3 = size
		}
	}

	for y := 0; y < len(walls); y++ {
		for x := 0; x < len(walls[y]); x++ {
			if walls[y][x] == ' ' {
				size := scan(x, y, walls)
				handleResult(size)
			}
		}
	}

	printMap(walls)

	fmt.Printf("%d * %d * %d = %d\n", max1, max2, max3, max1*max2*max3)
}

func printMap(walls [][]rune) {
	rows := make([]string, len(walls))

	for i, row := range walls {
		rows[i] = string(row)
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func parseMap() [][]rune {
	wallVals := [][]rune{}

	for _, line := range util.ReadInputLines("./input.txt") {
		wallRow := []rune{}
		for _, cell := range line {
			if cell == '9' {
				wallRow = append(wallRow, 'X')
			} else {
				wallRow = append(wallRow, ' ')
			}
		}

		wallVals = append(wallVals, wallRow)
	}

	return wallVals
}

type Point struct {
	X int
	Y int
}

func scan(x, y int, walls [][]rune) int {
	// Implements the FloodFill algorithm here: https://en.wikipedia.org/wiki/Flood_fill#Span_Filling

	queue := []Point{
		{X: x, Y: y},
	}

	size := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		x := p.X
		y := p.Y

		leftX := x
		for leftX-1 >= 0 && walls[y][leftX-1] == ' ' {
			size++
			walls[y][leftX-1] = '.'
			leftX--
		}

		rightX := x
		for rightX < len(walls[y]) && walls[y][rightX] == ' ' {
			size++
			walls[y][rightX] = '.'
			rightX++
		}

		queue = enequeSpans(leftX, rightX-1, y-1, walls, queue)
		queue = enequeSpans(leftX, rightX-1, y+1, walls, queue)
	}

	return size
}

func enequeSpans(leftX, rightX, y int, walls [][]rune, queue []Point) []Point {
	if y < 0 || y >= len(walls) {
		return queue
	}

	added := false
	for x := leftX; x <= rightX; x++ {
		if walls[y][x] == 'X' {
			added = false
		} else if !added {
			queue = append(queue, Point{X: x, Y: y})
			added = true
		}
	}

	return queue
}
