package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	// "strconv"
	"sort"
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

func isLineCorrupted(line string) (corrupted bool, illegalValue int) {
	chars := strings.Split(line, "")

	var openBrackets = []string{}

	for _, c := range chars {
		if isOpeningBracket(c) {
			openBrackets = append(openBrackets, c)
			continue
		}
		if isClosingBracketCorrect(c, openBrackets[len(openBrackets)-1]) {
			openBrackets = openBrackets[:len(openBrackets)-1]
			continue
		} else {
			return true, valueOfBracket(c)
		}
	}
	return false, 0
}

func getIncompleteLine(line string) []string {
	chars := strings.Split(line, "")

	var openBrackets = []string{}

	for _, c := range chars {
		if isOpeningBracket(c) {
			openBrackets = append(openBrackets, c)
			continue
		}
		if isClosingBracketCorrect(c, openBrackets[len(openBrackets)-1]) {
			openBrackets = openBrackets[:len(openBrackets)-1]
			continue
		}
	}
	return openBrackets
}

func valueOfBracket(c string) int {
	switch c {
	case ")":
		return 3
	case ">":
		return 25137
	case "]":
		return 57
	case "}":
		return 1197
	default:
		return 0
	}
}

func isOpeningBracket(c string) bool {
	if c == "(" || c == "<" || c == "[" || c == "{" {
		return true
	}
	return false
}

func isClosingBracketCorrect(closing string, opening string) bool {
	if (opening == "(" && closing == ")") || (opening == "<" && closing == ">") || (opening == "[" && closing == "]") || (opening == "{" && closing == "}") {
		return true
	}
	return false
}

func task1(fileName string) {
	lines := readCompleteFile(fileName)
	var sumOfCorruption int = 0
	for _, line := range lines {
		c, v := isLineCorrupted(line)
		if c {
			sumOfCorruption += v
		}
	}
	fmt.Printf("Sum of corruption: %v\n", sumOfCorruption)
}

func finishLine(brackets []string) int {
	var ret = 0
	for i := len(brackets) - 1; i >= 0; i-- {
		ret *= 5
		ret += valueOfBracket2(brackets[i])
	}
	return ret
}

func valueOfBracket2(s string) int {
	switch s {
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	}
	return 0
}

func task2(fileName string) {
	lines := readCompleteFile(fileName)
	var finishedLines []int = []int{}
	for _, line := range lines {
		c, _ := isLineCorrupted(line)
		if c {
			continue
		}
		bracketsToComplete := getIncompleteLine(line)
		finishedLines = append(finishedLines, finishLine(bracketsToComplete))
	}
	sort.Ints(finishedLines)
	fmt.Printf("Middle score of lines: %v\n", finishedLines[len(finishedLines)/2])

}

func main() {
	task1("day10-testdata.txt")
	task1("day10-inputdata.txt")
	task2("day10-testdata.txt")
	task2("day10-inputdata.txt")
}
