package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func decodeBinary(binaryBits string, start, end int) int {
	bits, err := strconv.ParseInt(binaryBits[start:end], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(bits)
}

func splitSubPackets(binaryBits string, id int, index *int) int {
	if id == 0 {
		val := decodeBinary(binaryBits, *index, *index+15)
		*index += 15
		return val
	} else {
		numOfSubPackets := decodeBinary(binaryBits, *index, *index+11)
		*index += 11
		return numOfSubPackets
	}
}

type subPacket struct {
	val int
}

type subPackets []subPacket

func (s subPackets) sum() int {
	var total int
	for _, subPacket := range s {
		total += subPacket.val
	}

	return total
}

func (s subPackets) multiply() int {
	var total = 1
	for _, subPacket := range s {
		total *= subPacket.val
	}

	return total
}

func (s subPackets) min() int {
	minVal := s[0].val
	for _, subPacket := range s {
		if minVal > subPacket.val {
			minVal = subPacket.val
		}
	}

	return minVal
}

func (s subPackets) max() int {
	maxVal := s[0].val
	for _, subPacket := range s {
		if maxVal < subPacket.val {
			maxVal = subPacket.val
		}
	}

	return maxVal
}

func (s subPackets) gt() int {
	if s[0].val > s[1].val {
		return 1
	}

	return 0
}

func (s subPackets) lt() int {
	if s[0].val < s[1].val {
		return 1
	}
	return 0
}

func (s subPackets) eq() int {
	if s[0].val == s[1].val {
		return 1
	}
	return 0
}

func operator(totalSubPackets subPackets, packetType int) int {
	fmt.Println(packetType)
	switch packetType {
	case 0:
		return totalSubPackets.sum()
	case 1:
		return totalSubPackets.multiply()
	case 2:
		return totalSubPackets.min()
	case 3:
		return totalSubPackets.max()
	case 5:
		return totalSubPackets.gt()
	case 6:
		return totalSubPackets.lt()
	case 7:
		return totalSubPackets.eq()
	}

	return 0
}

func splitPacket(binaryBits string, index *int) (int, int) {
	version := decodeBinary(binaryBits, *index, *index+3)
	*index += 3
	packetType := decodeBinary(binaryBits, *index, *index+3)
	*index += 3
	if packetType == 4 {
		literalNumber := func(binaryBits string, index *int) int {
			var temp string
			for {
				endBit := binaryBits[*index : *index+1]
				temp += binaryBits[*index+1 : *index+5]
				*index += 5
				if endBit == "0" {
					break
				}
			}

			return decodeBinary(temp, 0, len(temp))
		}
		return version, literalNumber(binaryBits, index)

	} else {
		lengthTypeID := decodeBinary(binaryBits, *index, *index+1)
		*index += 1
		splitSubPacketsValue := splitSubPackets(binaryBits, lengthTypeID, index)
		totalSubPackets := subPackets{}

		switch lengthTypeID {
		case 0:
			lastIdx := *index + splitSubPacketsValue
			for *index < lastIdx {
				subPacketVersion, subPacketVal := splitPacket(binaryBits, index)
				totalSubPackets = append(totalSubPackets, subPacket{subPacketVal})
				version += subPacketVersion
			}
		case 1:
			for i := 0; i < splitSubPacketsValue; i++ {
				subPacketVersion, subPacketVal := splitPacket(binaryBits, index)
				totalSubPackets = append(totalSubPackets, subPacket{subPacketVal})
				version += subPacketVersion
			}
		}
		return version, operator(totalSubPackets, packetType)
	}
}

func solution(start int, binaryBits string) {
	version, val := splitPacket(binaryBits, &start)
	fmt.Println(version, val)
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error when opening the input file: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()

	hex := sc.Text()
	binaryBits := ""
	for _, h := range hex {
		bit, _ := strconv.ParseInt(string(h), 16, 64)
		binaryBits += fmt.Sprintf("%04b", bit)
	}
	fmt.Println(len(binaryBits))
	solution(0, binaryBits)
}
