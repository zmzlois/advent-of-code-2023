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

func formatAndCheckingValidSet(set string) bool {

	fmt.Println("Set looks like", set)
	// At this point the cube will be 1 red, 1 blue, 4 green
	cubes := strings.Split(set, ",")
	cubeMap := make(map[string]int, len(cubes))
	for _, cubeColor := range cubes {
		// At this point the cube will be one of 1 red, 1 blue, 4 green
		// There are white space around the character so needs to trim it
		trimmedCubeColor := strings.TrimSpace(cubeColor)
		parts := strings.Split(trimmedCubeColor, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid parts format: ", parts)
			fmt.Println("Length of parts", len(parts))

			continue
		}
		color := parts[1]
		numberOfCube, err := strconv.Atoi(parts[0])
		if err != nil {
			// fmt.Println("Failed to format numberOfCube")
			continue
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
		fmt.Printf("cubMap: %d %v", cubeMap[color], color)
		if cubeMap[color] > Requirement[color] {
			return false
		}

	}

	return true
}

func CheckSetOk(boolSet []bool) bool {
	for _, bools := range boolSet {
		if bools == false {
			return false
		}
	}
	return true
}

func main() {
	file, fileErr := os.ReadFile("../input.txt")
	if fileErr != nil {
		fmt.Println("Error reading file")
	}
	myFile := string(file[:])
	lines := strings.Split(myFile, "\n")

	// var allGameID []int
	for _, line := range lines {

		game := strings.Split(line, ":")

		if len(game) != 2 {
			fmt.Println("length of game", len(game))
			continue
		}

		// Game name: Game 89
		gameName := game[0]
		// 6 green, 5 blue; 4 green, 4 blue; 3 red, 5 blue
		gameSets := game[1]

		gameNameSplit := strings.Split(gameName, " ")

		if len(gameNameSplit) != 2 {
			fmt.Println("Game name split error", gameNameSplit)
		}
		gameId, IdErr := strconv.Atoi(gameNameSplit[1])

		if IdErr != nil {
			fmt.Println("Atoi error parsing gameId", gameId)
		}

		gameSet := strings.Split(gameSets, ";")

		for _, sets := range gameSet {
			// declare a slice to store ok
			var okSlices []bool
			// convert sets' rune value to string
			setString := string(sets)

			ok := formatAndCheckingValidSet(setString)
			okSlices = append(okSlices, ok)
			fmt.Println("ok", ok)
			fmt.Println("okSlice", okSlices)

		}

	}
	// fmt.Println("Game ID Set", allGameID)

}
