package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getFile()
	ans := Postition(input)
	fmt.Println(ans)
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

func Postition(inp []string) int {
	returnVal := 0
	depth := 0
	horzPos := 0
	for i := 0; i < len(inp); i++ {
		fmt.Println(inp[i])
		s := strings.Split(inp[i], " ")
		val, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			horzPos = horzPos + val
		case "up":
			depth = depth - val
		case "down":
			depth = depth + val
		}
	}
	returnVal = depth * horzPos
	return returnVal
}
