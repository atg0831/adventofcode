package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	// sc.Split(bufio.ScanWords)
}

func isLargerThanPrev(prev, cur int) bool {
	if prev < cur {
		return true
	}

	return false
}

type slidingWindow struct {
	a, b, c int
}

func sumSlidingWindow(s slidingWindow) int {
	total := s.a + s.b + s.c
	return total
}

func main() {
	answer := 0
	prev := -1
	i := 0
	j := 1
	numSlice := make([]int, 0)

	for sc.Scan() {
		input := sc.Text()
		if len(input) == 0 {
			break
		}
		measurement, _ := strconv.Atoi(input)
		numSlice = append(numSlice, measurement)

		i += 1

		if i == 3 {
			prev = sumSlidingWindow(slidingWindow{numSlice[0], numSlice[1], numSlice[2]})
		} else if i > 3 {
			cur := sumSlidingWindow(slidingWindow{numSlice[j], numSlice[j+1], numSlice[j+2]})
			if isLargerThanPrev(prev, cur) {
				answer += 1
			}

			prev = cur
			j += 1
		}

	}

	fmt.Println(answer)
}
