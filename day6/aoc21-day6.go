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

func readCompleteFile(fileName string) []string {
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func readInitialSetup(s string) []int64 {
	var ret []int64 = []int64{}

	tokens := strings.Split(s, ",")
	for _, v := range tokens {
		i, _ := strconv.ParseInt(v, 10, 64)
		ret = append(ret, i)
	}

	return ret
}

func readOptimizedSetup(s string) (ret [9]int64) {
	tokens := strings.Split(s, ",")
	for _, v := range tokens {
		i, _ := strconv.ParseInt(v, 10, 64)
		ret[i]++
	}
	return
}

func task1(fileName string, numberOfDays int) {
	lines := readCompleteFile(fileName)
	lanternFishes := readInitialSetup(lines[0])

	for i := 0; i < numberOfDays; i++ {
		var newFishes = []int64{}
		var appendedFishes = []int64{}
		for _, v := range lanternFishes {
			if v == 0 {
				newFishes = append(newFishes, 6)
				appendedFishes = append(appendedFishes, 8)
			} else {
				newFishes = append(newFishes, v-1)
			}
		}
		newFishes = append(newFishes, appendedFishes...)
		lanternFishes = newFishes
	}
	fmt.Printf("# Fishes after %v cycles: %v\n", numberOfDays, len(lanternFishes))
}

func task2(fileName string, numberOfDays int) {
	lines := readCompleteFile(fileName)
	fishes := readOptimizedSetup(lines[0])
	fmt.Printf("# initial fishes %v\n", fishes)
	for i := 0; i < numberOfDays; i++ {
		fishes = processDay(fishes)
	}
	fmt.Printf("# afterwards fishes %v\n", sumState(fishes))
}

func processDay(oldState [9]int64) (newState [9]int64) {
	for i := 8; i >= 0; i-- {
		if i == 0 {
			newState[6] += oldState[0]
			newState[8] += oldState[0]
		} else {
			newState[i-1] += oldState[i]
		}
	}
	return
}

func sumState(state [9]int64) (ret int64) {
	for i := 0; i < 9; i++ {
		ret += state[i]
	}
	return
}

func main() {
	task1("day6-testdata.txt", 80)
	task1("day6-inputdata.txt", 80)

	task2("day6-testdata.txt", 80)
	task2("day6-inputdata.txt", 80)
	task2("day6-testdata.txt", 256)
	task2("day6-inputdata.txt", 256)

}
