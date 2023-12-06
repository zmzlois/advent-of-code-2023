package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

type Location struct {
	startNumb int
	endNumb   int
	Range     int
	lineIndex int
}

type Map struct {
	from string
	to   string
}

var seeds []int

// func findNumbers(line string, lineIndex int) []*Location {
//
// 	var numbers []*Location
// 	res := regexp.MustCompile(`\d+`)
//
// 	numberExists := res.FindAllStringSubmatchIndex(line, -1)
// 	for _, existNumb := range numberExists[0] {
// 		if lineIndex == 0 {
// 			seeds = append(seeds, existNumb)
// 			// fmt.Println("seeds:", seeds)
// 		}
// 	}
// 	return numbers
// }

func main() {
	file, fileErr := os.ReadFile("../input.txt")
	if fileErr != nil {
		fmt.Println("File parsing error")
	}
	doc := string(file[:])
	lines := strings.Split(doc, "\n")
	fmt.Println("lines:", lines)
	// for lineIndex, line := range lines {
	// 	numbers := findNumbers(line, lineIndex)
	// 	fmt.Println(numbers)
	// }

}
