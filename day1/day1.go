package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// day1
	// go over all integers in input
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

	var increased int
	// skip first element
	for i := range values[1:] {
		if values[i+1] > values[i] {
			increased += 1
		}
	}
	fmt.Println(increased)

}
