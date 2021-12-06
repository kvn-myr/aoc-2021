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
	var values []string

	// convert all values in slice to integer
	for sc.Scan() {
		//num, _ := strconv.Atoi(sc.Text())
		values = append(values, sc.Text())
	}

	return values
}

func updateFish(fish []int) []int {
	var newFishes []int

	for i := range fish {
		if fish[i] == 0 { // fish dies and adds a new fish
			newFishes = append(newFishes, 8)
			fish[i] = 6
		} else { // decrease fish life
			fish[i] -= 1
		}
	}
	return append(fish, newFishes...) // varadic parameter
}

func updateFishOnDay(fish []int) []int {
	var fishOnDays = make([]int, 9)

	for i := 1; i < 9; i++ {
		fishOnDays[i-1] = fish[i]
	}
	fishOnDays[6] += fish[0]
	fishOnDays[8] += fish[0]
	return fishOnDays
}

func secondPart() {
	data := readFile("input.txt")

	var fishSlice = make([]int, 9)
	splitData := strings.Split(data[0], ",")

	// go through all fishes and append them to a slice as an integer
	for _, fish := range splitData {
		number, err := strconv.Atoi(fish)
		if err != nil {
			fmt.Println("error!!", err)
		}
		fishSlice[number] += 1
	}

	days := 256
	for day := 0; day < days; day++ { // < 80 because we starte at 0 ...
		fishSlice = updateFishOnDay(fishSlice)
	}

	//fmt.Println(fishSlice)

	var sum int64
	for _, v := range fishSlice {
		sum += int64(v)
	}
	fmt.Println(sum)

	//fmt.Println(data)
	//fmt.Println(fishSlice)
}

func firstPart() {
	data := readFile("input.txt")

	var fishSlice []int
	splitData := strings.Split(data[0], ",")

	// go through all fishes and append them to a slice as an integer
	for _, fish := range splitData {
		number, err := strconv.Atoi(fish)
		if err != nil {
			fmt.Println("error!!", err)
		}
		fishSlice = append(fishSlice, number)
	}

	days := 80
	for day := 0; day < days; day++ { // < 80 because we starte at 0 ...
		fishSlice = updateFish(fishSlice)
	}

	fmt.Println(len(fishSlice))

	//fmt.Println(data)
	//fmt.Println(fishSlice)
}

func main() {
	firstPart()
	secondPart()
}
