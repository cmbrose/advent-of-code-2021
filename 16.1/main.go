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
	versionBits := s.Take(3)
	version := util.ParseBitString(versionBits)

	typeBits := s.Take(3)
	typeId := util.ParseBitString(typeBits)

	if typeId == 4 {
		valueBits := ""
		for bits := s.Take(5); bits[0] != '0'; bits = s.Take(5) {
			// TODO - this doesn't capture the last one!
			valueBits += bits[1:]
		}
		// Not needed for part 1
		// value := util.ParseBitString(valueBits)

		return version
	}

	opType := s.Take(1)[0]

	subVersionSum := 0
	if opType == '0' {
		lengthBits := s.Take(15)
		length := util.ParseBitString(lengthBits)

		subScanner := s.SubScanner(length)

		for !subScanner.IsDone() {
			subVersionSum += ProcessPacket(subScanner)
		}
	} else {
		packetCountBits := s.Take(11)
		packetCount := util.ParseBitString(packetCountBits)

		for i := 0; i < packetCount; i++ {
			subVersionSum += ProcessPacket(s)
		}
	}

	return version + subVersionSum
}
