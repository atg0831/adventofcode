package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solution1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error when opening input file: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	bits := make([][]int, 0)
	for sc.Scan() {
		input := sc.Text()

		line := make([]int, len(input))
		for idx, i := range input {
			eachBit, err := strconv.Atoi(string(i))
			if err != nil {
				log.Fatalf("error when converting from string to integer: %v", err)
			}
			line[idx] = eachBit
		}
		bits = append(bits, line)
	}

	row := len(bits)
	col := len(bits[0])

	var gammaRate, epsilonRate string
	for c := 0; c < col; c++ {
		cnt := 0
		for r := 0; r < row; r++ {
			if bits[r][c] == 1 {
				cnt += 1
			}
		}

		if 2*cnt > row {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gRDeciaml, _ := strconv.ParseInt(gammaRate, 2, 64)
	eRDecial, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Println(gRDeciaml * eRDecial)
}

func main() {
	solution1()
}
