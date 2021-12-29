package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	// "strconv"
	// "sort"
)

var primes [20]int = [20]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71}
var currentPrimeIndex = 0

type cave struct {
	name           string
	connectedCaves []*cave
	prime          int
}

func (me *cave) connectCave(otherCave *cave) {
	if !me.isConnected(otherCave) {
		me.connectedCaves = append(me.connectedCaves, otherCave)
		otherCave.connectCave(me)
	}
}

func (me *cave) isConnected(otherCave *cave) bool {
	for _, c := range me.connectedCaves {
		if c.name == otherCave.name {
			return true
		}
	}
	return false
}

type path struct {
	passedCaves  []*cave
	alreadyTwice bool
	pathPrimes   int
}

func (me *path) isMovingPossible(nextCave *cave) bool {
	if strings.ToUpper(nextCave.name) == nextCave.name {
		return true
	}
	if me.pathPrimes%nextCave.prime == 0 {
		return false
	}
	return true
}

func (me *path) isMovingPossibleSmallTwice(nextCave *cave) bool {
	if strings.ToUpper(nextCave.name) == nextCave.name {
		return true
	}
	if nextCave.name == "start" {
		return false
	}
	if me.pathPrimes%nextCave.prime == 0 && me.alreadyTwice {
		return false
	}

	return true
}

func (me *path) nextCave(nextCave *cave) {
	me.passedCaves = append(me.passedCaves, nextCave)
	if me.pathPrimes%nextCave.prime == 0 && strings.ToLower(nextCave.name) == nextCave.name {
		me.alreadyTwice = true
	}
	me.pathPrimes *= nextCave.prime
}

func (me *path) isFinished() bool {
	return me.lastVisitedCave().name == "end"
}

func createNewPath() *path {
	newPath := path{}
	newPath.pathPrimes = 1
	return &newPath
}

func copyPath(oldPath *path) *path {
	newPath := createNewPath()
	for _, c := range oldPath.passedCaves {
		newPath.nextCave(c)
	}
	return newPath
}

func (me *path) moveToNextCaves() []path {
	newPaths := []path{}

	lastCave := me.lastVisitedCave()

	for _, c := range lastCave.connectedCaves {
		if me.isMovingPossible(c) {
			newPath := copyPath(me)

			newPath.nextCave(c)
			newPaths = appendDistinct(newPaths, *newPath)
		}
	}

	return newPaths
}

func (me *path) moveToNextCavesSmallTwice() []path {
	newPaths := []path{}

	lastCave := me.lastVisitedCave()

	for _, c := range lastCave.connectedCaves {
		if me.isMovingPossibleSmallTwice(c) {
			newPath := copyPath(me)

			newPath.nextCave(c)
			newPaths = appendDistinct(newPaths, *newPath)
		}
	}

	return newPaths
}

func (me *path) lastVisitedCave() *cave {
	return me.passedCaves[len(me.passedCaves)-1]
}

func calculateAllPaths(caves *[]*cave) *[]path {
	var finishedPaths []path = []path{}
	var workInProgressPaths []path = []path{}
	var initialPath path = *createNewPath()
	initialPath.nextCave(getOrCreateCave(caves, "start"))
	workInProgressPaths = append(workInProgressPaths, initialPath)
	for {
		var tempPaths []path = []path{}
		for _, p := range workInProgressPaths {
			paths := p.moveToNextCaves()
			for _, pp := range paths {
				if pp.isFinished() {
					finishedPaths = appendDistinct(finishedPaths, pp)
				} else {
					tempPaths = appendDistinct(tempPaths, pp)
				}
			}
		}

		if len(tempPaths) == 0 {
			break
		}
		workInProgressPaths = tempPaths
	}
	return &finishedPaths
}

func calculateAllPathsSmallTwice(caves *[]*cave) *[]path {
	var finishedPaths []path = []path{}
	var workInProgressPaths []path = []path{}
	var initialPath path = *createNewPath()
	initialPath.nextCave(getOrCreateCave(caves, "start"))
	workInProgressPaths = append(workInProgressPaths, initialPath)
	for {
		var tempPaths []path = []path{}
		for _, p := range workInProgressPaths {
			paths := p.moveToNextCavesSmallTwice()
			for _, pp := range paths {
				if pp.isFinished() {
					finishedPaths = appendDistinct(finishedPaths, pp)
				} else {
					tempPaths = appendDistinct(tempPaths, pp)
				}
			}
		}

		if len(tempPaths) == 0 {
			break
		}
		workInProgressPaths = tempPaths
		fmt.Printf("Currently %v paths in WIP\n", len(workInProgressPaths))
	}
	return &finishedPaths
}

func appendDistinct(paths []path, p path) []path {
	var ret = paths[:]
	takenPaths := make(map[string]bool)
	for _, e := range ret {
		takenPaths[pathString(e)] = true
	}
	_, prs := takenPaths[pathString(p)]
	if !prs {
		ret = append(ret, p)
	}
	return ret
}

func pathString(p path) string {
	var ret string = ""
	for _, e := range p.passedCaves {
		ret += e.name
	}
	return ret
}

func getOrCreateCave(caves *[]*cave, name string) *cave {
	var c *cave = nil
	for _, n := range *caves {
		if n.name == name {
			c = n
			break
		}
	}
	if c == nil {
		c = &cave{name, []*cave{}, primes[currentPrimeIndex]}
		currentPrimeIndex++
		*caves = append(*caves, c)
	}
	return c
}

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

func parseLine(line string) (nameOfFirstCave string, nameOfSecondCave string) {
	lineParts := strings.Split(line, "-")
	if len(lineParts) == 2 {
		nameOfFirstCave = lineParts[0]
		nameOfSecondCave = lineParts[1]
	}
	return
}

func parseFile(fileName string) *[]*cave {
	caves := []*cave{}
	lines := readCompleteFile(fileName)
	for _, line := range lines {
		first, second := parseLine(line)
		if len(first) == 0 {
			break
		}
		firstCave := getOrCreateCave(&caves, first)
		secondCave := getOrCreateCave(&caves, second)
		firstCave.connectCave(secondCave)
	}
	return &caves
}

func task1(fileName string) {
	currentPrimeIndex = 0
	caves := parseFile(fileName)
	paths := calculateAllPaths(caves)
	currentPrimeIndex = 0
	fmt.Printf("Found %v possible paths with no small cave visited twice\n", len(*paths))
}

func task2(fileName string) {
	currentPrimeIndex = 0
	caves := parseFile(fileName)
	paths := calculateAllPathsSmallTwice(caves) //

	fmt.Printf("Found %v possible paths with up to one small cave visited twice\n", len(*paths))
}

func main() {
	task1("day12-testdata1.txt")
	// task1("day12-testdata2.txt")
	// task1("day12-testdata3.txt")
	// task1("day12-inputdata.txt")
	task2("day12-testdata1.txt")
	// task2("day12-testdata2.txt")
	// task2("day12-testdata3.txt")
	task2("day12-inputdata.txt")
}
