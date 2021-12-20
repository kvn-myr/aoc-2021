package main

import (
	"bufio"
	"fmt"
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

}

func main() {
	firstPart()
	//	secondPart()
}
