package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
var Requirement = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func formatAndCheckingValidSet(set string) int {

	// At this point the cube will be 1 red, 1 blue, 4 green
	cubes := strings.Split(set, ",")
	cubeMap := make(map[string]int, len(cubes))
	for index, cubeColor := range cubes {
		// At this point the cube will be one of 1 red, 1 blue, 4 green
		color := strings.Split(cubeColor, " ")[1]
		numberOfCube, err := strconv.Atoi(strings.Split(cubeColor, " ")[0])
		if err != nil {
			fmt.Println("Failed to format numberOfCube")
		}

		// update each cubeMap's color and number value
		if _, ok := cubeMap[color]; ok {
			cubeMap[color] = numberOfCube
		} else {
			originalNumber := cubeMap[color]
			updatedNumber := originalNumber + numberOfCube
			// plus them together and then update the map
			cubeMap[color] = updatedNumber
		}

	}

	// just put return 0 for now
	return 0
}

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	myFile := string(file[:])
	lines := strings.Split(myFile, "\n")
	for _, line := range lines {

		game := strings.Split(line, ":")[0]

		// get the game id's number here
		gameId, err := strconv.Atoi(strings.Split(game, " ")[1])

		if err != nil {
			fmt.Println("Atoi failed", err)
		}
		// the second part of each lines are the sets of games
		games := strings.Split(line, ":")[1]
		for _, sets := range games {
			// convert sets' rune value to string
			setString := string(sets)
			// split the sets by ;
			set := strings.Split(setString, ";")

		}

	}
}
