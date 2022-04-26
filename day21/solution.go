package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var endConditionScore = 1000

func readInput() (int, int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	sc.Scan()
	player1 := sc.Text()
	sc.Scan()
	player2 := sc.Text()

	startPosOfPlayer1, err := strconv.Atoi(strings.Split(strings.Split(player1, ":")[1], " ")[1])
	if err != nil {
		log.Fatal(err)
	}

	startPosOfPlayer2, err := strconv.Atoi(strings.Split(strings.Split(player2, ":")[1], " ")[1])
	if err != nil {
		log.Fatal(err)
	}

	return startPosOfPlayer1, startPosOfPlayer2
}

type player struct {
	pos   int
	score int
}

func solution1() {
	posOfPlayer1, posOfPlayer2 := readInput()

	players := []player{{posOfPlayer1, 0}, {posOfPlayer2, 0}}
	fmt.Println(players)

	cnt, turn := 0, 0

	for {
		advanceSpace := 0
		for i := 1; i <= 3; i++ {
			advanceSpace += (cnt + i)
			// diceNum = i
		}
		cnt += 3
		players[turn].pos = caculateAdvance(players[turn].pos, advanceSpace)
		players[turn].score += players[turn].pos

		if players[turn].score >= endConditionScore {

			fmt.Println(players[(turn+1)%2].score * cnt)
			break
		}

		turn = (turn + 1) % 2
	}
}

func caculateAdvance(pos, advanceSpace int) int {
	// advanceSpace = advanceSpace % 10
	// if advanceSpace == 0 {
	// 	advanceSpace = 10
	// }

	pos = (pos + advanceSpace) % 10
	if pos == 0 {
		pos = 10
	}

	return pos
}

func rollCases() map[int]int {
	rolls := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls[i+j+k] += 1
			}
		}
	}
	fmt.Println(rolls)

	return rolls
}

var rolls = rollCases()

func calculateWhowin(players []player, p int) []int {

	if players[0].score >= 21 {
		return []int{1, 0}
	}
	if players[1].score >= 21 {
		return []int{0, 1}
	}

	var universeCnt = []int{0, 0}
	for k, v := range rolls {
		prevPos := players[p].pos
		prevScore := players[p].score
		if p == 0 {
			players[p].pos = caculateAdvance(players[p].pos, k)
			players[p].score += players[p].pos
			winner := calculateWhowin(players, 1)
			universeCnt[0] += v * winner[0]
			universeCnt[1] += v * winner[1]

		} else {
			players[p].pos = caculateAdvance(players[p].pos, k)
			players[p].score += players[p].pos
			winner := calculateWhowin(players, 0)
			universeCnt[0] += v * winner[0]
			universeCnt[1] += v * winner[1]
		}

		players[p].pos = prevPos
		players[p].score = prevScore
	}

	return universeCnt
}

func solution2() {
	posOfPlayer1, posOfPlayer2 := readInput()

	universeCnt := calculateWhowin([]player{{posOfPlayer1, 0}, {posOfPlayer2, 0}}, 0)
	if universeCnt[0] > universeCnt[1] {
		fmt.Println(universeCnt[0])
	} else {
		fmt.Println(universeCnt[1])
	}
}

func main() {
	solution1()
	solution2()
}
