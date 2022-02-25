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

func main() {
	answer := -1
	prev := -1
	for sc.Scan() {
		input := sc.Text()
		if len(input) == 0 {
			break
		}

		cur, _ := strconv.Atoi(input)
		if isLargerThanPrev(prev, cur) {
			answer += 1
		}

		prev = cur
	}

	fmt.Println(answer)
}
