package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type NumberPosition struct {
	startPos  int
	endPos    int
	lineIndex int
	number    int
}

// For part 1
// * and & are helpers in Golang to get the memory position(pointers) of that data
func findNumbers(input string, lineIndex int) []*NumberPosition {
	// match anything that looks like a sequence of numbers, from 20 to 3259
	expression := regexp.MustCompile(`\d+`)

	// Find all string matches the index and return [][]int convert them to integer at the end
	matches := expression.FindAllStringSubmatchIndex(input, -1)

	// Extract numbers with their start and end position and put them into the result
	//
	result := make([]*NumberPosition, len(matches))

	for i, match := range matches {
		start := match[0]
		end := match[1]
		number, _ := strconv.Atoi(input[start:end])

		result[i] = &NumberPosition{
			number:    number,
			startPos:  start,
			lineIndex: lineIndex,
			endPos:    end,
		}
	}
	return result

}

func thereIsSymbol(numb *NumberPosition, lines []string) bool {
	from := numb.startPos - 1

	if from < 0 {
		from = 0
	}

	to := numb.endPos + 1
	if to > len(lines[0]) {
		to = len(lines[0])
	} // assume all lines have same length

	// loop three lines
	for loop := numb.lineIndex - 1; loop <= numb.lineIndex+1; loop++ {
		// if it is the first line we skip
		if loop < 0 || loop >= len(lines) {
			continue
		}

		// inspect line characters
		findSymbol := strings.IndexAny(lines[loop][from:to], "+#@*=/-%&$")
		if findSymbol > -1 {
			return true
		}
	}
	return false
}

// Part 2
type SymbolPosition struct {
	startPos  int
	endPos    int
	lineIndex int
}

func findSymbolPosition(line string, star rune) []int {
	positions := []int{}

	for i, char := range line {
		if char == star {
			positions = append(positions, i)
		}
	}
	return positions
}

func numberAroundSymbol(numb *NumberPosition, lineIndex, startPos, numberOfLine int) bool {
	// loop three lines
	for loop := lineIndex - 1; loop <= lineIndex+1; loop++ {
		// if it is not within this number's index is outside of these three line we don't care and continue
		if numb.lineIndex != loop {
			continue
		}
		// if it is before the first line and beyond current star position we don't care and continue
		if loop < 0 || loop > numberOfLine {
			continue
		}

		// check if this number is around the star
		if numb.startPos == startPos || numb.startPos == startPos+1 {
			return true // start position of the number seemed fine
		}
		if numb.endPos == startPos || numb.endPos-1 == startPos {
			return true
		}

		if numb.lineIndex != lineIndex {
			if numb.startPos <= startPos && numb.endPos >= startPos {
				return true
			}
		}
	}
	return false
}
func main() {
	file, _ := os.Open("../input.txt")

	fileScan := bufio.NewScanner(file)

	fileScan.Split(bufio.ScanLines)

	var fileLines []string

	for fileScan.Scan() {
		fileLines = append(fileLines, fileScan.Text())
	}
	file.Close()

	lines := fileLines

	var allNumbers []*NumberPosition
	for lineIndex, line := range lines {

		numbersForLines := findNumbers(line, lineIndex)

		allNumbers = append(allNumbers, numbersForLines...)
	}
	sum := 0
	for _, numb := range allNumbers {
		if thereIsSymbol(numb, lines) {
			sum += numb.number
		}
	}
	numberStarSum := 0
	fmt.Println("Total:", sum)
	for lineIndexOfStar, line := range lines {
		starPosition := findSymbolPosition(line, '*')
		for _, starPos := range starPosition {
			var foundNumbers []*NumberPosition
			for _, number := range allNumbers {
				if numberAroundSymbol(number, lineIndexOfStar, starPos, len(lines)) {
					foundNumbers = append(foundNumbers, number)
				}
			}

			if len(foundNumbers) == 2 {
				numberStarSum += foundNumbers[0].number * foundNumbers[1].number
			}
		}
	}
	fmt.Printf("Number Star Sum: %d\n", numberStarSum)

}
