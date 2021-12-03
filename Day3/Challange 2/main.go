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
	val := CalcLifeSupport(input)
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

func CalcLifeSupport(inp []string) int {
	PosBitSlice := []string{}
	NegBitSlice := []string{}
	Oxygen := []string{}
	Co2 := []string{}
	returnVal := 0
	for i := range inp {
		switch inp[i][0] {
		case 49:
			PosBitSlice = append(PosBitSlice, inp[i])
		case 48:
			NegBitSlice = append(NegBitSlice, inp[i])
		}
	}
	if len(PosBitSlice) >= len(NegBitSlice) {
		Oxygen = PosBitSlice
		Co2 = NegBitSlice
	} else {
		Co2 = PosBitSlice
		Oxygen = NegBitSlice
	}
	OxyVal := CalcOxy(Oxygen)
	Co2Val := CalcCo2(Co2)

	returnVal = OxyVal * Co2Val
	return returnVal
}

func CalcOxy(inp []string) int {
	oxySlice := inp
	j := 1
	returnVal := 0
	for len(oxySlice) > 1 {
		PosBitSlice := []string{}
		NegBitSlice := []string{}
		for i := range oxySlice {
			switch oxySlice[i][j] {
			case 49:
				PosBitSlice = append(PosBitSlice, oxySlice[i])
			case 48:
				NegBitSlice = append(NegBitSlice, oxySlice[i])
			}
		}
		if len(PosBitSlice) >= len(NegBitSlice) {
			oxySlice = PosBitSlice
		} else {
			oxySlice = NegBitSlice
		}
		j++
	}

	for index := range oxySlice[0] {
		if oxySlice[0][index] == 49 {
			returnVal = returnVal + int(math.Pow(2, float64(11-index)))
		}
	}

	fmt.Println(returnVal)

	return returnVal
}

func CalcCo2(inp []string) int {
	Co2Slice := inp
	j := 1
	returnVal := 0
	for len(Co2Slice) > 1 {
		PosBitSlice := []string{}
		NegBitSlice := []string{}
		for i := range Co2Slice {
			switch Co2Slice[i][j] {
			case 49:
				PosBitSlice = append(PosBitSlice, Co2Slice[i])
			case 48:
				NegBitSlice = append(NegBitSlice, Co2Slice[i])
			}
		}
		if len(PosBitSlice) >= len(NegBitSlice) {
			Co2Slice = NegBitSlice
		} else {
			Co2Slice = PosBitSlice
		}
		j++
	}

	for index := range Co2Slice[0] {
		if Co2Slice[0][index] == 49 {
			returnVal = returnVal + int(math.Pow(2, float64(11-index)))
		}
	}

	fmt.Println(returnVal)
	return returnVal
}

// func CalcPower(inp []string) int {
// 	reutrnPower := 0
// 	GammaRate := 0
// 	EpsilonRate := 0
// 	posBit := [12]int{}
// 	negBit := [12]int{}
// 	for i := range inp {
// 		for j := range inp[i] {
// 			switch inp[i][j] {
// 			case 48:
// 				negBit[j]++
// 			case 49:
// 				posBit[j]++
// 			}

// 		}
// 	}

// 	for k := range posBit {
// 		if posBit[k] > negBit[k] {
// 			GammaRate = GammaRate + int(math.Pow(2, float64(11-k)))

// 		} else {
// 			EpsilonRate = EpsilonRate + int(math.Pow(2, float64(11-k)))
// 		}
// 	}
// 	reutrnPower = GammaRate * EpsilonRate
// 	return reutrnPower
// }
