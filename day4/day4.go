package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getBoards(data []string) [][]int {
	var boards [][]int
	for i := 2; i < len(data); i += 6 {
		var board []int
		boardCut := strings.Split(strings.Join(data[i:i+5], " "), " ")
		fmt.Println(boardCut)
		for _, s := range boardCut {
			if s == "" {
				continue
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}

			board = append(board, i)
		}

		boards = append(boards, board)
	}
	return boards
}

func winRow(b []int) bool {
	win := true
	for i := 0; i < 5; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 5; i < 10; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 10; i < 15; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 15; i < 20; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 20; i < 25; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	return false
}

func winCol(b []int) bool {
	/*
		win := true
		for i := 0; i < 5; i++ {
			for j := i; j < 25; j += 5 {
				//fmt.Println(j)
				if b[j] != 0 {
					win = false
				}
			}
			if win {
				return true
			}
		}
		return false
	*/
	win := true
	for i := 0; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 1; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 3; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 4; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}
	return false
}

func checkWin(b []int) bool {
	if winRow(b) || winCol(b) {
		return true
	}
	return false
}

func secondPart() {
	data := readFile("input.txt")

	var numbers []int
	for _, s := range strings.Split(data[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}

	var boards [][]int
	boards = getBoards(data)
	boardWin := make([]bool, len(boards))

	// ok up to here
	//fmt.Println(numbers, boards)

	// zero numbers in all boards
	for _, number := range numbers {
		for board := range boards {

			if boardWin[board] {
				continue
			}
			for i, v := range boards[board] {
				if v == number {
					boards[board][i] = 0
					break
				}
			}
			// check for win
			if checkWin(boards[board]) {
				sum := 0
				for _, j := range boards[board] {
					sum += j
				}
				fmt.Println(number, boards[board], sum*number)
				boardWin[board] = true
			}
		}
	}

}

func firstPart() {
	data := readFile("input.txt")

	var numbers []int
	for _, s := range strings.Split(data[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}

	var boards [][]int
	boards = getBoards(data)

	// ok up to here
	//fmt.Println(numbers, boards)

	// zero numbers in all boards
	for _, number := range numbers {
		for _, board := range boards {
			for i, v := range board {
				//fmt.Println(i, v)
				if v == number {
					board[i] = 0
					break
				}
			}
			// check for win
			if checkWin(board) {
				fmt.Println("Winning board: ", board)
				fmt.Println("Winning number: ", number)
				sum := 0
				for _, j := range board {
					sum += j
				}
				fmt.Println(number, "*", sum, "=", number*sum)
				return
			}
		}
	}

}

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
	var values []string

	// convert all values in slice to integer
	for sc.Scan() {
		//num, _ := strconv.Atoi(sc.Text())
		values = append(values, sc.Text())
	}

	return values
}

func main() {
	// get the most common and least common bit of a binary number (each bit)
	// gamma rate: most common bit for each char row
	// epsilon rate: least common bit for each char row (or inverse of gamma rate)
	// gamma/epsilon rate -> decimal
	// return mulitplication of both

	firstPart()
	//secondPart()
}
