package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

type scannerStack []int

func (s *scannerStack) push(element int) {
	*s = append(*s, element)
}

func (s *scannerStack) pop() int {
	top := len(*s) - 1
	element := (*s)[top]
	*s = (*s)[:top]
	return element
}

type coordinates struct {
	x, y, z int
}
type scanner struct {
	beacons  []coordinates
	position coordinates
}

func readInput() []scanner {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scanners []scanner
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		var scanner scanner
		if strings.Contains(line, "---") {
			for sc.Scan() {
				line := sc.Text()
				if len(line) == 0 {
					break
				}
				reader := strings.NewReader(line)
				var beacon coordinates
				_, err := fmt.Fscanf(reader, "%d,%d,%d", &beacon.x, &beacon.y, &beacon.z)
				if err != nil {
					log.Fatal(err)
				}
				scanner.beacons = append(scanner.beacons, coordinates{beacon.x, beacon.y, beacon.z})
			}
		}
		scanners = append(scanners, scanner)
	}

	// scanner 0의 좌표를 0, 0, 0으로 초기화
	scanners[0].position = coordinates{0, 0, 0}
	return scanners
}

func (c coordinates) allOrientations(i int) coordinates {
	allOrientations := [24]coordinates{
		{c.x, c.y, c.z}, {-c.x, -c.y, c.z}, {c.y, -c.x, c.z}, {-c.y, c.x, c.z},
		{c.y, c.z, c.x}, {-c.y, -c.z, c.x}, {c.z, -c.y, c.x}, {-c.z, c.y, c.x},
		{c.z, c.x, c.y}, {-c.z, -c.x, c.y}, {c.x, -c.z, c.y}, {-c.x, c.z, c.y},
		{c.y, c.x, -c.z}, {-c.y, -c.x, -c.z}, {c.x, -c.y, -c.z}, {-c.x, c.y, -c.z},
		{c.x, c.z, -c.y}, {-c.x, -c.z, -c.y}, {c.z, -c.x, -c.y}, {-c.z, c.x, -c.y},
		{c.z, c.y, -c.x}, {-c.z, -c.y, -c.x}, {c.y, -c.z, -c.x}, {-c.y, c.z, -c.x},
	}

	return allOrientations[i]
}

func getRotatedScanner(s scanner, r int) scanner {
	rotatedScanner := scanner{}
	for _, beacon := range s.beacons {
		rotatedBeacon := beacon.allOrientations(r)
		rotatedScanner.beacons = append(rotatedScanner.beacons, rotatedBeacon)
	}

	return rotatedScanner
}

func getOverlappedBeacons(pivot, target scanner) int {
	pivotBeacons := pivot.beacons
	targetBeacons := target.beacons

	overlapped := 0
	for p := 0; p < len(pivotBeacons); p++ {
		for t := 0; t < len(targetBeacons); t++ {
			if pivotBeacons[p] == targetBeacons[t] {
				overlapped++
				break
			}
		}
	}

	return overlapped
}

func getRelativeScanner(pivotScanner, targetScanner scanner) (scanner, bool) {
	for i := 0; i < len(pivotScanner.beacons); i++ {
		pivotBeacon := pivotScanner.beacons[i]
		for r := 0; r < 24; r++ {
			rotatedScanner := getRotatedScanner(targetScanner, r)
			for j := 0; j < len(rotatedScanner.beacons); j++ {
				targetBeacon := rotatedScanner.beacons[j]

				distx, disty, distz := pivotBeacon.x-targetBeacon.x, pivotBeacon.y-targetBeacon.y, pivotBeacon.z-targetBeacon.z
				movedBeacons := make([]coordinates, len(rotatedScanner.beacons))
				// pivot beacon 위치와 target beacon 위치의 차이만큼 targetScanner의 beacons 들을
				// 이동시키고 이동시킨 scanner의 beacons와 pivot scanner의 beacons가 겹치는 부분이 12개 이상인지 판별
				for k := 0; k < len(rotatedScanner.beacons); k++ {
					curBeacon := rotatedScanner.beacons[k]
					movedBeacon := coordinates{curBeacon.x + distx, curBeacon.y + disty, curBeacon.z + distz}
					movedBeacons[k] = movedBeacon
				}
				movedScanner := scanner{beacons: movedBeacons, position: coordinates{distx, disty, distz}}

				if getOverlappedBeacons(pivotScanner, movedScanner) >= 12 {
					return movedScanner, true
				}
			}
		}
	}

	return scanner{}, false
}

func solution() {
	scanners := readInput()
	beacons := mapset.NewSet()
	visited := make([]bool, len(scanners))
	scannerStack := scannerStack{}
	scannerStack.push(0)
	visited[0] = true

	for len(scannerStack) > 0 {
		scannerIdx := scannerStack.pop()
		pivotScanner := scanners[scannerIdx]

		// 전체 beacon에 중복없이 add
		for i := 0; i < len(pivotScanner.beacons); i++ {
			beacons.Add(pivotScanner.beacons[i])
		}

		for i, targetScanner := range scanners {
			if !visited[i] {
				// 24방향의 rotation 경우 loop 돌면서 기준 beacon과의 거리만큼 이동시켜서 overlapped 되는 부분 있는지 파악
				relativeScanner, ok := getRelativeScanner(pivotScanner, targetScanner)
				if ok {
					// log.Println(i, movedScanner)
					scanners[i] = relativeScanner
					scannerStack.push(i)
					visited[i] = true
				}
			}
		}
	}
	fmt.Println(beacons.Cardinality())

	maximum := -int(math.MaxInt)
	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			distance := calculateManhattanDist(scanners[i], scanners[j])
			if distance > maximum {
				maximum = distance
			}
		}
	}

	fmt.Println(maximum)
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func calculateManhattanDist(leftScanner, rightScanner scanner) int {
	distance := 0

	distance += Abs(leftScanner.position.x - rightScanner.position.x)
	distance += Abs(leftScanner.position.y - rightScanner.position.y)
	distance += Abs(leftScanner.position.z - rightScanner.position.z)

	return distance
}

func main() {
	solution()
}
