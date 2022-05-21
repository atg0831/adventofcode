package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput() [][]string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	cucumberMap := make([][]string, 0)

	// rowLen := len(line)
	for sc.Scan() {
		line := sc.Text()
		cucumberMap = append(cucumberMap, strings.Split(line, ""))
	}
	return cucumberMap
}

type updateField struct {
	row  int
	col  int
	move string
}

func main() {
	cucumberMap := readInput()

	updateStorage := make([]updateField, 0)
	step := 1
	for {
		var movable = false
		for row := 0; row < len(cucumberMap); row++ {
			for col := 0; col < len(cucumberMap[row])-1; col++ {
				if cucumberMap[row][col] == ">" && cucumberMap[row][col+1] == "." {
					updateStorage = append(updateStorage, updateField{row, col, "."})
					updateStorage = append(updateStorage, updateField{row, col + 1, ">"})
					movable = true
				}

			}
			if cucumberMap[row][len(cucumberMap[row])-1] == ">" && cucumberMap[row][0] == "." {
				updateStorage = append(updateStorage, updateField{row, len(cucumberMap[row]) - 1, "."})
				updateStorage = append(updateStorage, updateField{row, 0, ">"})
				movable = true
			}
		}

		for i := 0; i < len(updateStorage); i++ {
			cucumberMap[updateStorage[i].row][updateStorage[i].col] = updateStorage[i].move
		}

		updateStorage = make([]updateField, 0)

		for row := 0; row < len(cucumberMap)-1; row++ {
			for col := 0; col < len(cucumberMap[row]); col++ {
				if cucumberMap[row][col] == "v" && cucumberMap[row+1][col] == "." {
					updateStorage = append(updateStorage, updateField{row, col, "."})
					updateStorage = append(updateStorage, updateField{row + 1, col, "v"})
					movable = true
				}

			}
		}

		for col := 0; col < len(cucumberMap[0]); col++ {
			if cucumberMap[len(cucumberMap)-1][col] == "v" && cucumberMap[0][col] == "." {
				updateStorage = append(updateStorage, updateField{len(cucumberMap) - 1, col, "."})
				updateStorage = append(updateStorage, updateField{0, col, "v"})
				movable = true
			}
		}

		if !movable {
			break
		}

		for i := 0; i < len(updateStorage); i++ {
			cucumberMap[updateStorage[i].row][updateStorage[i].col] = updateStorage[i].move
		}
		updateStorage = make([]updateField, 0)

		step += 1
	}

	fmt.Println(step)
}
