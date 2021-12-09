package main

import (
	"bufio"
	"fmt"
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

	// convert all values in slice to integer
	for sc.Scan() {
		//num, _ := strconv.Atoi(sc.Text())
		valuesStr = append(valuesStr, sc.Text())
	}

	return valuesStr
}

func isOne(d string) bool {
	if len(d) == 2 {
		return true
	} else {
		return false
	}
}

func isFour(d string) bool {
	if len(d) == 4 {
		return true
	} else {
		return false
	}
}

func isSeven(d string) bool {
	if len(d) == 3 {
		return true
	} else {
		return false
	}
}

func isEight(d string) bool {
	if len(d) == 7 {
		return true
	} else {
		return false
	}
}

func secondPart() {
	// frequency analysis based on elements
	// a: 8
	// b: 6
	// c: 8
	// d: 7
	// e: 4
	// f: 9
	// g: 7
	// therefore, the numbers can determined base on their frequency:
	// 0 contains: a, b, c, e, f, g : 8+6+8+4+9+7=42
	// 1 contains: c, f				: 8+9=17
	// and so on

	// done this by hand...
	var numberMap = map[int]int{
		42: 0, // abcefg
		17: 1, // cf
		34: 2, // acdeg
		39: 3, // acdfg
		30: 4, // bcdf
		37: 5, // abdfg
		41: 6, // abdefg
		25: 7, // acf
		49: 8, // abcdefg
		45: 9, // abcdfg
	}

	var data = readFile("input.txt")
	var resSlice [][]string

	for _, row := range data {
		dataSplit := strings.Split(row, " | ")
		leftSplit := dataSplit[0]
		rightSplit := dataSplit[1]
		var resNumber []string

		// go through all character blocks on the right side
		for _, charBlock := range strings.Split(rightSplit, " ") {
			cnt := 0

			// go through all chars in the character block
			for _, char := range charBlock {
				cnt += freqCounter(string(char), leftSplit)
			}
			resNumber = append(resNumber, strconv.Itoa(numberMap[cnt]))
		}
		resSlice = append(resSlice, resNumber)
	}

	solSum := 0
	var valJoin string
	var intVal int

	// sum everything up
	for _, val := range resSlice {
		valJoin = strings.Join(val[:], "")
		intVal, _ = strconv.Atoi(valJoin)
		solSum += intVal
	}
	fmt.Println("Second part: ", solSum)
}

// count the frequency of char in data
func freqCounter(char string, data string) int {
	cnt := 0

	for _, c := range data {
		if string(c) == char {
			cnt += 1
		}
	}
	return cnt
}

func firstPart() {
	var data = readFile("input.txt")
	var cnt int

	for _, s := range data {
		fSplit := strings.Split(s, " | ")       // split at |
		sSplit := strings.Split(fSplit[1], " ") // split all values

		for _, d := range sSplit {
			if isOne(d) || isFour(d) || isSeven(d) || isEight(d) {
				cnt += 1
			}
		}
	}

	fmt.Println("First part: ", cnt)
}

func main() {
	firstPart()
	secondPart()
}
