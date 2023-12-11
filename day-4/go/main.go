package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CardPoints struct {
	index  int
	copy   int
	points int
}

func calculatePow(input int) int {
	result := 0

	for n := 1; n <= input; n++ {
		if n == 1 {
			result = 1
		} else {
			result *= 2
		}
	}
	return result
}

// task one
func findMatch(lines []string) []int {
	res := regexp.MustCompile(`\d+`)

	var points []int

	for _, line := range lines {
		afterComma := strings.Split(line, ":")
		splitString := strings.Split(afterComma[1], "|")

		if len(splitString) != 2 {
			fmt.Println("the length of this data is not correct")
		}
		winningString := splitString[0]
		ownnedString := splitString[1]

		winningStrings := res.FindAllString(winningString, -1)
		ownStrings := res.FindAllString(ownnedString, -1)

		wonPoint := 0
		for _, winning := range winningStrings {
			winNumber, _ := strconv.Atoi(winning)
			for _, owned := range ownStrings {
				ownedNumb, _ := strconv.Atoi(owned)

				if winNumber == ownedNumb {
					wonPoint = wonPoint + 1
				}

				continue
			}
			continue
		}
		PowResult := calculatePow(wonPoint)
		points = append(points, PowResult)
	}
	return points
}

func sum(input []int) int {
	result := 0
	for _, numb := range input {
		result = numb + result
	}

	return result
}

// Task two
func TaskTwoMatch(lines []string) map[int]CardPoints {
	cardPoints := make(map[int]CardPoints, len(lines))

	// var points []int
	res := regexp.MustCompile(`\d+`)

	for lineIndex, line := range lines {
		afterComma := strings.Split(line, ":")
		splitString := strings.Split(afterComma[1], "|")

		if len(splitString) != 2 {
			fmt.Println("the length of this data is not correct")
		}
		winningString := splitString[0]
		ownnedString := splitString[1]

		winningStrings := res.FindAllString(winningString, -1)
		ownStrings := res.FindAllString(ownnedString, -1)

		wonPoint := 0
		for _, winning := range winningStrings {
			winNumber, _ := strconv.Atoi(winning)
			for _, owned := range ownStrings {
				ownedNumb, _ := strconv.Atoi(owned)

				if winNumber == ownedNumb {
					wonPoint = wonPoint + 1
				}
				continue

			}
		}

		cardIndex := lineIndex + 1
		finalPoint := wonPoint
		cardPoints[cardIndex] = CardPoints{
			points: finalPoint,
			copy:   1,
		}

	}

	return cardPoints
}

func turnPointsIntoCopy(cardPoints map[int]CardPoints) map[int]int {
	copies := make(map[int]int, len(cardPoints))
	fmt.Println("Cardpoints", cardPoints)

	// for each points gained from the cards, calculate the copy of the rest of them
	for idx, card := range cardPoints {
		countEnd := idx + card.points

		if countEnd >= len(cardPoints) {
			countEnd = len(cardPoints)
		}
		for n := idx + 1; n <= countEnd; n++ {
			cardPoints[n] = CardPoints{
				copy:   cardPoints[n].copy + cardPoints[idx].copy,
				points: cardPoints[n].points,
			}
		}
		copies[idx] = cardPoints[idx].copy
	}
	// we have updated the copies from the first round, and now multiply and copies and points to apply them to the rest

	fmt.Println("Calculated cardPoints:", cardPoints)

	return copies
}

func calculateCopy(copies map[int]int) int {
	count := 0
	for idx := range copies {
		count += copies[idx]
	}

	return count
}

func main() {
	// use bufio to scan file for better precision than Os.ReadFile
	file, _ := os.Open("../input.txt")
	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	var lines []string
	for fileScan.Scan() {
		lines = append(lines, fileScan.Text())
	}
	file.Close()

	taskOneMid := findMatch(lines)
	taskOneResult := sum(taskOneMid)
	fmt.Println("TaskOneResult:", taskOneResult)

	taskTwoMid := TaskTwoMatch(lines)
	calTaskTwo := turnPointsIntoCopy(taskTwoMid)
	copyCount := calculateCopy(calTaskTwo)
	fmt.Println("Task Two Result:", copyCount)
}
