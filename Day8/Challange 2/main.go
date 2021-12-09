package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const InputFile string = "input.txt"

type Digits struct {
	inputDigits  []string
	outputDigits []string
}

type Decoder struct {
	Num0 string
	Num1 string
	Num2 string
	Num3 string
	Num4 string
	Num5 string
	Num6 string
	Num7 string
	Num8 string
	Num9 string
}

func main() {
	input := getFile()
	Digit := GetDigits(input)
	Deco := Decode(Digit.inputDigits)
	OutputVal(Digit.outputDigits, Deco)
	// fmt.Println(Digit)
}

func OutputVal(inp []string, decoder []Decoder) int {
	returnVal := 0
	totalVal := 0
	fmt.Println(len(decoder))
	for i := range inp {
		deco := decoder[i]
		outVal := strings.Split(inp[i], " ")
		for j := range outVal {
			Num1 := strings.Split(deco.Num1, "")
			Num4 := strings.Split(deco.Num4, "")
			switch len(outVal[j]) {
			case 2:
				returnVal = returnVal + 1*int(math.Pow(10, float64(len(outVal)-1-j)))
			case 3:
				returnVal = returnVal + 7*int(math.Pow(10, float64(len(outVal)-1-j)))
			case 4:
				returnVal = returnVal + 4*int(math.Pow(10, float64(len(outVal)-1-j)))
			case 7:
				returnVal = returnVal + 8*int(math.Pow(10, float64(len(outVal)-1-j)))
			}
			if len(outVal[j]) == 5 {
				if strings.Contains(outVal[j], Num1[0]) && strings.Contains(outVal[j], Num1[1]) {
					returnVal = returnVal + 3*int(math.Pow(10, float64(len(outVal)-1-j)))
				} else if (strings.Contains(outVal[j], Num4[0]) && strings.Contains(outVal[j], Num4[1])) || (strings.Contains(outVal[j], Num4[2]) && strings.Contains(outVal[j], Num4[3])) {
					returnVal = returnVal + 5*int(math.Pow(10, float64(len(outVal)-1-j)))
				} else {
					returnVal = returnVal + 2*int(math.Pow(10, float64(len(outVal)-1-j)))

				}
			}
			if len(outVal[j]) == 6 {
				if strings.Contains(outVal[j], Num4[0]) && strings.Contains(outVal[j], Num4[1]) && strings.Contains(outVal[j], Num4[2]) && strings.Contains(outVal[j], Num4[3]) {
					returnVal = returnVal + 9*int(math.Pow(10, float64(len(outVal)-1-j)))
				} else if strings.Contains(outVal[j], Num1[0]) && strings.Contains(outVal[j], Num1[1]) {
					returnVal = returnVal + 0*int(math.Pow(10, float64(len(outVal)-1-j)))
				} else {
					returnVal = returnVal + 6*int(math.Pow(10, float64(len(outVal)-1-j)))
				}
			}
		}
		fmt.Println(returnVal, i)
		totalVal = totalVal + returnVal
		returnVal = 0
	}
	// for i := range totalVal {
	// 	fmt.Println(returnVal)
	// 	returnVal = returnVal + totalVal[i]
	// }

	fmt.Println(totalVal)
	return returnVal
}

func Decode(inp []string) []Decoder {
	returnVal := Decoder{}
	returnValSlice := []Decoder{}
	for i := range inp {
		outVal := strings.Split(inp[i], " ")
		for j := range outVal {
			switch len(outVal[j]) {
			case 2:
				returnVal.Num1 = string(outVal[j][:])
			case 3:
				returnVal.Num7 = string(outVal[j][:])
			case 4:
				returnVal.Num4 = string(outVal[j][:])
			case 7:
				returnVal.Num8 = string(outVal[j][:])
			}
		}
		for j := range outVal {
			Num1 := strings.Split(returnVal.Num1, "")
			Num4 := strings.Split(returnVal.Num4, "")
			if len(outVal[j]) == 5 {
				if strings.Contains(outVal[j], Num1[0]) && strings.Contains(outVal[j], Num1[1]) {
					returnVal.Num3 = outVal[j]
				} else if (strings.Contains(outVal[j], Num4[0]) && strings.Contains(outVal[j], Num4[1])) || (strings.Contains(outVal[j], Num4[2]) && strings.Contains(outVal[j], Num4[3])) {
					returnVal.Num5 = outVal[j]
				} else {
					returnVal.Num2 = outVal[j]
				}
			}

			if len(outVal[j]) == 6 {
				if strings.Contains(outVal[j], Num4[0]) && strings.Contains(outVal[j], Num4[1]) && strings.Contains(outVal[j], Num4[2]) && strings.Contains(outVal[j], Num4[3]) {
					returnVal.Num9 = outVal[j]
				} else if strings.Contains(outVal[j], Num1[0]) && strings.Contains(outVal[j], Num1[1]) {
					returnVal.Num0 = outVal[j]
				} else {
					returnVal.Num6 = outVal[j]
				}
			}
		}
		returnValSlice = append(returnValSlice, returnVal)
		returnVal = Decoder{}
	}
	return returnValSlice
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
