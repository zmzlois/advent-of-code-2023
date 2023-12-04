package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbols = []string{"*", "/", "-", "&", "$", "=", "+", "@", "#", "%"}
var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

var symbolIndex = make(map[int][]int)

// for a digit array looks like [2, 3, 4] and they should be joined together as a single digit
func convertDigit(digitString []string) int {
	var digitArray []int
	result := 0
	for _, digit := range digitString {
		if len(digit) <= 1 {
			continue
		}
		digitConvert, digitErr := strconv.Atoi(digit)
		if digitErr != nil {
			fmt.Println("error converting digit")
		}
		digitArray = append(digitArray, digitConvert)
	}
	// for every number in this updated digit array, make sure they become a real, single number
	for _, digit := range digitArray {
		result = result*10 + digit
	}
	fmt.Println("convertedDigit:", result)
	return result
}

// for every position of the symbol, unless it is first line, check the line above
// here we want to get an array of strings looks like [432 573 418 318]
func checkLineAbove(HashMap [][]string, symbolIndex map[int][]int, i int, j int) []string {

	var result []string
	return result
}

// Make symbol position index
func GetSymbolIndex(HashMap [][]string, i int, j int, char string) map[int][]int {
	// var lineIndex []int

	for _, symbol := range symbols {
		if char == symbol {
			symbolIndex[i] = append(symbolIndex[i], j)
		}
	}
	return symbolIndex
}
func main() {
	file, _ := os.ReadFile("../input.txt")
	outputFile := string(file[:])

	lines := strings.Split(outputFile, "\n")
	HashMap := make([][]string, len(lines))
	// numbStore := make([]int, 0)
	for i, line := range lines {
		// store them in a hashMap first
		HashMap[i] = strings.Split(line, "")

		chars := strings.Split(line, "")

		for j, char := range chars {
			// loop through the symbols and store the position in HashMap
			symbolIndex = GetSymbolIndex(HashMap, i, j, char)

		}
	}
	// fmt.Println("symbolIndex:", symbolIndex)
	// fmt.Println("Hashmap?", HashMap)

}
