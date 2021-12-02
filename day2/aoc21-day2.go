package main

import (
  "fmt"
  "io/ioutil"
	"strconv"
  "os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func task1 (fileName string) {

	var lines = readCompleteFile(fileName)
	var horizontal int64 = 0
	var depth int64 = 0

	for _, line := range lines {
		var splitLine = strings.Split(line, " ")
		if(len(splitLine)<2){
			continue
		}
		var value,_ = strconv.ParseInt(splitLine[1], 10, 64)
		switch splitLine[0] {
		case "forward":
			horizontal = horizontal + value
		case "down":
			depth = depth + value
		case "up":
			depth = depth - value
		default:
			fmt.Printf("Error while reading line %v\n",line)
		}
	}

	fmt.Printf("With file %v - horizontal: %v, depth: %v, product: %v\n",fileName,horizontal,depth,horizontal*depth)

}

func task2 (fileName string) {

	var lines = readCompleteFile(fileName)
	var horizontal int64 = 0
	var aim int64 = 0
	var depth int64 = 0

	for _, line := range lines {
		var splitLine = strings.Split(line, " ")
		if(len(splitLine)<2){
			continue
		}
		var value,_ = strconv.ParseInt(splitLine[1], 10, 64)
		switch splitLine[0] {
		case "forward":
			horizontal = horizontal + value
			depth = depth + aim*value
		case "down":
			aim = aim + value
		case "up":
			aim = aim - value
		default:
			fmt.Printf("Error while reading line %v\n",line)
		}
	}

	fmt.Printf("With file %v - horizontal: %v, depth: %v, aim: %v, product: %v\n",fileName,horizontal,depth,aim,horizontal*depth)

}

func readCompleteFile(fileName string) []string{
	pwd, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(pwd + "/" + fileName)
	check(err)

	return strings.Split(string(fileBytes), "\n")
}


func main() {
	task1("testdata.txt")
	task1("day2-inputdata-1.txt")
	task2("testdata.txt")
	task2("day2-inputdata-1.txt")
}
