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

func getMaxYVelocity(lowerboundY, upperboundY int) int {
	if lowerboundY+upperboundY > 0 {
		return upperboundY
	}

	return (-1 * lowerboundY) - 1
}

func sunOfOneToNum(x int) int {
	return x * (x + 1) / 2
}

func inTargetRange(x, y, lowerboundX, upperboundX, lowerboundY, upperboundY int) bool {
	return x >= lowerboundX && x <= upperboundX && y >= lowerboundY && y <= upperboundY
}

func outOfTargetRange(x, y, upperboundX, lowerboundY int) bool {
	return x > upperboundX || y < lowerboundY
}
func solution1(lowerboundY, upperboundY int) {
	answer := sunOfOneToNum(getMaxYVelocity(lowerboundY, upperboundY))
	fmt.Println(answer)
}

func solution2(lowerboundX, upperboundX, lowerboundY, upperboundY int) {
	answer := 0
	for i := int(math.Sqrt(float64(lowerboundX))); i <= upperboundX; i++ {
		for j := -getMaxYVelocity(lowerboundY, upperboundY) - 1; j < getMaxYVelocity(lowerboundY, upperboundY)+1; j++ {
			xvelocity, yvelocity := i, j
			xLoc, yLoc := 0, 0
			for {
				xLoc += xvelocity
				yLoc += yvelocity

				if outOfTargetRange(xLoc, yLoc, upperboundX, lowerboundY) {
					break
				}

				if inTargetRange(xLoc, yLoc, lowerboundX, upperboundX, lowerboundY, upperboundY) {
					answer += 1
					break
				}

				yvelocity -= 1
				if xvelocity > 0 {
					xvelocity -= 1
				} else if xvelocity < 0 {
					xvelocity += 1
				}
			}
		}
	}
	fmt.Println(answer)
}

func main() {
	targetXRange, targetYRange := readInput()

	lowerboundX := convertToInteger(targetXRange[0])
	upperboundX := convertToInteger(targetXRange[1])

	lowerboundY := convertToInteger(targetYRange[0])
	upperboundY := convertToInteger(targetYRange[1])

	solution1(lowerboundY, upperboundY)
	solution2(lowerboundX, upperboundX, lowerboundY, upperboundY)
}
