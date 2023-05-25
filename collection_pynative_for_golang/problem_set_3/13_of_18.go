// Write a program to split a given string on hyphens and display each substring.
package main

import (
	"fmt"
	"strings"
)

func Solve13() {
	str := "Joe-is-a-data-scientist"
	words := strings.Split(str, "-")
	printWords(words)
}

func printWords(words []string) {
	for _, word := range words {
		fmt.Println(word)
	}
}
