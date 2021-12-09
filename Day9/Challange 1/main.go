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

func main() {
	input := getFile()
	ans := GetLowPoints(input)
	fmt.Println(ans)
}

func GetLowPoints(inp []string) int {
	returnVal := 0
	OldRow := []string{}
	nextrow := []string{}
	for i := range inp {
		row := strings.Split(inp[i], "")
		if i < len(inp)-1 {
			nextrow = strings.Split(inp[i+1], "")
		}
		for j := range row {
			if j > 0 {
				if j < len(row)-1 && i > 0 && i < len(inp)-1 {
					if row[j] < row[j-1] && row[j] < row[j+1] && row[j] < OldRow[j] && row[j] < nextrow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("1", tempval)
						returnVal = returnVal + tempval + 1
					}
				}

				if j == len(row)-1 && i > 0 && i < len(inp)-1 {
					if row[j] < row[j-1] && row[j] < OldRow[j] && row[j] < nextrow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("1", tempval)
						returnVal = returnVal + tempval + 1
					}
				}

				if i == 0 && j < len(row)-1 && j < len(row)-1 {
					if row[j] < row[j-1] && row[j] < row[j+1] && row[j] < nextrow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("2", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
				if i == len(inp)-1 && j < len(row)-1 {
					if row[j] < row[j-1] && row[j] < row[j+1] && row[j] < OldRow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("3", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
			}

			if j == 0 {
				if i > 0 && i < len(inp)-1 {
					if row[j] < row[j+1] && row[j] < OldRow[j] && row[j] < nextrow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("4", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
				if i == 0 {
					if row[j] < row[j+1] && row[j] < nextrow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("5", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
				if i == len(inp)-1 {
					if row[j] < row[j+1] && row[j] < OldRow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("6", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
			}

			if j == len(row)-1 {
				if i == 0 {
					if row[j] < row[j-1] && row[j] < nextrow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("8", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
				if i == len(inp)-1 {
					if row[j] < row[j-1] && row[j] < OldRow[j] {
						tempval, _ := strconv.Atoi(row[j])
						fmt.Println("9", tempval)
						returnVal = returnVal + tempval + 1
					}
				}
			}

		}
		OldRow = row
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
