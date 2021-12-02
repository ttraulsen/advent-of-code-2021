package main

import (
  "fmt"
  "io/ioutil"
	"strconv"
  "os"
	"strings"
)

type task interface {
	processCommand([]string)
	getResult() int64
}

type cTask1 struct {
	horizontal int64
	depth int64
}

func NewTask1() cTask1 {
	t := cTask1{0,0}
	return t
}

func (t *cTask1) processCommand(line []string){
	if(len(line)!=2){
		fmt.Printf("Error processing line: %s\n", line)
		return
	}

	var value,_ = strconv.ParseInt(line[1], 10, 64)
	switch line[0] {
	case "forward":
		t.horizontal = t.horizontal + value
	case "down":
		t.depth = t.depth + value
	case "up":
		t.depth = t.depth - value
	default:
		fmt.Printf("Error while reading line %v\n",line)
	}
}

func (t *cTask1) getResult() int64 {
	return t.depth*t.horizontal
}


type cTask2 struct {
	horizontal int64
	depth int64
	aim int64
}

func NewTask2() cTask2 {
	t := cTask2{0,0,0}
	return t
}

func (t *cTask2) processCommand(line []string){
	if(len(line)!=2){
		fmt.Printf("Error processing line: %s\n", line)
		return
	}
	var value,_ = strconv.ParseInt(line[1], 10, 64)
	switch line[0] {
	case "forward":
		t.horizontal = t.horizontal + value
		t.depth = t.depth + t.aim*value
	case "down":
		t.aim = t.aim + value
	case "up":
		t.aim = t.aim - value								
	default:
		fmt.Printf("Error while reading line %v\n",line)
	}
}

func (t *cTask2) getResult() int64 {
	return t.depth*t.horizontal
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func processFile (fileName string,t task) {
	var lines = readCompleteFile(fileName)


	for _, line := range lines {
		var splitLine = strings.Split(line, " ")
		if(len(splitLine)<2){
			continue
		}
		t.processCommand(splitLine)
	}

	fmt.Printf("With file %v - task: %T, product: %v\n",fileName,t ,t.getResult())

}

func readCompleteFile(fileName string) []string{
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}


func main() {
	var t1 cTask1 = NewTask1()
	processFile("testdata.txt",&t1)
	t1 = NewTask1()
	processFile("day2-inputdata-1.txt",&t1)

	var t2 cTask2 = NewTask2()
	processFile("testdata.txt",&t2)
	t2 = NewTask2()
	processFile("day2-inputdata-1.txt",&t2)

}
