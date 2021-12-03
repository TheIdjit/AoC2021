package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	input := getFile()
	val := CalcPower(input)
	fmt.Println(val)
}

func getFile() []string {
	val := []string{}
	file, err := os.Open("input.txt")
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

func CalcPower(inp []string) int {
	reutrnPower := 0
	GammaRate := 0
	EpsilonRate := 0
	posBit := [12]int{}
	negBit := [12]int{}
	for i := range inp {
		for j := range inp[i] {
			switch inp[i][j] {
			case 48:
				negBit[j]++
			case 49:
				posBit[j]++
			}

		}
	}

	for k := range posBit {
		if posBit[k] > negBit[k] {
			GammaRate = GammaRate + int(math.Pow(2, float64(11-k)))

		} else {
			EpsilonRate = EpsilonRate + int(math.Pow(2, float64(11-k)))
		}
	}
	reutrnPower = GammaRate * EpsilonRate
	return reutrnPower
}
