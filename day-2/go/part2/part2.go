package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumAll(numbers []int) int {
	number := 0
	for _, numb := range numbers {
		number += numb
	}
	return number
}

func powerAll(numbers []int) int {
	init := 1
	for _, number := range numbers {
		init = init * number
	}
	return init
}

var ColorMap = map[string]int{
	"red":   0,
	"blue":  0,
	"green": 0,
}

func findLargest(gameSets string) []int {
	// This will include all the gameSets here
	gameSet := strings.Split(gameSets, ";")
	// gameSet will look like 3 blue, 4 red, 7 green
	// we further down split them by comma
	for _, sets := range gameSet {

		setsWithSpaces := strings.Split(sets, ",")
		// there might be whitespace too so need to trim it
		for _, setsWithSpace := range setsWithSpaces {

			trimmedSet := strings.TrimSpace(setsWithSpace)
			// after trimmed them separate them by space
			set := strings.Split(trimmedSet, " ")
			if len(set) != 2 {
				fmt.Println("Length of set:", len(set))

				continue
			}
			color := set[1]
			numberOfColor := set[0]

		}
	}

}
func main() {
	file, fileErr := os.ReadFile("../../input.txt")
	if fileErr != nil {
		fmt.Println("Error reading file")
	}
	stringFile := string(file[:])
	fmt.Println(stringFile)
	lines := strings.Split(stringFile, "\n")

	// var total []int
	for _, line := range lines {
		game := strings.Split(line, ":")

		if len(game) != 2 {
			fmt.Println("The length of game is not 2")
			continue
		}
		// gameIndex := game[0]
		gameSets := game[1]

		fmt.Println("gamesets:", gameSets)
	}
}
