package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
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

func splitLine(s string) ([]string, []string) {
	divided := strings.Split(s, " | ")
	testinput := strings.Split(divided[0], " ")
	digits := strings.Split(divided[1], " ")
	return testinput, digits
}

func task1(fileName string) {
	var uniqueDigitNumbers = int64(0)

	lines := readCompleteFile(fileName)

	uniques := [4]int{2, 3, 4, 7}

	for _, l := range lines {
		if len(l) < 2 {
			continue
		}
		_, digits := splitLine(l)
		for _, d := range digits {
			if containsInt(uniques[:], len(d)) {
				uniqueDigitNumbers++
			}
		}
	}
	fmt.Printf("Found %v times a unique-digit element\n", uniqueDigitNumbers)
}

func containsInt(arr []int, e int) bool {
	for _, a := range arr {
		if a == e {
			return true
		}
	}
	return false
}

func containsString(arr []string, e string) bool {
	for _, a := range arr {
		if a == e {
			return true
		}
	}
	return false
}

func determineNumbers(input []string) map[string]string {
	var twoDigits = ""
	var fourDigits = ""
	var fiveDigits = []string{}
	var sixDigits = []string{}
	var sixOrNine = []string{}
	var ret = make(map[string]string)

	for _, v := range input {
		var sorted = sortString(v)
		switch len(v) {
		case 2:
			twoDigits = sorted
			ret[sorted] = "1"
		case 3:
			ret[sorted] = "7"
		case 4:
			fourDigits = sorted
			ret[sorted] = "4"
		case 5:
			fiveDigits = append(fiveDigits, sorted)
		case 6:
			sixDigits = append(sixDigits, sorted)
		case 7:
			ret[sorted] = "8"
		}
	}

	var missingSegmentsOfFive = ""

	for _, v := range sixDigits {
		if isItTheNull(twoDigits, fourDigits, v) {
			ret[v] = "0"
		} else {
			sixOrNine = append(sixOrNine, v)
			missingSegmentsOfFive += findMissingSegments(v)
		}
	}

	fiveString := findMissingSegments(missingSegmentsOfFive)
	ret[fiveString] = "5"

	if isItTheNine(twoDigits, sixOrNine[0]) {
		ret[sixOrNine[0]] = "9"
		ret[sixOrNine[1]] = "6"
	} else {
		ret[sixOrNine[0]] = "6"
		ret[sixOrNine[1]] = "9"
	}

	var threeOrTwo = []string{}

	for _, v := range fiveDigits {
		if v != fiveString {
			threeOrTwo = append(threeOrTwo, v)
		}
	}

	if isItTheTree(twoDigits, threeOrTwo[0]) {
		ret[threeOrTwo[0]] = "3"
		ret[threeOrTwo[1]] = "2"
	} else {
		ret[threeOrTwo[0]] = "2"
		ret[threeOrTwo[1]] = "3"

	}

	return ret
}

func task2(fileName string) {
	lines := readCompleteFile(fileName)

	var overAllSum int64 = int64(0)

	for _, l := range lines {
		if len(l) < 2 {
			continue
		}
		probes, digits := splitLine(l)
		checkStrings := determineNumbers(probes)
		outputString := ""
		for _, d := range digits {
			outputString += checkStrings[sortString(d)]
		}
		i, _ := strconv.ParseInt(outputString, 10, 64)
		overAllSum += i
	}
	fmt.Printf("Overall sum of inputs: %v\n", overAllSum)
}

func isItTheTree(one string, probe string) bool {
	oneSegments := strings.Split(one, "")
	checkSegments := strings.Split(probe, "")

	for _, v := range oneSegments {
		if !containsString(checkSegments, v) {
			return false
		}
	}
	return true
}

func isItTheNull(one string, four string, probe string) bool {
	missingSegments := strings.Split(findMissingSegments(probe), "")
	fourSegments := strings.Split(four, "")
	oneSegments := strings.Split(one, "")
	checkSegments := []string{}

	for _, v := range fourSegments {
		if !containsString(oneSegments, v) {
			checkSegments = append(checkSegments, v)
		}
	}
	for _, v := range missingSegments {
		if containsString(checkSegments, v) {
			return true
		}
	}
	return false
}

func isItTheNine(one string, probe string) bool {
	missingSegments := strings.Split(findMissingSegments(probe), "")
	onesSegments := strings.Split(one, "")
	for _, v := range missingSegments {
		if containsString(onesSegments, v) {
			return false
		}
	}
	return true
}

func findMissingSegments(input string) string {
	s := strings.Split(input, "")
	var ret = ""
	allSegments := strings.Split("abcdefg", "")
	for _, v := range allSegments {
		if !containsString(s, v) {
			ret += v
		}
	}
	return ret
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	task1("day8-testdata.txt")
	task1("day8-inputdata.txt")
	task2("day8-testdata.txt")
	task2("day8-inputdata.txt")
}
