// Given string contains a combination of the lower and upper case letters.
// Write a program to arrange the characters of a string so that all lowercase letters should come first.
package main

import (
	"fmt"
	"unicode"
)

func Solve4() {
	fmt.Println("Enter a string to format: ")
	val := GetStrInput()
	fmt.Println(formatInput(val))
}

func formatInput(val string) string {
	lower := ""
	upper := ""
	for _, char := range val {
		if unicode.IsUpper(char) {
			upper += string(char)
		} else {
			lower += string(char)
		}
	}
	return lower + upper
}
