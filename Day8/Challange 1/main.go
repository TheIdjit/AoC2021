package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const InputFile string = "input.txt"

type Digits struct {
	inputDigits  []string
	outputDigits []string
}

func main() {
	input := getFile()
	Digit := GetDigits(input)
	ans := GetOneFourSevenEight(Digit)
	fmt.Println(ans)
}

func GetOneFourSevenEight(input Digits) int {
	count := 0

	for i := range input.outputDigits {
		outputVal := strings.Split(input.outputDigits[i], " ")
		for j := range outputVal {
			if len(outputVal[j]) == 2 || len(outputVal[j]) == 3 || len(outputVal[j]) == 4 || len(outputVal[j]) == 7 {
				count++
			}
		}
	}

	return count
}

func GetDigits(input []string) Digits {
	allDigits := []string{}
	inDigits := []string{}
	outDigits := []string{}

	for i := range input {
		allDigits = strings.Split(input[i], " | ")
		inDigits = append(inDigits, allDigits[0])
		outDigits = append(outDigits, allDigits[1])
	}

	return Digits{
		inputDigits:  inDigits,
		outputDigits: outDigits,
	}
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
