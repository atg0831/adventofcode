package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	dotCoordinate := make([][]int, 0)
	var row, col int // row, col은 code의 range

	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			for sc.Scan() {
				overlapped := 0 // 중복되는 자표 카운트 위한 변수
				condition := strings.Split((strings.Split(sc.Text(), " "))[2], "=")
				foldPivot, _ := strconv.Atoi(condition[1])

				if condition[0] == "y" { // fold y=...
					doFolding(dotCoordinate, foldPivot, 1)
					row = foldPivot
				} else { // fold x=...
					doFolding(dotCoordinate, foldPivot, 0)
					col = foldPivot
				}

				// 같은 level 순서대로(같은 y 위치의 좌표들 순서로 소팅(같은 y에 있다면 x 좌표가 작은 것부터 소팅))
				sort.Slice(dotCoordinate, func(i, j int) bool {
					if dotCoordinate[i][1] == dotCoordinate[j][1] {
						return dotCoordinate[i][0] < dotCoordinate[j][0]
					}
					return dotCoordinate[i][1] < dotCoordinate[j][1]
				})

				// prev와 현재 cur과 같다면 겹쳐지는 것
				prev := []int{-1, -1}
				for idx, cur := range dotCoordinate {
					if prev[0] == cur[0] && prev[1] == cur[1] {
						overlapped += 1
						dotCoordinate[idx][0] = -1
						dotCoordinate[idx][1] = -1
					}
					prev = cur
				}

				answer := len(dotCoordinate) - overlapped
				fmt.Printf("fold along %s=%d -> %d\n", condition[0], foldPivot, answer)
			}
		} else {
			coordinate := strings.Split(line, ",")
			x, _ := strconv.Atoi(coordinate[0])
			y, _ := strconv.Atoi(coordinate[1])

			dotCoordinate = append(dotCoordinate, []int{x, y})
		}
	}

	// row, col는 마지막 folding 크기
	background := make([][]int, row)
	for i := 0; i < row; i++ {
		background[i] = make([]int, col)
	}

	makeCode(background, dotCoordinate, row, col)
}

func doFolding(dotCoordinate [][]int, foldPivot, axis int) {
	for k := 0; k < len(dotCoordinate); k++ {
		val := dotCoordinate[k][axis]
		if val < foldPivot {
			continue
		}
		dotCoordinate[k][axis] -= 2 * (val - foldPivot) // folding point와 현재 좌표값을 뺀 것의 2배하면 접었을 때의 좌표로 이동됨
	}
}

func makeCode(background, dotCoordinate [][]int, row, col int) {
	// dot == -1 -> 중복되는 부분
	for _, dot := range dotCoordinate {
		if dot[0] == -1 || dot[1] == -1 {
			continue
		}

		background[dot[1]][dot[0]] = 1
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if background[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
