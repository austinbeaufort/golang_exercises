// Given string contains a combination of the lower and upper case letters.
// Write a program to arrange the characters of a string so that all lowercase letters should come first.
package main

import (
	"fmt"
	"unicode"
)

func Solve5() {
	fmt.Println("Enter a string to count chars, digits, and symbols: ")
	val := GetStrInput()
	counts := getCounts(val)
	printCounts(counts)
}

func printCounts(counts map[string]int) {
	for key, value := range counts {
		fmt.Printf("%s = %d\n", key, value)
	}
}

func getCounts(val string) map[string]int {
	counts := map[string]int{
		"Chars":  0,
		"Digits": 0,
		"Symbol": 0,
	}

	for _, char := range val {
		if unicode.IsDigit(char) {
			counts["Digits"] += 1
		} else if unicode.IsLetter(char) {
			counts["Chars"] += 1
		} else {
			counts["Symbol"] += 1
		}
	}
	return counts
}
