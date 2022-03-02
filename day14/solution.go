package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()
	template := sc.Text()

	patternMap := make(map[string]string)

	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}

		rule := strings.Split(line, " -> ")
		src, in := rule[0], rule[1]
		patternMap[src] = in
	}

	elementCnt := make(map[string]int)
	for _, char := range template {
		elementCnt[string(char)] += 1
	}

	polymer := template
	step := 1
	for step < 41 {
		var temp string
		for i := 0; i < len(polymer)-1; i++ {
			src := string(polymer[i]) + string(polymer[i+1])
			temp += string(src[0]) + patternMap[src] // 기존의 왼쪽 문자와 새로 insert 되는 문자 합쳐서 문자열 만들기
			elementCnt[patternMap[src]] += 1
		}
		polymer = temp + string(polymer[len(polymer)-1]) // 마지막 문자 붙여주기
		step += 1
	}

	maxCnt := -1
	minCnt := len(polymer) + 1
	for k, v := range elementCnt {
		fmt.Println(k, v, maxCnt, minCnt)
		if v > maxCnt {
			maxCnt = v
		}

		if v < minCnt {
			minCnt = v
		}
	}

	fmt.Println(maxCnt - minCnt)
}
