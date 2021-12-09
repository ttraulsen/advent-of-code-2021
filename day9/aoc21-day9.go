package main

import (
  "fmt"
  "io/ioutil"
  "os"
	"strings"
	"strconv"
	"sort"
)



func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readCompleteFile(fileName string) []string{
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}

func readMap(lines []string) *[][]int64 {
	
	var ret = [][]int64{}

	for _, l := range lines {
		if(len(l)<1){continue}
		line := make([]int64,len(l))

		chars := strings.Split(l,"")

		for i, c := range chars {
			n,_ := strconv.ParseInt(c,10, 64)
			line[i]=n
		}
		ret=append(ret,line)
	}
	return &ret
}

func readColoredMap(lines []string) *[][]int64 {
	var ret = [][]int64{}

	for _, l := range lines {
		if(len(l)<1){continue}
		line := make([]int64,len(l))
		ret=append(ret,line)
	}
	return &ret
}

func isLocalMinimum(m *[][]int64,x int,y int) bool{
	myMap := *m
	if(x < 0 || y < 0 || x >= len(myMap) || y >= len(myMap[x])){return false}

	if(x > 0 && myMap[x-1][y] <= myMap[x][y]){
		return false
	}

	if(x < len(myMap)-1 && myMap[x+1][y] <= myMap[x][y]){
		return false
	}
	
	if(y > 0 && myMap[x][y-1] <= myMap[x][y]){
		return false
	}

	if(y < len(myMap[x])-1 && myMap[x][y+1] <= myMap[x][y]){
		return false
	}	
	return true
}

func task1 (fileName string) {
	myMap := *(readMap(readCompleteFile(fileName)))


	var sumOfMinimums = int64(0)

	for x := 0; x < len(myMap); x++ {
		for y := 0; y < len(myMap[x]);y++ {
			if(isLocalMinimum(&myMap,x,y)){
				sumOfMinimums+=myMap[x][y]+1
			}
		}
	}

	fmt.Printf("Sum of minimums %v\n",sumOfMinimums)
}

type minimum struct {
	x int
	y int
	id int
}

func task2 (fileName string) {
	myMap := *(readMap(readCompleteFile(fileName)))
	coloredMap := *(readColoredMap(readCompleteFile(fileName)))
	
	var listOfMinimums = []minimum{}

	var id = 1

	for x := 0; x < len(myMap); x++ {
		for y := 0; y < len(myMap[x]);y++ {
			if(isLocalMinimum(&myMap,x,y)){
				var localMinimum = minimum{x,y,id}
				listOfMinimums = append(listOfMinimums,localMinimum)
				coloredMap[x][y] = int64(id)
				id++
			}
		}
	}


	for {
		hasColored := false

		for x := 0; x < len(myMap); x++ {
			for y := 0; y < len(myMap[x]);y++ {
				if(myMap[x][y] == 9){
					continue
				}
				if(coloredMap[x][y] != 0){
					continue
				}

				if(x > 0 && coloredMap[x-1][y] != 0){
					coloredMap[x][y]=coloredMap[x-1][y]
					hasColored=true
				}

				if(x < len(myMap)-1 && coloredMap[x+1][y] != 0){
					coloredMap[x][y]=coloredMap[x+1][y]
					hasColored=true				
				}
				
				if(y > 0 && coloredMap[x][y-1] != 0){					
					coloredMap[x][y]=coloredMap[x][y-1]
					hasColored=true
				}

				if(y < len(myMap[x])-1 && coloredMap[x][y+1] != 0){
					coloredMap[x][y]=coloredMap[x][y+1]
					hasColored=true
				}
			}
		}
		if(!hasColored){break}
	}

	var sizeOfBasins = map[int64]int{}

	for x := 0; x < len(myMap); x++ {
		for y := 0; y < len(myMap[x]);y++ {
			sizeOfBasins[coloredMap[x][y]]++
		}
	}	
	delete(sizeOfBasins,0)

	values := make([]int, 0, len(sizeOfBasins))

	for _, v := range sizeOfBasins {
		values = append(values, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	fmt.Printf("Basin sizes %v\n",values)

	fmt.Printf("Answer to task2 is %v\n",values[0]*values[1]*values[2])
}



func main() {
	task1("day9-testdata.txt")
	task1("day9-inputdata.txt")
	task2("day9-testdata.txt")
	task2("day9-inputdata.txt")
}
