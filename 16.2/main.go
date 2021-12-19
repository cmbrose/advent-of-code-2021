package main

import (
	"bytes"
	"fmt"

	"../util"
)

func main() {
	input := util.ReadInputLines("./input.txt")[0]

	scanner := newScanner(input)

	sum := 0
	for !scanner.IsDone() {
		sum += ProcessPacket(scanner)
	}

	fmt.Println(sum)
}

type Scanner struct {
	hex string
	bin string

	// extra binary that exists after hex is used up. Only relevant for sub-scanners
	tailBin string
}

func newScanner(hex string) *Scanner {
	return &Scanner{hex, "", ""}
}

func (s *Scanner) Take(n int) string {
	var out bytes.Buffer

	for n > 0 {
		if len(s.bin) == 0 {
			if len(s.hex) > 0 {
				s.bin = util.HexToBinary(s.hex[0])
				s.hex = s.hex[1:]
			} else {
				s.bin = s.tailBin
				s.tailBin = ""
			}
		}

		if n > len(s.bin) {
			out.WriteString(s.bin)
			n -= len(s.bin)
			s.bin = ""
		} else {
			out.WriteString(s.bin[:n])
			s.bin = s.bin[n:]
			n = 0
		}
	}

	return out.String()
}

func (s *Scanner) SubScanner(length int) *Scanner {
	sub := &Scanner{"", "", ""}

	if len(s.bin) >= length {
		sub.bin = s.bin[:length]
		s.bin = s.bin[length:]
		return sub
	}

	sub.bin = s.bin
	length -= len(s.bin)
	s.bin = ""

	hexLength := length / 4
	sub.hex = s.hex[:hexLength]
	length -= hexLength * 4
	s.hex = s.hex[hexLength:]

	sub.tailBin = s.Take(length)

	return sub
}

func (s Scanner) IsDone() bool {
	if s.Len() == 0 {
		return true
	}

	for _, b := range s.bin {
		if b != '0' {
			return false
		}
	}

	for _, h := range s.hex {
		if h != '0' {
			return false
		}
	}

	return true
}

func (s Scanner) Len() int {
	return len(s.hex)*4 + len(s.bin) + len(s.tailBin)
}

// Returns sum of versions
func ProcessPacket(s *Scanner) int {
	s.Take(3)
	//versionBits := s.Take(3)
	//version := util.ParseBitString(versionBits)

	typeBits := s.Take(3)
	typeId := util.ParseBitString(typeBits)

	if typeId == 4 {
		valueBits := ""

		bits := s.Take(5)
		for ; bits[0] != '0'; bits = s.Take(5) {
			valueBits += bits[1:]
		}
		valueBits += bits[1:]

		value := util.ParseBitString(valueBits)

		return value
	}

	acc := GetAccumulator(typeId)

	opType := s.Take(1)[0]

	var result int
	if opType == '0' {
		lengthBits := s.Take(15)
		length := util.ParseBitString(lengthBits)

		subScanner := s.SubScanner(length)

		for !subScanner.IsDone() {
			value := ProcessPacket(subScanner)
			result = acc(value)
		}
	} else {
		packetCountBits := s.Take(11)
		packetCount := util.ParseBitString(packetCountBits)

		for i := 0; i < packetCount; i++ {
			value := ProcessPacket(s)
			result = acc(value)
		}
	}

	return result
}

type Accumulator func(value int) int

func GetAccumulator(typeId int) Accumulator {
	switch typeId {
	case 0:
		return GetSumAccumulator()
	case 1:
		return GetProductAccumulator()
	case 2:
		return GetMinAccumulator()
	case 3:
		return GetMaxAccumulator()
	case 5:
		return GetGreaterThanAccumulator()
	case 6:
		return GetLessThanAccumulator()
	case 7:
		return GetEqualToAccumulator()
	}

	panic(fmt.Sprintf("Unknown typeId: %d", typeId))
}

func GetSumAccumulator() Accumulator {
	acc := 0
	return func(value int) int {
		acc += value
		return acc
	}
}

func GetProductAccumulator() Accumulator {
	isFirst := true
	acc := 0
	return func(value int) int {
		if isFirst {
			acc = value
		} else {
			acc *= value
		}

		isFirst = false
		return acc
	}
}

func GetMinAccumulator() Accumulator {
	isFirst := true
	acc := 0
	return func(value int) int {
		if isFirst {
			acc = value
		} else if value < acc {
			acc = value
		}

		isFirst = false
		return acc
	}
}

func GetMaxAccumulator() Accumulator {
	isFirst := true
	acc := 0
	return func(value int) int {
		if isFirst {
			acc = value
		} else if value > acc {
			acc = value
		}

		isFirst = false
		return acc
	}
}

func GetGreaterThanAccumulator() Accumulator {
	isFirst := true
	acc := 0
	return func(value int) int {
		if isFirst {
			acc = value
		} else if acc > value {
			acc = 1
		} else {
			acc = 0
		}

		isFirst = false
		return acc
	}
}

func GetLessThanAccumulator() Accumulator {
	isFirst := true
	acc := 0
	return func(value int) int {
		if isFirst {
			acc = value
		} else if acc < value {
			acc = 1
		} else {
			acc = 0
		}

		isFirst = false
		return acc
	}
}

func GetEqualToAccumulator() Accumulator {
	isFirst := true
	acc := 0
	return func(value int) int {
		if isFirst {
			acc = value
		} else if acc == value {
			acc = 1
		} else {
			acc = 0
		}

		isFirst = false
		return acc
	}
}
