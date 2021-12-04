package util

import (
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInputLines(path string) []string {
	content, err := os.ReadFile("./input.txt")
	Check(err)

	return strings.Split(string(content), "\n")
}

func ParseBitString(str string) int {
	val := 0
	for _, c := range str {
		val <<= 1
		if c == '1' {
			val++
		}
	}

	return val
}
