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
	i, j, k := Bingo(input)
	Ans := CalcScore(i, j, k)
	fmt.Println(Ans)
}

func Bingo(inp []string) ([][5]string, [][5]bool, string) {
	BingoAnswers := [3][5][5]bool{}
	BingoAnswerCard := [3][5][5]string{}
	// Bingo := [100]bool{}
	BingoOrder := strings.Split(inp[0], ",")
	BingoCards := inp[2:]

	indexRow := 0
	indexCard := 0
	indexNum := 0

	for i := range BingoOrder {
		for j := range BingoCards {
			row := strings.Split(BingoCards[j], " ")
			if len(row) > 1 {
				for k := range row {
					if row[k] != " " && row[k] != "" {
						BingoAnswerCard[indexCard][indexRow][indexNum] = row[k]
						if row[k] == BingoOrder[i] {
							BingoAnswers[indexCard][indexRow][indexNum] = true
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
			ans, ind := CheckBingo(BingoAnswers[:][:])
			if ans {
				return BingoAnswerCard[ind][:], BingoAnswers[ind][:], BingoOrder[i]
			}
		}
	}
	return nil, nil, ""
}

func CheckBingo(BingoAnswers [][5][5]bool) (bool, int) {
	Bingoval := 0
	for j := 0; j <= len(BingoAnswers)-1; j++ {
		for k := 0; k <= 4; k++ {
			for l := 0; l <= 4; l++ {
				if BingoAnswers[j][l][k] == true {
					Bingoval++
				} else {
					Bingoval = 0
					break
				}
				if Bingoval == 5 {
					fmt.Println("BINGO VERT", BingoAnswers[j])
					return true, j
				}
			}
		}
	}
	// Over de bingo antwoorden heen gaan om te checken of er Bingo is(Horizontaal)
	for j := 0; j <= len(BingoAnswers)-1; j++ {
		for k := 0; k <= 4; k++ {
			for l := 0; l <= 4; l++ {
				if BingoAnswers[j][k][l] == true {
					Bingoval++
				} else {
					Bingoval = 0
					break
				}
				if Bingoval == 5 {
					fmt.Println("BINGO", j, BingoAnswers[j])
					return true, j
				}
			}
		}
	}
	return false, 0
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
