package main

import (
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

func main() {
	file, fileErr := os.ReadFile("../input.txt")
	if fileErr != nil {
		fmt.Println("Error reading file")
	}
	stringFile := string(file[:])
	fmt.Println(stringFile)

}
