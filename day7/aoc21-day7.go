package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

var fuelNeeded = map[int64]int64{}
var crabPositions = map[int64]int64{}

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

func task1(fileName string) {
	lines := readCompleteFile(fileName)
	crabs := readInitialSetup(lines[0])
	s := stats.LoadRawData(crabs)
	med, err := stats.Median(s)
	check(err)
	medRounded, err := stats.Round(med, 0)
	check(err)
	medInt := int64(medRounded)

	var fuel int64 = 0

	for _, v := range crabs {
		if v < medInt {
			fuel += medInt - v
		} else {
			fuel += v - medInt
		}
	}

	fmt.Printf("Median - %v, Fuel needed - %v\n", med, fuel)
}

func mapCrabs(input []int64) map[int64]int64 {
	ret := map[int64]int64{}
	for _, v := range input {
		ret[v]++
	}

	//fill blanks
	var max = int64(0)
	for v := range ret {
		if max < v {
			max = v
		}
	}
	for i := int64(0); i < max; i++ {
		_, present := ret[i]

		if !present {
			ret[i] = 0
		}
	}

	return ret
}

func task2(fileName string) {
	lines := readCompleteFile(fileName)
	crabs := readInitialSetup(lines[0])
	crabPositions = mapCrabs(crabs)
	var minimalPosition int64 = -1

	var totalDistance = map[int64]int64{}

	for k := range crabPositions {
		totalDistance[k] = calculateFuelNeeded(k)
		if minimalPosition == -1 {
			minimalPosition = k
		}
		if totalDistance[k] < totalDistance[minimalPosition] {
			minimalPosition = k
		}
	}
	fmt.Printf("Median - %v, Fuel needed - %v\n", minimalPosition, totalDistance[minimalPosition])
}

func calculateFuelNeeded(position int64) int64 {
	var ret = int64(0)
	for k := range crabPositions {
		var distance = int64(0)
		if position < k {
			distance = k - position
		} else {
			distance = position - k
		}
		ret += calculateFuel(distance) * crabPositions[k]
	}
	return ret
}

func calculateFuel(distance int64) (ret int64) {

	value, present := fuelNeeded[distance]

	if present {
		return value
	}

	for i := int64(0); i <= distance; i++ {
		ret += i
	}

	fuelNeeded[distance] = ret
	return
}

func main() {
	task1("day7-testdata.txt")
	task1("day7-inputdata.txt")
	task2("day7-testdata.txt")
	task2("day7-inputdata.txt")

}
