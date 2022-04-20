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

func calculateAddition(curPos, totalScore, totalDiceNum int) (int, int) {
	wrapBackTotalDiceNum := totalDiceNum % 10
	if wrapBackTotalDiceNum == 0 {
		wrapBackTotalDiceNum = 10
	}
	wrapBackSpace := (curPos + wrapBackTotalDiceNum) % 10
	if wrapBackSpace == 0 {
		wrapBackSpace = 10
	}

	curPos = (curPos + wrapBackTotalDiceNum) % 10
	if curPos == 0 {
		curPos = 10
	}

	totalScore = totalScore + wrapBackSpace
	return totalScore, curPos
}
func solution1() {
	posOfPlayer1, posOfPlayer2 := readInput()

	answer := 0
	player1Score, player2Score := 0, 0
	for i := 1; ; i += 6 {
		player1Score, posOfPlayer1 = calculateAddition(posOfPlayer1, player1Score, 3*(i+1))
		fmt.Println(player1Score, i, posOfPlayer1)
		if player1Score >= endConditionScore {
			answer = (i + 2) * player2Score
			break
		}
		player2Score, posOfPlayer2 = calculateAddition(posOfPlayer2, player2Score, 3*(i+4))
		fmt.Println(player2Score, i, posOfPlayer2)
		if player2Score >= endConditionScore {
			answer = (i + 2) * player1Score
			break
		}
	}
	fmt.Println(answer)
}

func main() {
	solution1()
}
