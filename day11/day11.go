package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) []string {
	// read input.txt
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
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

func oneStep(grid [][]int) int {
	flashed := [10][10]bool{}

	// increment
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x] += 1
		}
	}

	// increment
	updatedBool := true
	for updatedBool {
		updatedBool = false

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if flashed[y][x] {
					continue
				}

				if grid[y][x] > 9 {
					updatedBool = true
					flashed[y][x] = true

					if y-1 >= 0 {
						if x-1 >= 0 {
							grid[y-1][x-1] += 1
						}
						grid[y-1][x] += 1

						if x+1 < len(grid[y]) {
							grid[y-1][x+1] += 1
						}
					}

					if x-1 >= 0 {
						grid[y][x-1] += 1
					}
					if x+1 < len(grid[y]) {
						grid[y][x+1] += 1
					}

					if y+1 < len(grid) {
						if x-1 >= 0 {
							grid[y+1][x-1] += 1
						}
						grid[y+1][x] += 1

						if x+1 < len(grid[y]) {
							grid[y+1][x+1] += 1
						}
					}
				}
			}
		}
	}

	count := 0
	for y := 0; y < len(flashed); y++ {
		for x := 0; x < len(flashed[y]); x++ {
			if flashed[y][x] {
				grid[y][x] = 0
				count += 1
			}
		}
	}
	return count
}

func firstPart() {
	var data = readFile("input.txt")
	var grid [][]int = make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)

		for j := 0; j < 10; j++ {
			grid[i][j] = int(data[i][j] - byte('0'))
		}
	}
	//fmt.Println(grid)
	//fmt.Println(oneStep(grid))
	cnt := 0
	for i := 0; i < 100; i++ {
		cnt += oneStep(grid)
	}
	fmt.Println("Solution for part 1: ", cnt)
}

func secondPart() {
	var data = readFile("input.txt")
	var grid [][]int = make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)

		for j := 0; j < 10; j++ {
			grid[i][j] = int(data[i][j] - byte('0'))
		}
	}
	//fmt.Println(grid)
	//fmt.Println(oneStep(grid))
	cnt := 0
	for i := 0; i < 1000; i++ {
		nCnt := oneStep(grid)
		cnt += nCnt
		if nCnt == 100 {
			fmt.Println("Solution for part 2: ", i+1)
			break
		}
	}
}

func main() {
	firstPart()
	secondPart()
}
