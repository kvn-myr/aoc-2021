package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readFile(filename string) []string {
	// read input.txt
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
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
	score := 0

	for _, line := range data {

		stack := []rune{}
		corruptedLine := rune(0)

	loop:
		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')', ']', '}', '>':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if pop == '(' && c == ')' ||
					pop == '[' && c == ']' ||
					pop == '{' && c == '}' ||
					pop == '<' && c == '>' {
					continue
				} else {
					corruptedLine = c

					break loop
				}
			}
		}

		switch corruptedLine {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		default:
		}
	}
	fmt.Println("Solution for part 1: ", score)
}

func secondPart() {
	var data = readFile("input.txt")
	score := 0
	completions := []string{}

	for _, line := range data {

		stack := []rune{}
		corruptedLine := rune(0)

	loop:
		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')', ']', '}', '>':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if pop == '(' && c == ')' ||
					pop == '[' && c == ']' ||
					pop == '{' && c == '}' ||
					pop == '<' && c == '>' {
					continue
				} else {
					corruptedLine = c

					break loop
				}
			}
		}

		switch corruptedLine {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		default:
			completions = append(completions, string(stack))
		}
	}
	//fmt.Println(completions)

	scores := make([]int, len(completions))

	for i, comp := range completions {
		for j := len(comp) - 1; j >= 0; j-- {
			scores[i] *= 5
			switch comp[j] {
			case '(':
				scores[i] += 1
			case '[':
				scores[i] += 2
			case '{':
				scores[i] += 3
			case '<':
				scores[i] += 4
			}
		}
	}
	sort.Ints(scores)
	fmt.Println("Solution for part 2: ", scores[len(scores)/2])
}

func main() {
	firstPart()
	secondPart()
}
