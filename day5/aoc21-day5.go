package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type coordinates struct {
	x int64
	y int64
}

var board [1000][1000]int64

func readCompleteFile(fileName string) []string {
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func task1(fileName string) {
	resetBoard()
	lines := readCompleteFile(fileName)
	var crossedPoints int64 = 0
	for _, line := range lines {
		processLine(line, true)
	}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if board[x][y] > 1 {
				crossedPoints++
			}
		}
	}
	fmt.Printf("Number of crossing points found: %v\n", crossedPoints)
}

func task2(fileName string) {
	resetBoard()
	lines := readCompleteFile(fileName)
	var crossedPoints int64 = 0
	for _, line := range lines {
		processLine(line, false)
	}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if board[x][y] > 1 {
				crossedPoints++
			}
		}
	}

	fmt.Printf("Number of crossing points found: %v\n", crossedPoints)
}

func processLine(line string, ignoreDiagonal bool) {
	if len(line) < 2 {
		return
	}
	s, e := splitLine(line)
	var lower coordinates
	var upper coordinates
	//vertical line
	if s.x == e.x {
		if s.y < e.y {
			lower = s
			upper = e
		} else {
			lower = e
			upper = s
		}
		for i := lower.y; i <= upper.y; i++ {
			board[lower.x][i]++
		}
		return
	}

	//horizontal line
	if s.y == e.y {
		if s.x < e.x {
			lower = s
			upper = e
		} else {
			lower = e
			upper = s
		}
		for i := lower.x; i <= upper.x; i++ {
			board[i][lower.y]++
		}
		return
	}

	if ignoreDiagonal {
		return
	}
	//diagonal line
	//not yet implemented
	if s.x < e.x {
		lower = s
		upper = e
	} else {
		lower = e
		upper = s
	}
	for i := int64(0); i <= upper.x-lower.x; i++ {
		if lower.y < upper.y {
			board[lower.x+i][lower.y+i]++
		} else {
			board[lower.x+i][lower.y-i]++
		}
	}
}

func splitLine(line string) (startingPoint coordinates, endingPoint coordinates) {
	if len(line) < 2 {
		return
	}
	pointStrings := strings.Split(line, " -> ")
	startingPoint = readPoint(pointStrings[0])
	endingPoint = readPoint(pointStrings[1])
	return
}

func readPoint(s string) (point coordinates) {
	c := strings.Split(s, ",")
	point.x, _ = strconv.ParseInt(c[0], 10, 64)
	point.y, _ = strconv.ParseInt(c[1], 10, 64)
	return
}

func resetBoard() {
	board = [1000][1000]int64{}
}

func main() {
	task1("day5-testdata.txt")
	task1("day5-inputdata.txt")
	task2("day5-testdata.txt")
	task2("day5-inputdata.txt")
}
