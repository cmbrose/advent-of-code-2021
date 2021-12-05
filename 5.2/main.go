package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

func main() {
	// States are:
	// - Not in map => count is 0
	// - In map and true => count is 1
	// - In map and false => count > 1 (don't count again)
	seenPoints := make(map[Point]bool)

	cnt := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		line := parseLine(line)

		lw := newLineWalker(line)

		for hasNext := true; hasNext; hasNext = lw.step() {
			countIt, ok := seenPoints[lw.cur]

			if countIt && ok {
				cnt++
				seenPoints[lw.cur] = false
			} else if !ok {
				seenPoints[lw.cur] = true
			}
		}
	}

	fmt.Println(cnt)
}

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func parseLine(str string) Line {
	// Who actually cares about formatting, just get the numbers only :)
	sepChars := func(c rune) bool {
		return c == ' ' || c == ',' || c == '-' || c == '>'
	}

	parts := strings.FieldsFunc(str, sepChars)

	x1, err := strconv.Atoi(parts[0])
	util.Check(err)
	y1, err := strconv.Atoi(parts[1])
	util.Check(err)
	x2, err := strconv.Atoi(parts[2])
	util.Check(err)
	y2, err := strconv.Atoi(parts[3])
	util.Check(err)

	return Line{
		Start: Point{X: x1, Y: y1},
		End:   Point{X: x2, Y: y2},
	}
}

type LineWalker struct {
	line Line

	xStep int
	yStep int

	cur Point
}

func newLineWalker(line Line) *LineWalker {
	walker := &LineWalker{line: line, cur: line.Start}

	getStep := func(s, e int) int {
		if s < e {
			return 1
		} else if s > e {
			return -1
		} else {
			return 0
		}
	}

	walker.xStep = getStep(line.Start.X, line.End.X)
	walker.yStep = getStep(line.Start.Y, line.End.Y)

	return walker
}

func (lw *LineWalker) step() bool {
	if lw.line.End == lw.cur {
		return false
	}

	lw.cur.X += lw.xStep
	lw.cur.Y += lw.yStep

	return true
}
