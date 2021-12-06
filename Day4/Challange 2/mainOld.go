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
	fmt.Println(j)
	ans := CalcScore(i, j, k)
	fmt.Println(ans)
}

func Bingo(inp []string) ([][5]string, [][5]bool, string) {
	BingoAnswers := [100][5][5]bool{}
	BingoAnswerCard := [100][5][5]string{}
	Bingo := [100]bool{}
	BingoOrder := strings.Split(inp[0], ",")
	BingoCards := inp[2:]

	indexRow := 0
	indexCard := 0
	indexNum := 0

	LastBingo := 0
	lastBingoVal := ""

	for i := range BingoOrder {
		for j := range BingoCards {
			row := strings.Split(BingoCards[j], " ")
			if len(row) > 1 {
				for k := range row {
					if row[k] != " " && row[k] != "" {
						BingoAnswerCard[indexCard][indexRow][indexNum] = row[k]
						if row[k] == BingoOrder[i] {
							if j == 54 {
								fmt.Println(BingoOrder[i], "Fuck My Life")
								fmt.Println(BingoAnswers[54])
							}
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
			ans := false
			ind := 0
			ans, ind, Bingo = CheckBingo(BingoAnswers[:][:], Bingo)
			if ans {
				lastBingoVal = BingoOrder[i]
				LastBingo = ind
				allDone := true

				if LastBingo == 54 {
					// fmt.Println(lastBingoVal, LastBingo, BingoAnswers[LastBingo][:])
					return BingoAnswerCard[LastBingo][:], BingoAnswers[LastBingo][:], lastBingoVal
				}

				for g := range Bingo {
					if Bingo[g] == false {
						allDone = false
						break
					}
				}
				if allDone == true {
					// fmt.Println(Bingo)
					// fmt.Println(lastBingoVal, LastBingo, BingoAnswers[LastBingo][:])
					return BingoAnswerCard[LastBingo][:], BingoAnswers[LastBingo][:], lastBingoVal
				}
			}
		}
	}

	fmt.Println(lastBingoVal, LastBingo, BingoAnswers[LastBingo])
	return BingoAnswerCard[LastBingo][:], BingoAnswers[LastBingo][:], lastBingoVal
}

func CheckBingo(BingoAnswers [][5][5]bool, Bingo [100]bool) (bool, int, [100]bool) {
	Bingoval := 0
	LastBingo := 0
	bingo := false
	for j := 0; j <= len(BingoAnswers)-1; j++ {
		if !Bingo[j] {
			for k := 0; k <= 4; k++ {
				for l := 0; l <= 4; l++ {
					if BingoAnswers[j][l][k] == true && Bingo[j] == false {
						Bingoval++
					} else {
						Bingoval = 0
						break
					}
					if Bingoval == 5 {
						// fmt.Println("BINGO VERT", j, BingoAnswers[j])
						Bingo[j] = true
						LastBingo = j
						bingo = true
						fmt.Println(j)
						// return true, j, Bingo
					}
				}
			}
		}
	}

	for j := 0; j <= len(BingoAnswers)-1; j++ {
		for k := 0; k <= 4; k++ {
			for l := 0; l <= 4; l++ {
				if BingoAnswers[j][k][l] == true && Bingo[j] == false {
					Bingoval++
				} else {
					Bingoval = 0
					break
				}
				if Bingoval == 5 {
					// fmt.Println("BINGO", j, BingoAnswers[j])
					Bingo[j] = true
					LastBingo = j
					bingo = true
					fmt.Println(j)

				}
			}
		}
	}
	if bingo {
		return true, LastBingo, Bingo
	} else {
		return false, LastBingo, Bingo
	}
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
