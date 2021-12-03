package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

func main() {
	forward := 0
	depth := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		pair := strings.SplitN(line, " ", 2)

		command := pair[0]

		distance, err := strconv.Atoi(pair[1])
		util.Check(err)

		switch command {
		case "forward":
			forward += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	fmt.Printf("%d * %d = %d\n", forward, depth, forward*depth)
}
