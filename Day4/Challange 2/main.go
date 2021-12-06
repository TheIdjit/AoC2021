package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoNumber struct {
	value  int
	marked bool
}

type BingoRow struct {
	values [5]BingoNumber
}

type BingoCard struct {
	card  [5]BingoRow
	BINGO bool
}

func main() {
	input := getFile()
	Bingo(input)
	// fmt.Println(i, j)
	// Ans := CalcScore(i, j, k)
	// fmt.Println(Ans)
}

func Bingo(inp []string) (BingoCard, string) {
	BingoAnswers := [3]BingoCard{}
	BingoOrder := strings.Split(inp[0], ",")
	BingoCards := inp[2:]

	indexCard := 0
	indexRow := 0
	indexNum := 0

	LastBingo := 0
	LastBingoVal := ""

	for i := range BingoOrder {
		for j := range BingoCards {
			row := strings.Split(BingoCards[j], " ")
			if len(row) > 1 {
				for k := range row {
					if row[k] != " " && row[k] != "" {
						BingoAnswers[indexCard].card[indexRow].values[indexNum].value, _ = strconv.Atoi(row[k])

						if row[k] == BingoOrder[i] {
							BingoAnswers[indexCard].card[indexRow].values[indexNum].marked = true
						}

						if indexNum < 4 {
							indexNum++
						} else {
							indexNum = 0
						}
					}
				}
				indexRow++
			} else {
				indexCard++
				indexRow = 0
				indexNum = 0
			}
		}
		indexRow = 0
		indexCard = 0
		if i > 4 {
			BingoCount := 0
			for l := range BingoAnswers {
				for j := 0; j <= 4; j++ {
					for k := 0; k <= 4; k++ {
						if !BingoAnswers[l].BINGO {
							if BingoAnswers[l].card[k].values[j].marked == true {
								BingoCount++
							} else {
								BingoCount = 0
							}
							if BingoCount == 5 {
								fmt.Println("Vert Bingo")
								LastBingo = l
								LastBingoVal = BingoOrder[i]
								BingoAnswers[l].BINGO = true

								fmt.Println(BingoAnswers[LastBingo], LastBingoVal)
								return BingoAnswers[LastBingo], LastBingoVal
							}
						}
					}
				}

				for j := 0; j <= 4; j++ {
					for k := 0; k <= 4; k++ {
						if !BingoAnswers[l].BINGO {
							if BingoAnswers[l].card[j].values[k].marked == true {
								BingoCount++
							} else {
								BingoCount = 0
							}
							if BingoCount == 5 {
								LastBingo = l
								LastBingoVal = BingoOrder[i]
								BingoAnswers[l].BINGO = true
								fmt.Println(BingoAnswers[LastBingo], LastBingoVal)
								return BingoAnswers[LastBingo], LastBingoVal
							}
						}
					}
				}

			}
		}
	}
	fmt.Println(LastBingo, LastBingoVal)
	return BingoAnswers[LastBingo], LastBingoVal

}

func CalcScore(values [][5]string, marked [][5]bool, lastVal string) int {
	calcVal := 0

	for i := range marked {
		for j := range marked[i] {
			tempVal := 0
			if !marked[i][j] {
				tempVal, _ = strconv.Atoi(values[i][j])
				calcVal = calcVal + tempVal
			}
		}
	}

	fmt.Println(calcVal)
	tempval, _ := strconv.Atoi(lastVal)
	fmt.Println(lastVal)
	calcVal = calcVal * tempval

	return calcVal
}

func getFile() []string {
	val := []string{}
	file, err := os.Open("test.txt")
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

/*
99 24 44 83 47 // 99+24+44+47
 2 21 94 35  4 // 21+94+35+4
96 87 31  1 22 // 96+87+31+1
67  3 37 43 46
85 55 10  6 80 // 85+55+10+6+80
*/
