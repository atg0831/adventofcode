package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	sc *bufio.Scanner
)

func init() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal("error when open the input.txt file")
	}

	sc = bufio.NewScanner(file)
	// sc.Split(bufio.ScanWords)
}

type position struct {
	hor int
	dep int
}

func (p *position) calculatePostion(command []string) *position {
	condition := command[0]
	num, _ := strconv.Atoi(command[1])

	if condition == "forward" {
		p.hor += num
	} else if condition == "up" {
		p.dep -= num
	} else if condition == "down" {
		p.dep += num
	}

	return p
}

func main() {

	curPosition := &position{
		hor: 0,
		dep: 0,
	}

	for sc.Scan() {
		input := sc.Text()
		if len(input) == 0 {
			break
		}

		command := strings.Split(input, " ")
		curPosition.calculatePostion(command)

		if err := sc.Err(); err != nil {
			log.Println("error when reading input from input.txt file")
			return
		}
	}

	answer := curPosition.hor * curPosition.dep
	fmt.Println(answer)
}
