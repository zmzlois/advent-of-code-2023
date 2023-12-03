package main

import (
	"fmt"
	"os"
	"strings"
)

var symbols = []string{"*", "/", "-", "&", "$", "=", "+", "@", "#", "%"}
var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func ifCharIsNumb(HashMap [][]string, i int, j int, char string, numb string) {
	var HashOne string
	var HashTwo string
	if char == numb {
		if j < len(HashMap[i]) {
			for _, numbOne := range numbers {
				if HashMap[i][j+1] == numbOne {

					HashOne = numbOne
					for _, numbTwo := range numbers {
						if HashMap[i][j+2] == numbTwo {

							HashTwo = numbTwo
						}
					}
				}
			}

		}
	}
}
func main() {
	file, _ := os.ReadFile("../input.txt")
	outputFile := string(file[:])

	lines := strings.Split(outputFile, "\n")
	HashMap := make([][]string, len(lines))
	numbStore := make([]int, 0)
	for i, line := range lines {
		// store them in a hashMap first
		HashMap[i] = strings.Split(line, "")

		chars := strings.Split(line, "")

		for j, char := range chars {
			// loop through the symbols and store the position in HashMap

			for _, numb := range numbers {
				if char == numb {

				}
			}

		}
	}
	fmt.Println("Hashmap?", HashMap)

}
