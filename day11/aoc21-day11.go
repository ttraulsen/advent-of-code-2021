package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"strconv"
	// "sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCompleteFile(fileName string) []string {
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func readMatrix(lines []string) [10][10]int64 {
	var matrix [10][10]int64 = [10][10]int64{}

	for y, l := range lines {
		splitLine := strings.Split(l, "")
		for x, c := range splitLine {
			i, _ := strconv.ParseInt(c, 10, 64)
			matrix[x][y] = i
		}
	}

	return matrix
}

func increaseEnergyLevelOnMatrix(matrix *[10][10]int64) {
	for x, _ := range *matrix {
		for y, _ := range matrix[x] {
			matrix[x][y]++
		}
	}
}

func resetFlashedCellsInMatrix(matrix *[10][10]int64) {
	for x, _ := range *matrix {
		for y, _ := range matrix[x] {
			if matrix[x][y] > 9 {
				matrix[x][y] = 0
			}
		}
	}
}

func flashCellsInMatrix(matrix *[10][10]int64) int64 {
	var flashes int64 = int64(0)
	for x, _ := range *matrix {
		for y, _ := range matrix[x] {
			if matrix[x][y] > 9 && matrix[x][y] < 1000000 {
				flashes++
				flashCell(matrix, x, y)
			}
		}
	}
	return flashes
}

func flashCell(matrix *[10][10]int64, x int, y int) {
	if x > 0 && y > 0 {
		matrix[x-1][y-1]++
	}
	if x > 0 {
		matrix[x-1][y]++
	}
	if x > 0 && y < 9 {
		matrix[x-1][y+1]++
	}
	if y > 0 {
		matrix[x][y-1]++
	}
	if y < 9 {
		matrix[x][y+1]++
	}
	if x < 9 && y > 0 {
		matrix[x+1][y-1]++
	}
	if x < 9 {
		matrix[x+1][y]++
	}
	if x < 9 && y < 9 {
		matrix[x+1][y+1]++
	}
	matrix[x][y] += 1000000
}

func task1(fileName string) {
	lines := readCompleteFile(fileName)
	matrix := readMatrix(lines)

	flashes := int64(0)

	const numberOfSteps int = 100

	for i := 0; i < numberOfSteps; i++ {
		increaseEnergyLevelOnMatrix(&matrix)
		for {
			flashesInStep := flashCellsInMatrix(&matrix)
			if flashesInStep == 0 {
				break
			}
			flashes += flashesInStep
		}
		resetFlashedCellsInMatrix(&matrix)
	}

	fmt.Printf("After %v steps we had %v flashes\n", numberOfSteps, flashes)
}

func task2(fileName string) {
	lines := readCompleteFile(fileName)
	matrix := readMatrix(lines)

	steps := 0

	for {
		steps++
		increaseEnergyLevelOnMatrix(&matrix)
		flashesInStep := int64(0)
		for {
			flashes := flashCellsInMatrix(&matrix)
			if flashes == 0 {
				break
			}
			flashesInStep += flashes
		}
		if flashesInStep >= 100 {
			break
		}
		resetFlashedCellsInMatrix(&matrix)
	}

	fmt.Printf("After %v steps we synchronized flashing\n", steps)
}

func main() {
	task1("day11-testdata.txt")
	task1("day11-inputdata.txt")
	task2("day11-testdata.txt")
	task2("day11-inputdata.txt")
}
