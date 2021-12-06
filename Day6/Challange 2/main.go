package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile string = "input.txt"

type Population struct {
	day8 int
	day7 int
	day6 int
	day5 int
	day4 int
	day3 int
	day2 int
	day1 int
	day0 int
}

func main() {
	input := getFile()
	initVal := getInitval(input)
	popu := SetUpPopu(initVal)
	ans := GetPopulationGroth(popu, 256)
	fmt.Println(ans)
}

func GetPopulationGroth(initVal Population, days int) int {

	for i := 1; i <= days; i++ {
		zeroVal := initVal.day0
		initVal.day0 = initVal.day1
		initVal.day1 = initVal.day2
		initVal.day2 = initVal.day3
		initVal.day3 = initVal.day4
		initVal.day4 = initVal.day5
		initVal.day5 = initVal.day6
		initVal.day6 = initVal.day7 + zeroVal
		initVal.day7 = initVal.day8
		initVal.day8 = zeroVal
	}

	returnVal := initVal.day0 + initVal.day1 + initVal.day2 + initVal.day3 + initVal.day4 + initVal.day5 + initVal.day6 + initVal.day7 + initVal.day8
	return returnVal
}

func SetUpPopu(input []int) Population {
	returnVal := Population{}
	for i := range input {
		// fmt.Println(input[i])
		switch input[i] {
		case 0:
			returnVal.day0++
		case 1:
			returnVal.day1++
		case 2:
			returnVal.day2++
		case 3:
			returnVal.day3++
		case 4:
			returnVal.day4++
		case 5:
			returnVal.day5++
		case 6:
			returnVal.day6++
		case 7:
			returnVal.day7++
		case 8:
			returnVal.day8++
		}
	}

	return returnVal
}

func getInitval(input []string) []int {
	returnVal := []int{}
	splitVal := strings.Split(input[0], ",")
	for i := range splitVal {
		tempval, _ := strconv.Atoi(splitVal[i])

		returnVal = append(returnVal, tempval)
	}

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
