// Write a program to create a new string made of the middle three characters of an input string.
package main

import "fmt"

func Solve1() {
	fmt.Println("Please enter a word with an Odd number of chars: ")
	val := GetValidInput()
	fmt.Println(getMiddleChars(val))
}

func getMiddleChars(val string) string {
	middleIndex := len(val) / 2
	leftMidIndex := (len(val) / 2) - 1
	rightMidIndex := (len(val) / 2) + 1
	return string(val[leftMidIndex]) + string(val[middleIndex]) + string(val[rightMidIndex])
}
