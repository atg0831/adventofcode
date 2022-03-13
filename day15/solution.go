package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ref) https://go.dev/src/container/heap/example_pq_test.go
type liskInfo struct {
	x        int
	y        int
	priority int // cost
	index    int
}

type PriorityQueue []*liskInfo

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	liskInfo := x.(*liskInfo)
	liskInfo.index = n
	*pq = append(*pq, liskInfo)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	liskInfo := old[n-1]
	old[n-1] = nil      // avoid memory leak
	liskInfo.index = -1 // for safety
	*pq = old[0 : n-1]
	return liskInfo
}

var dir = [4][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
var visited [][]bool

func isPossiblePath(x, y int, visited [][]bool) bool {
	maxX := len(visited)
	maxY := len(visited[0])
	if x >= 0 && x < maxX && y >= 0 && y < maxY {
		if !visited[x][y] {
			return true
		}
	}
	return false
}

func minimumRiskPath(startx, starty, goalx, goaly int, riskLevel [][]int, pq *PriorityQueue) int {
	heap.Push(pq, &liskInfo{startx, starty, riskLevel[startx][starty], 0})
	visited[startx][starty] = true
	for pq.Len() > 0 {
		curInfo := heap.Pop(pq).(*liskInfo)
		curx, cury, curcost := curInfo.x, curInfo.y, curInfo.priority
		if curx == goalx && cury == goaly {
			return curcost
		}

		for _, next := range dir {
			nextx := next[0] + curx
			nexty := next[1] + cury

			if isPossiblePath(nextx, nexty, visited) {
				heap.Push(pq, &liskInfo{nextx, nexty, curcost + riskLevel[nextx][nexty], 0})
				visited[nextx][nexty] = true
			}
		}
	}
	return -1
}

func solve_1(riskLevel [][]int) {
	row, col := len(riskLevel), len(riskLevel[0])
	visited = make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	pq := new(PriorityQueue)
	heap.Init(pq)
	answer := minimumRiskPath(0, 0, row-1, col-1, riskLevel, pq)
	fmt.Printf("part1 answer: %d\n", answer-riskLevel[0][0])
}

func extendMap(original, extended [][]int) [][]int {
	for i := 0; i < len(original); i++ {
		for j := 0; j < len(original[0]); j++ {
			extended[i][j] = original[i][j]
		}
	}

	for k := 1; k < 5; k++ {
		for i := k * len(original); i < (k+1)*len(original); i++ {
			for j := 0; j < len(original[0]); j++ {
				addition := extended[i-len(original)][j] + 1
				if addition > 9 {
					addition = 1
				}
				extended[i][j] = addition
			}
		}
	}

	for k := 1; k < 5; k++ {
		for i := 0; i < len(extended); i++ {
			for j := k * len(original[0]); j < (k+1)*len(original[0]); j++ {
				addition := extended[i][j-len(original[0])] + 1
				if addition > 9 {
					addition = 1
				}
				extended[i][j] = addition
			}
		}
	}

	return extended
}

func solve_2(riskLevel [][]int) {
	extended := make([][]int, len(riskLevel)*5)
	for i := 0; i < len(extended); i++ {
		extended[i] = make([]int, len(riskLevel[0])*5)
	}

	extended = extendMap(riskLevel, extended)
	row, col := len(extended), len(extended[0])
	visited = make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	pq := new(PriorityQueue)
	// *pq = append(*pq, &item{-1, -1, math.MaxInt, 0})
	heap.Init(pq)
	answer := minimumRiskPath(0, 0, row-1, col-1, extended, pq)
	fmt.Printf("part2 answer: %d\n", answer-extended[0][0])
}

func initialize(sc *bufio.Scanner) [][]int {
	riskLevel := make([][]int, 0)
	for sc.Scan() {
		line := sc.Text()

		row := make([]int, len(line))
		for idx, l := range line {
			eachPosition, err := strconv.Atoi(string(l))
			if err != nil {
				log.Fatalf("error when converting from string to integer: %v", err)
			}
			row[idx] = eachPosition
		}
		riskLevel = append(riskLevel, row)
	}

	return riskLevel
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	riskLevel := initialize(sc)
	solve_1(riskLevel)
	solve_2(riskLevel)
}
