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

func determineRates(lines []string) (string, string) {
	var countLines int64 = 0
	var arrayOfOnes = make([]int64, len(lines[0]))

	var gammaRate string = ""
	var epsilonRate string = ""

	for _, line := range lines {
		countLines++
		var chars = []byte(line)
		for i := 0; i < len(chars); i++ {
			if chars[i] == '1' {
				arrayOfOnes[i]++
			}
		}
	}

	for i := 0; i < len(arrayOfOnes); i++ {
		if (arrayOfOnes[i] == countLines/2) && (countLines%2 == 0) {
			gammaRate = gammaRate + "1"
			epsilonRate = epsilonRate + "0"
			continue
		}
		if arrayOfOnes[i] > (countLines / 2) {
			gammaRate = gammaRate + "1"
			epsilonRate = epsilonRate + "0"
		} else {
			gammaRate = gammaRate + "0"
			epsilonRate = epsilonRate + "1"
		}
	}
	return gammaRate, epsilonRate
}

func task1(fileName string) {

	var lines = readCompleteFile(fileName)

	gammaRate, epsilonRate := determineRates(lines)

	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Printf("Gamma rate is %v - %v, epsilon rate is %v - %v, product - %v\n", gammaRate, gammaRateInt, epsilonRate, epsilonRateInt, gammaRateInt*epsilonRateInt)

}

func task2(fileName string) {
	var lines = readCompleteFile(fileName)

	oxygenGeneratorRating := filterItems(lines, "gamma", 0)
	co2ScrubberRating := filterItems(lines, "epsilon", 0)

	oxygenGeneratorRatingInt, _ := strconv.ParseInt(oxygenGeneratorRating[0], 2, 64)
	co2ScrubberRatingInt, _ := strconv.ParseInt(co2ScrubberRating[0], 2, 64)

	fmt.Printf("O2 rating - %v, co2 rating - %v, product - %v\n", oxygenGeneratorRatingInt, co2ScrubberRatingInt, oxygenGeneratorRatingInt*co2ScrubberRatingInt)
}

func filterItems(lines []string, pattern string, position int) (ret []string) {

	gammaRate, epsilonRate := determineRates(lines)
	var rate string
	switch pattern {
	case "gamma":
		rate = gammaRate
	case "epsilon":
		rate = epsilonRate
	default:
		return
	}

	if len(lines) <= 1 {
		return lines
	}
	if position >= len(rate) {
		return
	}
	for _, s := range lines {
		if len(s) < len(rate) {
			continue
		}
		if s[position] == rate[position] {
			ret = append(ret, s)
		}
	}
	return filterItems(ret, pattern, position+1)
}

func readCompleteFile(fileName string) []string {
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func main() {
	task1("testdata.txt")
	task1("day3-inputdata-1.txt")
	task2("testdata.txt")
	task2("day3-inputdata-1.txt")
}
