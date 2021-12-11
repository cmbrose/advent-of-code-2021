package main

import (
	"fmt"
	"sort"
	"strings"

	"../util"
)

func main() {
	sum := 0
	for _, line := range util.ReadInputLines("./input.txt") {
		digits := strings.Split(line, " | ")[0]
		output := strings.Split(line, " | ")[1]

		digitValues := strings.Split(digits, " ")

		outputValues := strings.Split(output, " ")

		value := solve(digitValues, outputValues)
		sum += value
	}

	fmt.Println(sum)
}

func solve(digits, outputs []string) int {
	var one []interface{}
	var four []interface{}
	var seven []interface{}
	var eight []interface{}
	fiveSegs := [][]interface{}{}
	sixSegs := [][]interface{}{}

	for _, digit := range digits {
		runes := []rune(digit)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		runesInterface := util.RuneSliceToInterfaceSlice(runes)

		switch len(runes) {
		case 2:
			one = runesInterface
		case 3:
			seven = runesInterface
		case 4:
			four = runesInterface
		case 5:
			fiveSegs = append(fiveSegs, runesInterface)
		case 6:
			sixSegs = append(sixSegs, runesInterface)
		case 7:
			eight = runesInterface
		}
	}

	fiveIntersect := util.IntersectAll(fiveSegs...)
	sixIntersect := util.IntersectAll(sixSegs...)

	topRes := util.Except(seven, one)
	assertSingle(topRes, "top")

	centerRes := util.Intersect(four, fiveIntersect)
	assertSingle(centerRes, "center")

	bottomRes := util.ExceptAll(sixIntersect, four, topRes)
	assertSingle(bottomRes, "bottom")

	bottomRightRes := util.Intersect(sixIntersect, one)
	assertSingle(bottomRightRes, "bottom right")

	topRightRes := util.Except(one, bottomRightRes)
	topRight := assertSingle(topRightRes, "top right")

	topLeftRes := util.ExceptAll(four, centerRes, one)
	topLeft := assertSingle(topLeftRes, "top left")

	bottomLeftRes := util.ExceptAll(eight, four, topRes, bottomRes)
	bottomLeft := assertSingle(bottomLeftRes, "bottom left")

	val := 0
	for _, output := range outputs {
		var position int
		switch len(output) {
		case 2:
			position = 1
		case 3:
			position = 7
		case 4:
			position = 4
		case 5:
			// 2, 3, or 5
			if strings.ContainsRune(output, topLeft) {
				position = 5
			} else if strings.ContainsRune(output, bottomLeft) {
				position = 2
			} else {
				position = 3
			}
		case 6:
			// 0, 6, or 9
			if !strings.ContainsRune(output, topRight) {
				position = 6
			} else if !strings.ContainsRune(output, bottomLeft) {
				position = 9
			} else {
				position = 0
			}
		case 7:
			position = 8
		}

		val = val*10 + position
	}

	return val
}

func assertSingle(res []interface{}, name string) rune {
	if len(res) != 1 {
		panic(fmt.Sprintf("Could not get %s segment", name))
	}
	val := res[0].(rune)
	//fmt.Printf("%s = %c\n", name, val)

	return val
}
