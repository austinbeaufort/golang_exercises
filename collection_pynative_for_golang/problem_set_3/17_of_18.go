// Print only words containing both letters and numbers
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Solve17() {
	sentence := "Penelope25 is a Data scientist50 and AI Expert"
	filteredWords := filterSentence(sentence)
	fmt.Println(filteredWords)
}

func filterSentence(val string) []string {
	filteredWords := []string{}
	for _, word := range getWords(val) {
		if isValid(word) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}

func isValid(word string) bool {
	hasLetter := false
	hasDigit := false

	for _, char := range word {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
	}
	return hasLetter && hasDigit
}

func getWords(val string) []string {
	return strings.Split(val, " ")
}
