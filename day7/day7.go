package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
	//	var valuesInt []int

	// convert all values in slice to integer
	for sc.Scan() {
		//num, _ := strconv.Atoi(sc.Text())
		valuesStr = append(valuesStr, sc.Text())
	}

	return valuesStr
}

func prepData(data []string) []int {
	var dataSlice = make([]int, len(data)-1) // forgot the -1. wow that was a lot of debugging...
	splitData := strings.Split(data[0], ",")

	for _, v := range splitData {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Panic(err)
		}
		dataSlice = append(dataSlice, num)
	}
	return dataSlice
}

func getLargest(slice []int) int {
	var lNum, tmp int

	for _, el := range slice {
		if el > tmp {
			tmp = el
			lNum = tmp
		}
	}
	return lNum
}

// triangle numbers: https://www.mathsisfun.com/algebra/triangular-numbers.html
func cost(val int) int {
	return (val * (val + 1)) / 2
}

func secondPart() {
	var data = readFile("input.txt")
	var maxVal, minVal int
	dataSlice := prepData(data)
	//fmt.Println(dataSlice)

	for _, v := range dataSlice {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	var fuel []int = make([]int, maxVal)
	for i := 0; i < maxVal; i++ {
		for _, val := range dataSlice {
			fuel[i] += cost(int(math.Abs(float64(val - i))))
			//fmt.Println(dataSlice)
			//fmt.Println(val)
			//fmt.Println(fuel)
		}
	}

	//fmt.Println(maxVal, minVal)
	//fmt.Println(fuel)
	//fmt.Println(dataSlice)

	// check for smallest fuel number
	var largestFuel = getLargest(fuel)
	for _, v := range fuel {
		if v < largestFuel {
			largestFuel = v
		}
	}

	fmt.Println(largestFuel)
}

func firstPart() {
	var data = readFile("input.txt")
	var maxVal, minVal int
	dataSlice := prepData(data)
	//fmt.Println(dataSlice)

	for _, v := range dataSlice {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	var fuel []int = make([]int, maxVal)
	for i := 0; i < maxVal; i++ {
		for _, val := range dataSlice {
			fuel[i] += int(math.Abs(float64(val - i)))
			//fmt.Println(dataSlice)
			//fmt.Println(val)
			//fmt.Println(fuel)
		}
	}
	//fmt.Println(maxVal, minVal)
	//fmt.Println(fuel)
	//fmt.Println(dataSlice)

	// check for smallest fuel number
	var largestFuel = getLargest(fuel)
	for _, v := range fuel {
		if v < largestFuel {
			largestFuel = v
		}
	}

	fmt.Println(largestFuel)
}

func main() {
	firstPart()
	secondPart()
}
