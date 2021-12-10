package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) []string {
	// read input.txt
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory!")
	}
	defer file.Close()

	// add file content to values slice
	sc := bufio.NewScanner(file)
	var valuesStr []string

	// convert all values in slice to integer
	for sc.Scan() {
		//num, _ := strconv.Atoi(sc.Text())
		valuesStr = append(valuesStr, sc.Text())
	}

	return valuesStr
}

func firstPart() {
	var data = readFile("input.txt")
	var dataGrid [][]int = make([][]int, len(data))

	for i, row := range data {
		dataGrid[i] = make([]int, len(row))

		var err error
		for j, val := range row {
			dataGrid[i][j], err = strconv.Atoi(string(val))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	maxX, maxY := len(data[0]), len(data)
	risk := 0

	// brute-force FTW
	for y := 0; y < len(dataGrid); y++ {
		for x := 0; x < len(dataGrid[y]); x++ {
			cur := dataGrid[y][x]

			// check up, down, left, right
			if y > 0 && dataGrid[y-1][x] <= cur {
				continue
			}
			if y < maxY-1 && dataGrid[y+1][x] <= cur {
				continue
			}

			if x > 0 && dataGrid[y][x-1] <= cur {
				continue
			}
			if x < maxX-1 && dataGrid[y][x+1] <= cur {
				continue
			}
			risk += cur + 1
		}
	}
	//fmt.Println(dataGrid)
	fmt.Println(risk)
}

func main() {
	firstPart()
	//	secondPart()
}
