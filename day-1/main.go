package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Slice struct {
	int
	string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var Numbers []int

//
// func isInteger(value interface{}) bool {
// 	valType := reflect.TypeOf(value)
//
// 	return valType.Kind() == reflect.Int ||
// 		valType.Kind() == reflect.Int8 ||
// 		valType.Kind() == reflect.Int16 ||
// 		valType.Kind() == reflect.Int32 ||
// 		valType.Kind() == reflect.Int64
// }

var Values = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func scanString(input string, valuesToCheck []string) []int {
	var result []int

	// Scan through each index of the sentence

	for i := 0; i < len(input); i++ {
		for _, pattern := range valuesToCheck {
			if strings.Contains(input[i:], pattern) {

				// resolve the value and update the index
				result = append(result, Values[pattern])
				i += len(pattern) - 1

				// exit the loop as soon as the result is found and start from the last index + 1
				break
			}
		}
	}
	return result
}

func orderResult(result []int, input string) []int {
	// make a map to store the index of each character in the original string
	indexes := make(map[rune]int)
	// not sure if this works but let's try
	for i, char := range input {
		indexes[char] = i
	}

	// sort the result based on original indexes
	sort.Slice(result, func(i, j int) bool {
		return indexes[rune(result[i])] < indexes[rune(result[j])]
	})
	return result
}

func scanSentenceBasedOnPattern(input string) []int {
	var result []int

	// So the loop will scan the sentence four times, each time scan through different length of the pattern to make sure we don't miss anything
	patterns := [][]string{
		{"three", "seven", "eight"},
		{"four", "five", "zero", "nine"},
		{"one", "two", "six"},
		{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}

	for _, pattern := range patterns {

		// Scan the string with the current patterns
		result = append(result, scanString(input, pattern)...)

		// Order the result based on the the original string
		result = orderResult(result, input)
	}
	// reorder them again to make sure
	result = orderResult(result, input)

	fmt.Println("result", result)

	return result
}

func main() {
	data, err := os.ReadFile("input.txt")

	check(err)

	myString := string(data[:])
	lines := strings.Split(myString, "\n")

	for _, line := range lines {
		scanSentenceBasedOnPattern(line)
	}
}
