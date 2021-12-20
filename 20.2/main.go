package main

import (
	"bytes"
	"fmt"

	"../util"
)

type OnOffPair []bool

func smushRows(top []OnOffPair, arr [][]OnOffPair, bottom []OnOffPair) [][]OnOffPair {
	arr = append(arr, bottom, nil) // bottom and then nil because it's about to get shifted down
	copy(arr[1:], arr)
	arr[0] = top
	return arr
}

func smushValues(left OnOffPair, arr []OnOffPair, right OnOffPair) []OnOffPair {
	arr = append(arr, right, nil) // right and then nil because it's about to get shifted over
	copy(arr[1:], arr)
	arr[0] = left
	return arr
}

func main() {
	lines := util.ReadInputLines("./input.txt")
	algoStr := lines[0]
	algo := make([]bool, len(algoStr))
	for i, c := range algoStr {
		algo[i] = c == '#'
	}

	if algo[0] && algo[len(algo)-1] {
		panic("Background is always on!")
	}

	on := 0
	off := 1
	var background OnOffPair = make([]bool, 2)
	background[on] = false
	background[off] = algo[0]

	grid := make([][]OnOffPair, len(lines)-2)

	for y, line := range lines[2:] {
		grid[y] = make([]OnOffPair, len(line))
		for x, c := range line {
			grid[y][x] = make([]bool, 2)
			grid[y][x][on] = c == '#'
		}
	}

	for iter := 0; iter < 50; iter++ {
		// printGrid(grid, background, on)

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				grid[y][x][off] = nextValue(x, y, grid, background, on, algo)
			}
		}

		// Create the left and right fills, but don't add them until all
		// of them have been calculated!
		leftFills := make([][]bool, len(grid))
		rightFills := make([][]bool, len(grid))
		for y := 0; y < len(grid); y++ {
			leftFills[y] = make([]bool, 2)
			leftFills[y][on] = background[on]
			leftFills[y][off] = nextValue(-1, y, grid, background, on, algo)

			rightFills[y] = make([]bool, 2)
			rightFills[y][on] = background[on]
			rightFills[y][off] = nextValue(len(grid[y]), y, grid, background, on, algo)
		}

		// Now add the left and right fills all at once
		for y := 0; y < len(grid); y++ {
			grid[y] = smushValues(leftFills[y], grid[y], rightFills[y])
		}

		// Next top and bottom fills
		topFill := make([]OnOffPair, len(grid[0])) // Now 2 longer from the above loop
		bottFill := make([]OnOffPair, len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			topFill[x] = make([]bool, 2)
			topFill[x][on] = background[on]
			topFill[x][off] = nextValue(x, -1, grid, background, on, algo)

			bottFill[x] = make([]bool, 2)
			bottFill[x][on] = background[on]
			bottFill[x][off] = nextValue(x, len(grid), grid, background, on, algo)
		}
		grid = smushRows(topFill, grid, bottFill)

		on = 1 - on
		off = 1 - off
	}

	// printGrid(grid, background, on)

	cnt := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x][on] {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

func getValue(x, y int, grid [][]OnOffPair, background OnOffPair, on int) bool {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return background[on]
	}

	return grid[y][x][on]
}

func nextValue(x, y int, grid [][]OnOffPair, background OnOffPair, on int, algo []bool) bool {
	val := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			val <<= 1
			cell := getValue(x+j, y+i, grid, background, on)
			if cell {
				val |= 1
			}
		}
	}

	return algo[val]
}

func printGrid(grid [][]OnOffPair, background OnOffPair, on int) {
	toChar := func(val bool) rune {
		if val {
			return '#'
		}
		return '.'
	}

	bgChar := toChar(background[on])

	totalWidth := len(grid[0]) + 2 + 3*2 // +2 is the border, 3*2 is the inf space

	var out bytes.Buffer

	topBottBorder := repeat(bgChar, totalWidth) + "\n"
	edgeBorder := repeat(bgChar, 3)

	out.WriteString(topBottBorder)
	out.WriteString(topBottBorder)

	out.WriteString(edgeBorder)
	out.WriteRune('+')
	out.WriteString(repeat('-', len(grid[0])))
	out.WriteRune('+')
	out.WriteString(edgeBorder)
	out.WriteRune('\n')

	for _, row := range grid {
		out.WriteString(edgeBorder)
		out.WriteRune('|')

		for _, val := range row {
			out.WriteRune(toChar(val[on]))
		}

		out.WriteRune('|')
		out.WriteString(edgeBorder)
		out.WriteRune('\n')
	}

	out.WriteString(edgeBorder)
	out.WriteRune('+')
	out.WriteString(repeat('-', len(grid[0])))
	out.WriteRune('+')
	out.WriteString(edgeBorder)
	out.WriteRune('\n')

	out.WriteString(topBottBorder)
	out.WriteString(topBottBorder)

	fmt.Println(out.String())
}

func repeat(r rune, cnt int) string {
	var out bytes.Buffer

	for i := 0; i < cnt; i++ {
		out.WriteRune(r)
	}

	return out.String()
}
