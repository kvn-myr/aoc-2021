package main

import (
	"bufio"
	"fmt"
	"os"
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
	// 8 is benchmark
	// it allows us to determine the mapping between wires ans segments
	// 1. look for 8
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

	//fmt.Println(data)
	//fmt.Println(splitData)
	fmt.Println(cnt)
}

func main() {
	firstPart()
	//secondPart()
}
