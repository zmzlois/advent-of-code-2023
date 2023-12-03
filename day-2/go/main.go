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

	// At this point the cube will be 1 red, 1 blue, 4 green
	cubes := strings.Split(set, ",")
	cubeMap := make(map[string]int, len(cubes))
	for _, cubeColor := range cubes {
		// At this point the cube will be one of 1 red, 1 blue, 4 green
		parts := strings.Split(cubeColor, " ")

		if len(parts) < 2 {
			// fmt.Println("Invalid cubColor format: ", cubeColor)

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
		if cubeMap[color] > Requirement[color] {
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

	var allGameID []int
	for _, line := range lines {

		game := strings.Split(line, ":")

		if len(game) <= 2 {
			continue 
		}

		fmt.Println("games", game)
		gameName := game[0]

		fmt.Println("Game name:", gameName)
		gameSets := game[1]
		fmt.Println("Game Sets: ", gameSets)

		gameNameSplit := strings.Split(gameName, " ")

		if len(gameNameSplit) < 2 {
			fmt.Println("Game name split error", gameNameSplit)
		}
		gameId, IdErr := strconv.Atoi(gameNameSplit[1])

		if IdErr != nil {
			fmt.Println("Atoi error parsing gameId", gameId)
		}

		for _, sets := range gameSets {
			// convert sets' rune value to string
			setString := string(sets)
			// split the sets by ;
			set := strings.Split(setString, ";")
			for _, smallset := range set {

				if ok := formatAndCheckingValidSet(smallset); ok {
					allGameID = append(allGameID, gameId)
				}
			}

		}

	} // fmt.Println("Game ID Set", allGameID)

}
