package main

import (
	"fmt"
	"os"
	"strings"
)

var symbols = []string{"*", "/", "-", "&", "$", "=", "+", "@", "#", "%"}
var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func ifCharIsNumb(HashMap [][]string, i int, j int, char string, numb string) {

	var numbSets = make([]string, 3)
	if char == numb {
		numbSets[0] = char
		if j < len(HashMap[i]) {
			for _, numbOne := range numbers {
				if HashMap[i][j+1] == numbOne {
					numbSets[1] = HashMap[i][j+1]
					for _, numbTwo := range numbers {
						if HashMap[i][j+2] == numbTwo {
							numbSets[2] = HashMap[i][j+2]

						}
						continue
					}
				}
				continue
			}

		}
	}
	fmt.Printf("row %d: numbSet: %v\n", i, numbSets)
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

			for _, numb := range numbers {
				ifCharIsNumb(HashMap, i, j, char, numb)

			}

		}
	}
	// fmt.Println("Hashmap?", HashMap)

}
