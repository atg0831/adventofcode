package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	// 14, 18
	var instructions [][]string
	for sc.Scan() {
		instructions = append(instructions, strings.Split(sc.Text(), " "))
	}

	return instructions
}

type monadField struct {
	z int
	w string
}

type monadFields []monadField
type variableNumber struct {
	v1, v2, v3 int
}

type variableNumbers []variableNumber

func main() {
	instructions := readInput()
	vs := &variableNumbers{}
	vs.getVariableNumbers(instructions)
	fmt.Println(*vs)

	var resultMonads = monadFields{{0, ""}} // last z value is 0(start value z from end to first)
	for i := len(*vs) - 1; i >= 0; i-- {
		fmt.Println(i)
		tempMonads := make(monadFields, 0)
		for _, resultMonad := range resultMonads {
			resultMonad.calculateCurrentZFromNextZ((*vs)[i], &tempMonads)
		}

		resultMonads = tempMonads
	}
	maxValue := 0
	for _, resultMonad := range resultMonads {
		intW, err := strconv.Atoi(resultMonad.w)
		if err != nil {
			log.Fatal(err)
		}
		if intW > maxValue {
			maxValue = intW
		}
	}

	minValue := maxValue
	for _, resultMonad := range resultMonads {
		intW, err := strconv.Atoi(resultMonad.w)
		if err != nil {
			log.Fatal(err)
		}
		if intW < minValue {
			minValue = intW
		}
	}

	fmt.Println("part1: ", maxValue)
	fmt.Println("part2: ", minValue)

}

func (r *monadField) calculateCurrentZFromNextZ(v variableNumber, tempMonad *monadFields) {
	for w := 1; w <= 9; w++ {
		zList := calculateRangeOfZ(v, w, r.z)
		for _, z := range zList {
			if calculateNextZ(v, w, z) == r.z {
				*tempMonad = append(*tempMonad, monadField{z, strconv.Itoa(w) + r.w})
			}
		}
	}
}

func calculateNextZ(v variableNumber, w, z int) int {
	var x int
	if (z%26 + v.v2) != w {
		x = 1
	}
	div := int(z / v.v1)
	return div*(25*x+1) + (w+v.v3)*x //next z

}

// x=1 -> (z/v.v1)*(26) + w+v.v3  = nextZ ==>  z=(v.v1) *(nextz - (w+v.v3))/26 or (v.v1) *((nextz - (w+v.v3))/26)+1)
// x=0 -> z/v.v1 = nextZ ==> z= nextZ*v.v1 or z = (nextZ+1) * v.v1
func calculateRangeOfZ(v variableNumber, w, nextZ int) []int {
	zList := make([]int, 0)

	// x = 0
	for z := v.v1 * nextZ; z < v.v1*(nextZ+1); z++ {
		zList = append(zList, z)
	}

	// x = 1
	for z := v.v1 * (nextZ - (w + v.v3)) / 26; z < v.v1*((nextZ-(w+v.v3))/26+1); z++ {
		zList = append(zList, z)
	}

	return zList
}

func (vs *variableNumbers) getVariableNumbers(instructions [][]string) {
	v := variableNumber{}
	for i := 0; i < 14*18; i++ {
		remainder := i % 18
		switch remainder {
		case 4:
			num, err := strconv.Atoi(instructions[i][2])
			if err != nil {
				log.Fatal(err)
			}
			v.v1 = num
		case 5:
			num, err := strconv.Atoi(instructions[i][2])
			if err != nil {
				log.Fatal(err)
			}
			v.v2 = num
		case 15:
			num, err := strconv.Atoi(instructions[i][2])
			if err != nil {
				log.Fatal(err)
			}
			v.v3 = num
		}

		if remainder == 15 {
			*vs = append(*vs, v)
		}
	}

}
