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

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func MaxInt(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// Must be pre-sorted!
func Intersect(a, b []interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return []interface{}{}
	}

	res := []interface{}{}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		comp := Compare(a, b)
		if comp == 0 {
			res = append(res, a)
			i++
			j++
		} else if comp > 0 { // a > b
			j++
		} else { // a < b
			i++
		}
	}

	return res
}

func IntersectAll(a ...[]interface{}) []interface{} {
	if len(a) == 0 {
		return []interface{}{}
	}

	cur := a[0]

	if len(a) == 1 {
		return cur
	}

	for _, b := range a[1:] {
		cur = Intersect(cur, b)
	}

	return cur
}

// Must be pre-sorted!
func Except(a, b []interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	res := []interface{}{}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		comp := Compare(a, b)
		if comp == 0 {
			i++
			j++
		} else if comp > 0 { // a > b
			j++
		} else { // a < b
			res = append(res, a)
			i++
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}

	return res
}

// Must be pre-sorted!
func ExceptAll(a []interface{}, b ...[]interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	cur := a

	for _, b := range b {
		cur = Except(cur, b)
	}

	return cur
}

func RuneSliceToInterfaceSlice(a []rune) []interface{} {
	res := make([]interface{}, len(a))

	for i, a := range a {
		res[i] = a
	}

	return res
}

func Compare(a, b interface{}) int {
	switch a := a.(type) {
	case int:
		b, ok := b.(int)
		if !ok {
			panic("type mismatch")
		}
		return a - b

	case rune:
		b, ok := b.(rune)
		if !ok {
			panic("type mismatch")
		}
		return int(a) - int(b)

	default:
		panic("Unhandled type")
	}
}
