package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"../util"
)

type Result byte

const (
	Success Result = iota
	FailTooShort
	FailTooFar
	FailPassedThrough
	Continue
)

func (r Result) String() string {
	switch r {
	case Success:
		return "Success"
	case FailTooShort:
		return "FailTooShort"
	case FailTooFar:
		return "FailTooFar"
	case FailPassedThrough:
		return "FailPassedThrough"
	case Continue:
		return "Continue"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

type CheckShotFn func(x, y, xVel, yVel int) Result

func main() {
	input := util.ReadInputLines("./input.txt")[0]

	parts := strings.FieldsFunc(input, func(c rune) bool { return c == ' ' || c == ',' || c == '=' || c == '.' })

	// target area: x=20..30, y=-10..-5
	// 0      1     2 3   4   5  6   7

	xMin, err := strconv.Atoi(parts[3])
	util.Check(err)
	xMax, err := strconv.Atoi(parts[4])
	util.Check(err)

	yMin, err := strconv.Atoi(parts[6])
	util.Check(err)
	yMax, err := strconv.Atoi(parts[7])
	util.Check(err)

	check := func(x, y, xVel, yVel int) Result {
		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			return Success
		}

		if x > xMax {
			return FailTooFar
		}

		if x < xMin && xVel == 0 {
			return FailTooShort
		}

		if x >= xMin && y < yMin {
			return FailPassedThrough
		}

		return Continue
	}

	var yVelCutoff int
	if yMax < 0 {
		yVelCutoff = -yMin
	} else if yMin > 0 {
		yVelCutoff = yMax
	} else {
		// The shot always returns to y=0. That means if the box contains y=0,
		// then as long as the x velocity is OK, any y velocity is also OK.
		// So it would go infinitely high...
		panic("Can't have a box around y=0!!")
	}

	maxHeight := math.MinInt
	for xVel := 1; xVel <= xMax; xVel++ {
		for yVel := 0; yVel < yVelCutoff; yVel++ {
			height, result := shoot(xVel, yVel, check)
			if result != Success {
				continue
			}

			if height > maxHeight {
				maxHeight = height
				fmt.Printf("New max: %d,%d => %d\n", xVel, yVel, maxHeight)
			}
		}
	}
}

func shoot(xVel, yVel int, check CheckShotFn) (int, Result) {
	x := 0
	y := 0

	yMax := 0

	var res Result
	for res = Continue; res == Continue; res = check(x, y, xVel, yVel) {
		x += xVel
		y += yVel

		if y > yMax {
			yMax = y
		}

		if xVel < 0 {
			xVel += 1
		} else if xVel > 0 {
			xVel -= 1
		}

		yVel -= 1
	}

	return yMax, res
}
