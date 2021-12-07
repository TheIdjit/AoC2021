package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const InputFile string = "input.txt"

func main() {
	input := getFile()
	initVal := getInitval(input)
	ans := GetCoordinates(initVal)
	fmt.Println(ans)
}

func GetCoordinates(input []int) int {
	maxVal := input[len(input)-1]
	tempInt := []int{}
	tempVal := 0
	for i := 0; i <= maxVal; i++ {
		for j := range input {
			if input[j] > i {
				for k := 0; k <= (input[j] - i); k++ {
					tempVal = tempVal + k
				}
			} else {
				for k := 0; k <= (i - input[j]); k++ {
					tempVal = tempVal + k
				}
			}
		}
		tempInt = append(tempInt, (tempVal))
		tempVal = 0
	}

	sort.Sort(sort.IntSlice(tempInt))
	return tempInt[0]
}

func getInitval(input []string) []int {
	returnVal := []int{}
	splitVal := strings.Split(input[0], ",")
	for i := range splitVal {
		tempval, _ := strconv.Atoi(splitVal[i])

		returnVal = append(returnVal, tempval)
	}
	//Sort file in order of small to big
	sort.Sort(sort.IntSlice(returnVal))
	return returnVal
}

func getFile() []string {
	val := []string{}
	file, err := os.Open(InputFile)
	if err != nil {
		log.Fatal(err)
	}
	read := bufio.NewReader(file)
	for {
		freqMod, _, err := read.ReadLine()
		if err != nil {
			break
		}
		l := string(freqMod)
		val = append(val, l)
	}
	return val
}
