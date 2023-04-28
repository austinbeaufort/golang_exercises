// Find all occurrences of a substring in a given string by ignoring the case.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve8() {
	fmt.Println("Enter a string, then enter a substring to search for within the string: ")
	inputReader := bufio.NewReader(os.Stdin)
	val1 := getSentence(inputReader)
	val2 := GetStrInput()
	fmt.Printf("Substring Count: %d\n", countSubstrings(val1, val2))
}

func countSubstrings(val1, val2 string) int {
	lowerVal1 := strings.ToLower(val1)
	lowerVal2 := strings.ToLower(val2)
	return strings.Count(lowerVal1, lowerVal2)
}

func getSentence(inputReader *bufio.Reader) string {
	input, err := inputReader.ReadString('\n')
	Check(err)
	return input
}
