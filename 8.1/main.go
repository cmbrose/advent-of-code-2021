package main

import (
	"fmt"
	"strings"

	"../util"
)

func main() {
	cnt := 0
	for _, line := range util.ReadInputLines("./input.txt") {
		output := strings.Split(line, " | ")[1]
		outputDigits := strings.Split(output, " ")

		for _, digitCode := range outputDigits {
			switch len(digitCode) {
			case 2, 3, 4, 7:
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
