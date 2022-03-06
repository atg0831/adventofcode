package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	rules := make(map[string]string)
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}

		pattern := strings.Split(line, " -> ")
		rules[pattern[0]] = pattern[1]
	}

	solution1(rules, template)
	solution2(rules, template)
}

func solution2(rules map[string]string, template string) {
	elementCnt := make(map[string]int)
	for _, char := range template {
		elementCnt[string(char)] += 1
	}

	pairCnt := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairCnt[string(template[i])+string(template[i+1])] += 1
	}

	var step = 1
	for step = 1; step < 41; step++ {
		pairCntByStep := make(map[string]int) // 각 step에서의 pair의 개수를 카운팅하기 위함
		for k, v := range pairCnt {
			first := string(k[0]) + rules[k]
			second := rules[k] + string(k[1])
			pairCntByStep[first] += v  // 해당되는 k 값의 개수만큼 first, second의 개수를 누적해야 됨 template -> sdsdffgkasgsdff 이고 k = sd, rules[sd] = a;
			pairCntByStep[second] += v // first = sa, secod = ad 결국 sd의 개수만큼 first, second가 누적됨
			elementCnt[rules[k]] += v  // template -> 'anbbcdankddsansdfan' 이고 현재 k 값이 'an' 이면 'an'에 해당되는 문자는 rules[an] 값으로 다 바뀌므로 v를 누적해준다(v는 결국 직전 step에서의 'an'의 개수)
		}

		pairCnt = pairCntByStep
	}

	maxCnt := -1
	minCnt := math.MaxInt64
	for _, v := range elementCnt {
		if v > maxCnt {
			maxCnt = v
		}

		if v < minCnt {
			minCnt = v
		}
	}
	fmt.Printf("part2\nstep: %d, 정답: %d\n", step-1, maxCnt-minCnt)
}

func solution1(rules map[string]string, template string) {
	elementCnt := make(map[string]int)
	for _, char := range template {
		elementCnt[string(char)] += 1
	}

	var step = 1
	polymer := template
	for step = 1; step < 11; step++ {
		var temp string
		for i := 0; i < len(polymer)-1; i++ {
			src := string(polymer[i]) + string(polymer[i+1])
			temp += string(src[0]) + rules[src] // 기존의 왼쪽 문자와 새로 insert 되는 문자 합쳐서 문자열 만들기
			elementCnt[rules[src]] += 1
		}
		polymer = temp + string(polymer[len(polymer)-1]) // 마지막 문자 붙여주기
	}

	maxCnt := -1
	minCnt := len(polymer) + 1
	for _, v := range elementCnt {
		if v > maxCnt {
			maxCnt = v
		}

		if v < minCnt {
			minCnt = v
		}
	}

	fmt.Printf("part1\nstep: %d, 정답: %d\n", step-1, maxCnt-minCnt)
}
