// Write a program to check if two strings are balanced.
// For example, strings s1 and s2 are balanced if all the characters in the s1 are present in s2.
// The character’s position doesn’t matter.
package main

import (
	"fmt"
	"strings"
)

func Solve7() {
	fmt.Println("Enter two Strings: ")
	val1, val2 := GetTwoStrings()
	fmt.Println(checkBalanced(val1, val2))
}

func checkBalanced(val1, val2 string) string {
	for _, char := range val1 {
		if !strings.Contains(val2, string(char)) {
			return "False: Strings not Balanced"
		}
	}
	return "True: Strings are Balanced"
}
