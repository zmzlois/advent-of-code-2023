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
		// This is correct here
		// fmt.Println("CardIndex:", cardIndex, "points:", wonPoint)
		cardPoints[cardIndex] = CardPoints{
			// index:  cardIndex,
			points: finalPoint,
			// just set one copy for now and reloop them later
			copy: 1,
		}

	}

	fmt.Println("cardPoints", cardPoints)
	// it's correct until below

	return cardPoints
}

func turnPointsIntoCopy(cardPoints map[int]CardPoints) []int {
	var points []int
	for cardIndex, card := range cardPoints {
		fmt.Println("uncalculated card index:", cardIndex, "points:", card.points, "copy:", card.copy)
		if card.points > 0 {
			// fmt.Println("card at position", cardIndex, "doesn't have a point so we skip")
			for n := 0; n <= card.points; n++ {
				cardCopy := cardPoints[cardIndex+n].copy + 1
				// cardCopy += 1
				cardPoints[cardIndex+n] = CardPoints{
					points: card.points,
					copy:   cardCopy,
				}
			}

			// if the card wins zero point, we don't do anything about the rest of the cards
		}
	}
	fmt.Println("CardPoints:", cardPoints)

	for _, card := range cardPoints {
		points = append(points, card.copy)
	}
	return points
}

func calculateCopy(copies []int) int {
	count := 0
	for _, copy := range copies {
		count = count + copy
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
