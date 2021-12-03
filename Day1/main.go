package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getFile()
	ans := compareLast(input)
	fmt.Println(ans)
}

func compareLast(inp []string) int {
	retVal := 0
	for i := 0; i < len(inp); i++ {
		if i >= 2 && i != len(inp)-1 {
			A1, _ := strconv.Atoi(inp[i])
			A2, _ := strconv.Atoi(inp[i-1])
			A3, _ := strconv.Atoi(inp[i-2])
			k := A1 + A2 + A3
			B1, _ := strconv.Atoi(inp[i-1])
			B2, _ := strconv.Atoi(inp[i])
			B3, _ := strconv.Atoi(inp[i+1])
			j := B1 + B2 + B3
			fmt.Println(k, j)
			if j > k {
				retVal++
			}
		}
	}
	return retVal
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
