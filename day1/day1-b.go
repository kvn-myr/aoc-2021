package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// day1-b
	// sum all elements in three-measrument sliding windows
	// add sums to new slice
	// check if element before was lower than the current element
	// if so, increase counter
	// print counter

	// read input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("You need a file called input.txt in this directory!")
	}
	defer file.Close()

	// add file content to values slice
	sc := bufio.NewScanner(file)
	var values []int

	// convert all values in slice to integer
	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		values = append(values, num)
	}

	// calculate summary of sliding windows (3 elements)
	// append to new slice (nValues)
	var nValues []int
	var curSum int
	for j := range values {

		if j <= len(values)-3 {
			curSum = values[j] + values[j+1] + values[j+2]
			nValues = append(nValues, curSum)
		}
	}

	// now iterate over the new slice instead of the inital one
	var increased int
	// skip first element
	for i := range nValues[1:] {
		if nValues[i+1] > nValues[i] {
			increased += 1
		}
	}
	fmt.Println(increased)

}
