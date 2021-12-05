package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	var values []string

	// convert all values in slice to integer
	for sc.Scan() {
		//num, _ := strconv.Atoi(sc.Text())
		values = append(values, sc.Text())
	}

	return values
}

type Coords struct {
	x1, y1 int
	x2, y2 int
}

func secondPart() {

	fName := "input.txt"

	data := readFile(fName)

	var maxX, maxY int
	var coordSlice = make([]Coords, len(data))
	consideredLines := []Coords{}
	for i, line := range data {
		fmt.Sscanf(line, "%d,%d -> %d,%d",
			&coordSlice[i].x1,
			&coordSlice[i].y1,
			&coordSlice[i].x2,
			&coordSlice[i].y2)

		// check for max number
		if coordSlice[i].x1 > maxX {
			maxX = coordSlice[i].x1
		}

		if coordSlice[i].x2 > maxX {
			maxX = coordSlice[i].x2
		}

		if coordSlice[i].y1 > maxY {
			maxY = coordSlice[i].y1
		}

		if coordSlice[i].y2 > maxY {
			maxY = coordSlice[i].y2
		}

		if coordSlice[i].x1 == coordSlice[i].x2 ||
			coordSlice[i].y1 == coordSlice[i].y2 {
			consideredLines = append(consideredLines, coordSlice[i])
		}
	}

	//fmt.Println(len(coordSlice))
	//fmt.Println(coordSlice)
	//fmt.Println(len(consideredLines))
	//fmt.Println(consideredLines)
	//fmt.Println(maxX, maxY) // max in input: 989

	// build 990x990 field
	field := make([][]int, 990)
	for i := 0; i < 990; i++ {
		field[i] = make([]int, 990)
	}

	for _, line := range coordSlice {
		if line.x1 == line.x2 {
			y1 := line.y1
			y2 := line.y2

			// swap them (line can decrease)
			if y1 > y2 {
				y1 = line.y2
				y2 = line.y1
			}

			for y := y1; y <= y2; y++ {
				field[y][line.x1] += 1
			}
		} else if line.y1 == line.y2 {
			x1 := line.x1
			x2 := line.x2

			// swpa them
			if x1 > x2 {
				x1 = line.x2
				x2 = line.x1
			}

			for x := x1; x <= x2; x++ {
				field[line.y1][x] += 1
			}
		} else { // diagonals
			diagX := int(float64(line.x2-line.x1) / math.Abs(float64(line.x2-line.x1)))
			diagY := int(float64(line.y2-line.y1) / math.Abs(float64(line.y2-line.y1)))

			x, y := line.x1, line.y1

			for y != line.y2+diagY {
				field[y][x] += 1
				y += diagY
				x += diagX
			}
		}
	}

	cnt := 0
	for _, row := range field {
		for _, v := range row {
			if v >= 2 {
				cnt += 1
			}
		}
	}

	fmt.Println(cnt)
}

func firstPart() {

	fName := "input.txt"

	data := readFile(fName)

	var maxX, maxY int
	var coordSlice = make([]Coords, len(data))
	consideredLines := []Coords{}
	for i, line := range data {
		fmt.Sscanf(line, "%d,%d -> %d,%d",
			&coordSlice[i].x1,
			&coordSlice[i].y1,
			&coordSlice[i].x2,
			&coordSlice[i].y2)

		// check for max number
		if coordSlice[i].x1 > maxX {
			maxX = coordSlice[i].x1
		}

		if coordSlice[i].x2 > maxX {
			maxX = coordSlice[i].x2
		}

		if coordSlice[i].y1 > maxY {
			maxY = coordSlice[i].y1
		}

		if coordSlice[i].y2 > maxY {
			maxY = coordSlice[i].y2
		}

		if coordSlice[i].x1 == coordSlice[i].x2 ||
			coordSlice[i].y1 == coordSlice[i].y2 {
			consideredLines = append(consideredLines, coordSlice[i])
		}
	}

	//fmt.Println(len(coordSlice))
	//fmt.Println(coordSlice)
	//fmt.Println(len(consideredLines))
	//fmt.Println(consideredLines)
	//fmt.Println(maxX, maxY) // max in input: 989

	// build 990x990 field
	field := make([][]int, 990)
	for i := 0; i < 990; i++ {
		field[i] = make([]int, 990)
	}

	for _, line := range consideredLines {
		if line.x1 == line.x2 {
			y1 := line.y1
			y2 := line.y2

			// swap them (line can decrease)
			if y1 > y2 {
				y1 = line.y2
				y2 = line.y1
			}

			for y := y1; y <= y2; y++ {
				field[y][line.x1] += 1
			}
		} else if line.y1 == line.y2 {
			x1 := line.x1
			x2 := line.x2

			// swpa them
			if x1 > x2 {
				x1 = line.x2
				x2 = line.x1
			}

			for x := x1; x <= x2; x++ {
				field[line.y1][x] += 1
			}
		} else {
			fmt.Println("This should not happen!")
		}
	}

	cnt := 0
	for _, row := range field {
		for _, v := range row {
			if v >= 2 {
				cnt += 1
			}
		}
	}

	fmt.Println(cnt)
}

func main() {
	firstPart()
	secondPart()
}
