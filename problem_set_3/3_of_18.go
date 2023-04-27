// Given two strings, s1 and s2, write a program to return a new string made of s1 and s2’s first, middle, and last characters.
package main

import "fmt"

func Solve3() {
	fmt.Println("Enter two strings: ")
	val1 := GetValidInput()
	val2 := GetValidInput()
	fmt.Printf("Result: %s\n", GetChars(val1)+GetChars(val2))
}
