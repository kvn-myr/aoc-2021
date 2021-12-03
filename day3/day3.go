package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func firstPart() {
	var filename = "input.txt"
	var values = readFile(filename)

	var res int64
	mostCommon, leastCommon := []string{}, []string{}
	cntOne, cntZero := 0, 0

	// go over lenght of one value (12)
	for i := 0; i < len(values[0]); i++ {

		// check corresponding value (line[i]) for each value
		for _, line := range values {
			if line[i] == '1' {
				cntOne += 1
			} else {
				cntZero += 1
			}
		}

		// check if we got more 1's or 0's
		if cntOne > cntZero {
			mostCommon = append(mostCommon, "1")
			leastCommon = append(leastCommon, "0")
		} else {
			mostCommon = append(mostCommon, "0")
			leastCommon = append(leastCommon, "1")
		}
		cntOne, cntZero = 0, 0
	}
	//fmt.Println(mostCommon, leastCommon)

	// convert string slice to integer
	gamma, err := strconv.ParseInt(strings.Join(mostCommon, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	epsilon, err := strconv.ParseInt(strings.Join(leastCommon, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	res = gamma * epsilon
	fmt.Println("Solution first part: ", res)

}

func getCommon(values []string, idx int) (uint8, uint8) {
	// [0 0 1 1 1 0 1 0 1 1 0 1] [1 1 0 0 0 1 0 1 0 0 1 0]
	//mC := []string{"0", "0", "1", "1", "1", "0", "1", "0", "1", "1", "0", "1"}
	//lC := []string{"1", "1", "0", "0", "0", "1", "0", "1", "0", "0", "1", "0"}

	cnt1, cnt0 := 0, 0

	for _, line := range values {
		if line[idx] == '0' {
			cnt0 += 1
		} else {
			cnt1 += 1
		}
	}

	if cnt0 > cnt1 {
		return '0', '1'
	} else {
		return '1', '0'
	}
}

func filter(values []string, idx int, mCommon bool) string {
	// break if there is only one element left
	if len(values) == 1 {
		return values[0]
	}

	most, least := getCommon(values, idx)

	cmp := least
	if mCommon {
		cmp = most
	}

	filtered := make([]string, 0)
	for _, l := range values {
		//fmt.Println(l)
		if l[idx] != cmp {
			filtered = append(filtered, l)
		}
	}

	return filter(filtered, idx+1, mCommon)
}

func secondPart() {
	filename := "input.txt"
	scrub := filter(readFile(filename), 0, true)
	oxy := filter(readFile(filename), 0, false)

	// convert string slice to integer
	resScrub, err := strconv.ParseInt(scrub, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	resOxy, err := strconv.ParseInt(oxy, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Solution second part: ", resOxy*resScrub)
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
	secondPart()

}
