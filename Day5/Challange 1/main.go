package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const Coordinates int = 1000
const structCoor int = 500
const InputFile string = "input.txt"

type coordinates struct {
	x1 int // Start coordinate y-domain
	y1 int // Start coordinate y-axes
	x2 int // End coordinate y-domain
	y2 int // End coordinate y-axes
}

func main() {
	input := getFile()
	Coordinates := GetCoordinates(input)
	MapOfVents := Makemap(Coordinates)

	// for i := range MapOfVents {
	// 	// fmt.Println(Coordinates[i])
	// 	fmt.Println(MapOfVents[i])
	// }
	ans := GetIntersections(MapOfVents)

	fmt.Println(ans)

	// fmt.Println(Coordinates[0])
}

func Makemap(Coor [structCoor]coordinates) [Coordinates][Coordinates]int {
	Map := [Coordinates][Coordinates]int{}

	for i := range Coor {
		if Coor[i].x1 == Coor[i].x2 {
			if Coor[i].y1 > Coor[i].y2 {
				for j := 0; j <= Coor[i].y1-Coor[i].y2; j++ {
					Map[(Coor[i].y2 + j)][Coor[i].x1]++
				}
			}
			if Coor[i].y2 > Coor[i].y1 {
				for j := 0; j <= Coor[i].y2-Coor[i].y1; j++ {
					Map[(Coor[i].y1 + j)][Coor[i].x1]++
				}
			}
		}

		if Coor[i].y1 == Coor[i].y2 {
			if Coor[i].x1 > Coor[i].x2 {
				for j := 0; j <= Coor[i].x1-Coor[i].x2; j++ {
					Map[(Coor[i].y1)][Coor[i].x2+j]++
				}
			}
			if Coor[i].x2 > Coor[i].x1 {
				for j := 0; j <= Coor[i].x2-Coor[i].x1; j++ {
					Map[(Coor[i].y1)][Coor[i].x1+j]++
				}
			}
		}

		if Coor[i].x1 > Coor[i].x2 {
			if Coor[i].y1 > Coor[i].y2 {
				for j := 0; j <= Coor[i].y1-Coor[i].y2; j++ {
					Map[(Coor[i].y1 - j)][Coor[i].x1-j]++
				}
			}
			if Coor[i].y2 > Coor[i].y1 {
				for j := 0; j <= Coor[i].y2-Coor[i].y1; j++ {
					Map[(Coor[i].y1 + j)][Coor[i].x1-j]++
				}
			}
		}

		if Coor[i].x1 < Coor[i].x2 {
			if Coor[i].y1 > Coor[i].y2 {
				for j := 0; j <= Coor[i].y1-Coor[i].y2; j++ {
					Map[(Coor[i].y1 - j)][Coor[i].x1+j]++
				}
			}
			if Coor[i].y2 > Coor[i].y1 {
				for j := 0; j <= Coor[i].y2-Coor[i].y1; j++ {
					Map[(Coor[i].y1 + j)][Coor[i].x1+j]++
				}
			}
		}
	}

	return Map
}

func GetIntersections(ventMap [Coordinates][Coordinates]int) int {
	returnval := 0
	for i := range ventMap {
		for j := range ventMap[i] {
			if ventMap[i][j] >= 2 {
				returnval++
			}

		}
	}

	return returnval
}

func GetCoordinates(input []string) [structCoor]coordinates {
	returnVal := [structCoor]coordinates{}

	for i := range input {
		Coor := strings.Split(input[i], " -> ")
		x := strings.Split(Coor[0], ",")
		y := strings.Split(Coor[1], ",")

		x1, _ := strconv.Atoi(x[0])
		x2, _ := strconv.Atoi(x[1])
		y1, _ := strconv.Atoi(y[0])
		y2, _ := strconv.Atoi(y[1])
		returnVal[i] = coordinates{
			x1,
			x2,
			y1,
			y2,
		}
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
