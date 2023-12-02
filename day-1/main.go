package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

var ErrNoMatch = errors.New("no match found")

func scanFromBeginning(input string) (int, error) {
	for i := 0; i < len(input); i++ {
		for key, value := range Values {
			if strings.HasPrefix(input[i:], key) {
				return value, nil
			}
		}
	}
	return 0, ErrNoMatch
}

func scanFromEnd(input string) (int, error) {
	for i := len(input) - 1; i >= 0; i-- {
		for key, value := range Values {
			if strings.HasPrefix(input[i:], key) {
				return value, nil
			}
		}
	}
	return 0, ErrNoMatch
}

func scanSentenceReturnNumber(input string) int {
	firstDigit, err := scanFromBeginning(input)
	if err != nil {
		fmt.Println("There is an error at the first digit", err)
	}
	lastDigit, err := scanFromEnd(input)
	if err != nil {
		fmt.Println("There is an error at the last digit", err)
	}

	result, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
	if err != nil {
		fmt.Println("Error converting result to integer", err)
	}
	fmt.Println("result", result)
	return result
}

func main() {
	data, err := os.ReadFile("input.txt")

	check(err)

	myString := string(data[:])
	lines := strings.Split(myString, "\n")

	for _, line := range lines {
		scanSentenceReturnNumber(line)
	}
}
