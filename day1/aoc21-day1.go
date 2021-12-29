package main

import (
	"bufio"
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

func task1(fileName string) {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + "/" + fileName)
	check(err)

	scanner := bufio.NewScanner(f)

	var lastLine int64 = -1
	var currentLine int64 = -1
	var increments = 0

	for scanner.Scan() {
		var line = scanner.Text()
		currentLine, _ = strconv.ParseInt(line, 10, 64)
		//fmt.Println(line)
		if lastLine != -1 {
			if lastLine < currentLine {
				increments++
			}
		}
		lastLine = currentLine
	}
	fmt.Printf("Number of increments in %v: %v\n", fileName, increments)

}

func readCompleteFile(fileName string) []string {
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func task2(fileName string) {
	var lines = readCompleteFile(fileName)
	var increments = 0
	for i := 0; i < len(lines)-3; i++ {
		if lines[i+3] > lines[i] {
			increments++
		}
	}
	fmt.Printf("Number of measured increments in %v: %v\n", fileName, increments)
}

func main() {
	task1("testdata.txt")
	task1("day1-inputdata-1.txt")
	task2("testdata.txt")
	task2("day1-inputdata-1.txt")
}
