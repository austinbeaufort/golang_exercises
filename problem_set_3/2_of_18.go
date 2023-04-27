// Given two strings, s1 and s2. Write a program to create a new string s3 by appending s2 in the middle of s1.
package main

import "fmt"

func Solve2() {
	fmt.Println("Enter two strings: ")
	val1, val2 := GetTwoStrings()
	fmt.Println(appendMid(val1, val2))
}

func appendMid(val1, val2 string) string {
	midIndex := getMiddleIndex(val1)
	firstHalf := val1[:midIndex]
	lastHalf := val1[midIndex:]
	return firstHalf + val2 + lastHalf
}

func getMiddleIndex(val string) int {
	isEven := len(val)%2 == 0
	midIndex := len(val) / 2
	if isEven {
		return midIndex
	}
	return midIndex + 1
}
