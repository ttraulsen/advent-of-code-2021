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

type bingoBoard struct {
	board         [5][5]int64
	markedNumbers []int64
	bingo         bool
}

func (b *bingoBoard) setBoard(board [5][5]int64) {
	b.board = board
}

func (b *bingoBoard) markNumber(n int64) {
	b.markedNumbers = append(b.markedNumbers, n)
}

func (b *bingoBoard) checkForBingo() bool {
	if b.bingo {
		return true
	}
	if len(b.markedNumbers) < 5 {
		return false
	}
	for i := 0; i < 5; i++ {
		horizontal := true
		vertical := true
		for j := 0; j < 5; j++ {
			if !b.isNumberAlreadyMarked(b.board[i][j]) {
				horizontal = false
			}
			if !b.isNumberAlreadyMarked(b.board[j][i]) {
				vertical = false
			}
		}
		if horizontal || vertical {
			b.bingo = true
			return true
		}
	}
	return false
}

func (b *bingoBoard) isNumberAlreadyMarked(e int64) bool {
	for _, a := range b.markedNumbers {
		if a == e {
			return true
		}
	}
	return false
}

func (b *bingoBoard) sumOfUnmarkedNumbers() (ret int64) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.isNumberAlreadyMarked(b.board[i][j]) {
				ret += b.board[i][j]
			}
		}
	}
	return
}

func readCompleteFile(fileName string) []string {
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func readBingoFile(fileName string) (numbers []int64, bingoBoards []*bingoBoard) {
	lines := readCompleteFile(fileName)
	var numbersToProcess []int64
	var boards []*bingoBoard

	for _, n := range strings.Split(lines[0], ",") {
		i, _ := strconv.ParseInt(n, 10, 64)
		numbersToProcess = append(numbersToProcess, i)
	}

	for i := 2; i < len(lines); i = i + 6 {
		boards = append(boards, readBingoBoard(lines, i))
	}

	return numbersToProcess[:], boards[:]
}

func task1(fileName string) {
	numbersToProcess, boards := readBingoFile(fileName)
	var winnerBoard *bingoBoard = nil
	for _, n := range numbersToProcess {
		for _, b := range boards {
			b.markNumber(n)
			if b.checkForBingo() {
				winnerBoard = b
				break
			}
			if winnerBoard != nil {
				break
			}
		}
		if winnerBoard != nil {
			break
		}
	}
	s := winnerBoard.sumOfUnmarkedNumbers()
	l := winnerBoard.markedNumbers[len(winnerBoard.markedNumbers)-1]
	fmt.Printf("We got a winning board %v - Sum %v - last number %v - product %v\n", winnerBoard, s, l, s*l)
}

func task2(fileName string) {
	numbersToProcess, boards := readBingoFile(fileName)
	var markedBoards []*bingoBoard = nil
	for _, n := range numbersToProcess {
		for _, b := range boards {
			if b.checkForBingo() {
				continue
			}
			b.markNumber(n)
			if b.checkForBingo() {
				markedBoards = append(markedBoards, b)
			}
		}
		if len(markedBoards) >= len(boards) {
			break
		}
	}
	loosingBoard := markedBoards[len(markedBoards)-1]
	s := loosingBoard.sumOfUnmarkedNumbers()
	l := loosingBoard.markedNumbers[len(loosingBoard.markedNumbers)-1]
	fmt.Printf("We got a loosing board %v - Sum %v - last number %v - product %v\n", loosingBoard, s, l, s*l)
}

func readBingoBoard(lines []string, position int) *bingoBoard {
	ret := [5][5]int64{}

	for i := 0; i <= 4; i++ {
		cleanedString := strings.TrimPrefix(strings.ReplaceAll(lines[position+i], "  ", " "), " ")
		numbers := strings.Split(cleanedString, " ")
		for j := 0; j < len(numbers); j++ {
			ret[i][j], _ = strconv.ParseInt(numbers[j], 10, 64)
		}
	}
	b := bingoBoard{ret, []int64{}, false}
	return &b
}

func main() {
	task1("day4/testdata.txt")
	task1("day4/day4-inputdata-1.txt")
	task2("day4/testdata.txt")
	task2("day4/day4-inputdata-1.txt")
}
