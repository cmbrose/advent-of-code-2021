package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

type Sector struct {
	MinX int
	MinY int
	MinZ int

	MaxX int
	MaxY int
	MaxZ int

	On bool
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func distance(a, b int) int {
	return max(a, b) - min(a, b)
}

func intersectLineSegments(s1, e1, s2, e2 int) (int, int, bool) {
	if s1 > e2 || e1 < s2 {
		return 0, 0, false
	}

	return max(s1, s2), min(e1, e2), true
}

func (s Sector) String() string {
	onStr := "on"
	if !s.On {
		onStr = "off"
	}

	return fmt.Sprintf("%s x=%d..%d, y=%d..%d, z=%d..%d", onStr, s.MinX, s.MaxX, s.MinY, s.MaxY, s.MinZ, s.MaxZ)
}

func (s Sector) Volume() int {
	dx := distance(s.MinX, s.MaxX) + 1
	dy := distance(s.MinY, s.MaxY) + 1
	dz := distance(s.MinZ, s.MaxZ) + 1

	return dx * dy * dz
}

func (s Sector) Equals(o Sector) bool {
	return s.MinX == o.MinX && s.MaxX == o.MaxX &&
		s.MinY == o.MinY && s.MaxY == o.MaxY &&
		s.MinZ == o.MinZ && s.MaxZ == o.MaxZ
}

func (s Sector) Intersect(o Sector) (Sector, bool) {
	res := Sector{}

	x1, x2, xOk := intersectLineSegments(s.MinX, s.MaxX, o.MinX, o.MaxX)
	y1, y2, yOk := intersectLineSegments(s.MinY, s.MaxY, o.MinY, o.MaxY)
	z1, z2, zOk := intersectLineSegments(s.MinZ, s.MaxZ, o.MinZ, o.MaxZ)

	if !(xOk && yOk && zOk) {
		return res, false
	}

	res.MinX = x1
	res.MaxX = x2
	res.MinY = y1
	res.MaxY = y2
	res.MinZ = z1
	res.MaxZ = z2

	return res, true
}

func (s Sector) Split(intersection Sector) []Sector {
	if s.Equals(intersection) {
		return nil
	}

	sectors := []Sector{}

	var test Sector

	// X blocks fill the whole depth and height

	if intersection.MinX > s.MinX {
		// Left
		test = Sector{
			MinX: s.MinX,
			MinY: s.MinY,
			MinZ: s.MinZ,
			MaxX: intersection.MinX - 1,
			MaxY: s.MaxY,
			MaxZ: s.MaxZ,
			On:   s.On,
		}
		sectors = append(sectors, test)
	}

	if intersection.MaxX < s.MaxX {
		// Right
		test = Sector{
			MinX: intersection.MaxX + 1,
			MinY: s.MinY,
			MinZ: s.MinZ,
			MaxX: s.MaxX,
			MaxY: s.MaxY,
			MaxZ: s.MaxZ,
			On:   s.On,
		}
		sectors = append(sectors, test)
	}

	// Y blocks fill the depth, but are the width of the intersection

	if intersection.MinY > s.MinY {
		// Below
		test = Sector{
			MinX: intersection.MinX,
			MinY: s.MinY,
			MinZ: s.MinZ,
			MaxX: intersection.MaxX,
			MaxY: intersection.MinY - 1,
			MaxZ: s.MaxZ,
			On:   s.On,
		}
		sectors = append(sectors, test)
	}

	if intersection.MaxY < s.MaxY {
		// Above
		test = Sector{
			MinX: intersection.MinX,
			MinY: intersection.MaxY + 1,
			MinZ: s.MinZ,
			MaxX: intersection.MaxX,
			MaxY: s.MaxY,
			MaxZ: s.MaxZ,
			On:   s.On,
		}
		sectors = append(sectors, test)
	}

	// Finally, Z blocks are the dimension of the intersection

	if intersection.MinZ > s.MinZ {
		// Forward (towards you)
		test = Sector{
			MinX: intersection.MinX,
			MinY: intersection.MinY,
			MinZ: s.MinZ,
			MaxX: intersection.MaxX,
			MaxY: intersection.MaxY,
			MaxZ: intersection.MinZ - 1,
			On:   s.On,
		}
		sectors = append(sectors, test)
	}

	if intersection.MaxZ < s.MaxZ {
		// Behind (away from you)
		test = Sector{
			MinX: intersection.MinX,
			MinY: intersection.MinY,
			MinZ: intersection.MaxZ + 1,
			MaxX: intersection.MaxX,
			MaxY: intersection.MaxY,
			MaxZ: s.MaxZ,
			On:   s.On,
		}
		sectors = append(sectors, test)
	}

	return sectors
}

func main() {
	initial := Sector{
		-50,
		-50,
		-50,
		50,
		50,
		50,
		false,
	}

	sectors := []Sector{initial}

	printStats(sectors, true)

	for _, line := range util.ReadInputLines("./input.txt") {
		changeSector := parseLine(line)

		if !isInBounds(changeSector) {
			fmt.Println("Skipping", changeSector)
			continue
		}

		startLen := len(sectors)
		for i := 0; i < startLen; i++ {
			s := sectors[i]
			if s.On == changeSector.On {
				continue
			}

			intersection, ok := s.Intersect(changeSector)
			if !ok {
				continue
			}

			intersection.On = changeSector.On
			sectors[i] = intersection

			if intersection.String() == "off x=2..23, y=23..26, z=-50..-22" {

				fmt.Println("It's an intersection")
			}

			newSectors := s.Split(intersection)
			if newSectors == nil {
				// The intersection is entirely the source sector, so just update it
				continue
			}

			vol := 0
			for _, ns := range newSectors {
				if ns.String() == "off x=2..23, y=23..26, z=-50..-22" {

					fmt.Println("It's a split")
				}
				vol += ns.Volume()
			}
			if vol+intersection.Volume() != s.Volume() {
				panic("Lost volume")
			}

			sectors = append(sectors, newSectors...)
		}

		printStats(sectors, false)
	}

	printStats(sectors, false)
}

func parseLine(line string) Sector {
	// on x=-20..26,y=-36..17,z=-47..7
	// 0  1 2    3  4 5    6  7 8    9
	parts := strings.FieldsFunc(line, func(c rune) bool { return c == ' ' || c == ',' || c == '=' || c == '.' })

	s := Sector{}
	var err error

	s.On = parts[0] == "on"

	s.MinX, err = strconv.Atoi(parts[2])
	util.Check(err)
	s.MaxX, err = strconv.Atoi(parts[3])
	util.Check(err)
	s.MinY, err = strconv.Atoi(parts[5])
	util.Check(err)
	s.MaxY, err = strconv.Atoi(parts[6])
	util.Check(err)
	s.MinZ, err = strconv.Atoi(parts[8])
	util.Check(err)
	s.MaxZ, err = strconv.Atoi(parts[9])
	util.Check(err)

	if s.MinX > s.MaxX {
		panic("X valus out of order")
	}
	if s.MinY > s.MaxY {
		panic("Y valus out of order")
	}
	if s.MinZ > s.MaxZ {
		panic("Z valus out of order")
	}

	return s
}

func printStats(sectors []Sector, printSectors bool) {
	cnt := 0
	vol := 0
	for _, s := range sectors {
		vol += s.Volume()

		if printSectors {
			fmt.Println(s)
		}

		if s.On {
			cnt += s.Volume()
		}
	}
	fmt.Println("Vol =", vol, ", Cnt =", cnt)
}

func isInBounds(s Sector) bool {
	return s.MinX >= -50 && s.MaxX <= 50 &&
		s.MinY >= -50 && s.MaxY <= 50 &&
		s.MinZ >= -50 && s.MaxZ <= 50
}
