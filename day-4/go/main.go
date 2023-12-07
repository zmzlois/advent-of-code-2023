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
		cardPoints[cardIndex] = CardPoints{
			index:  cardIndex,
			points: wonPoint,
			// just set one copy for now and reloop them later
			copy: 1,
		}

	}

	// it's correct until below
	for cardIndex, card := range cardPoints {
		fmt.Println("uncalculated card index:", card.index, "points:", card.points, "copy:", card.copy)
		if card.points == 0 {
			fmt.Println("card", card.index, "at position", cardIndex, "doesn't have a point so we skip")
			// if the card wins zero point, we don't do anything about the rest of the cards
			continue
		}
		for n := 1; n <= card.points; n++ {
			// cardCopy := card.copy
			// cardCopy += 1
			cardPoints[card.index+n] = CardPoints{
				copy: cardPoints[card.index+n].copy + 1,
			}
			continue
		}

		fmt.Println("calculated card index:", card.index, "points:", card.points, "copy:", card.copy)
		continue
	}

	return cardPoints
}

func calculateCopy(cardPoints map[int]CardPoints) []int {
	var allCopy []int
	for _, card := range cardPoints {
		allCopy = append(allCopy, card.copy)
	}
	fmt.Println("All copy:", allCopy)

	return allCopy
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
	calTaskTwo := calculateCopy(taskTwoMid)
	taskTwoResult := sum(calTaskTwo)
	fmt.Println("Task Two Result:", taskTwoResult)
}
