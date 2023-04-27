// Write a program to create a new string made of an input string’s first, middle, and last character.
// If the string has an even number of chars, re-prompt the user for a word with odd number of chars.
package main

import "fmt"

func Solve0() {
	fmt.Println("Please enter a word with an Odd number of chars: ")
	val := GetValidInput()
	fmt.Println(GetChars(val))
}

func GetChars(val string) string {
	lastIndex := len(val) - 1
	middleIndex := len(val) / 2

	first := string(val[0])
	middle := string(val[middleIndex])
	last := string(val[lastIndex])
	return first + middle + last
}
