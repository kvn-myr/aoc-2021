package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// different commants like
	// forward X increases the horizontal position by X units
	// down X increases the depth by X units
	// up X decreases the depth by X units
	// depth is negative (-1, -2, -3, -4, ...)
	// puzzle input is a series of "forward X", ...
	// multiply horizontal position and depth at the end
	// -----------
	// go over all values from input
	// if "forward" increase horizontal position
	// if "down" increase depth (-1)
	// if "up" decrease depth (+1)
	// multiply horizontal and depth
	// return result

	// read input.txt
	file, err := os.Open("input.txt")
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

	var horz int
	var depth int
	ints := make([]int, len(values))

	for i, v := range values {
		split := strings.Split(v, " ")

		// split[0] == action
		// split[1] == value
		if split[0] == "forward" {
			ints[i], _ = strconv.Atoi(split[1])
			horz += ints[i]
		} else if split[0] == "down" {
			ints[i], _ = strconv.Atoi(split[1])
			depth -= ints[i]
		} else if split[0] == "up" {
			ints[i], _ = strconv.Atoi(split[1])
			depth += ints[i]
		}
	}
	fmt.Println(horz * depth)
}
