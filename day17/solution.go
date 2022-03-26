package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]string, []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error when opening file: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()

	input := sc.Text()
	targetGroup := strings.Split(strings.Split(input, ":")[1:][0], ",") // [x=a..b, y=c..d]
	targetXRange := strings.Split(strings.Split(targetGroup[0], "=")[1], "..")
	targetYRange := strings.Split(strings.Split(targetGroup[1], "=")[1], "..")

	return targetXRange, targetYRange
}

func convertToInteger(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}
func main() {
	targetXRange, targetYRange := readInput()

	lowerboundX := convertToInteger(targetXRange[0])
	upperboundX := convertToInteger(targetXRange[1])

	lowerboundY := convertToInteger(targetYRange[1])
	upperboundY := convertToInteger(targetYRange[0])
	// startX, startY := 0, 0

	answer := 0
	for xvelocity := int(math.Sqrt(float64(2*upperboundX - 1))); xvelocity >= int(math.Sqrt(float64(2*lowerboundX-1))); xvelocity-- {
		for yvelocity := xvelocity + 1 - upperboundY; yvelocity >= 1; yvelocity-- {
			yLoc := 0

			for k := 0; ; k++ {
				yLoc += yvelocity - k
				if yLoc <= lowerboundY && yLoc >= upperboundY {
					answer = yvelocity * (yvelocity + 1) / 2
					// answer = yLoc
					fmt.Println(answer)
					return
				}

				if yLoc < upperboundY {
					break
				}

			}

		}
		// for k := 0; k < xvelocity; k++ {
		// 	yvelocity +1
		// 	for yvelocity := 1; yvelocity
		// 	yvelocity := 1
		// 	for {

		// 	}
		// 	yLoc := 0
		// 	for yvelocity := 1; yvelocity-lowerboundY+1 < xvelocity; yvelocity++{
		// 		yvelocity+1
		// 	}
		// 	if
		// 	for 0-lowerboundY + yvelocity+1; 0-upperboundY + yvelocity+1

		// }
	}
	fmt.Println(answer)
}

func validOfXvelocity(x, lowerboundX int) bool {
	maximumDistanceByX := x * (x + 1) / 2
	if maximumDistanceByX < lowerboundX {
		return false
	}

	return true
}

// func
