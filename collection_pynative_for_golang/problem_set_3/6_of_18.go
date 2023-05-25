// Given two strings, s1 and s2. Write a program to create a new string s3 made of the first char of s1, then the last char of s2, Next,
// the second char of s1 and second last char of s2, and so on. Any leftover chars go at the end of the result.
// if leftover chars are from value 2, the leftovers should be non-reversed.
package main

import (
	"fmt"
)

type Value struct {
	length          int
	name, remainder string
}

func Solve6() {
	fmt.Println("Enter two strings: ")
	val1, val2 := GetTwoStrings()
	fmt.Println(formatStrings(val1, val2))
}

func formatStrings(val1, val2 string) string {
	revVal2 := reverseString(val2)
	strBuilderInfo := buildStr(val1, val2)

	newStr := ""
	for i := 0; i < strBuilderInfo.length; i++ {
		newStr += string(val1[i])
		newStr += string(revVal2[i])
	}
	newStr += strBuilderInfo.remainder
	return newStr
}

func buildStr(val1, val2 string) Value {
	if len(val1) < len(val2) {
		remainder := val2[:len(val1)+1]
		return Value{len(val1), "val1", remainder}
	}
	remainder := val1[len(val2):]
	return Value{len(val2), "val2", remainder}
}

func reverseString(val string) string {
	revStr := ""
	lastIndex := len(val) - 1
	for i := lastIndex; i >= 0; i-- {
		revStr += string(val[i])
	}
	return revStr
}
