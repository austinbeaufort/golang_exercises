// Write a program to find the last position of a substring “Emma” in a given string.
package main

import (
	"fmt"
	"strings"
)

func Solve12() {
	// given string
	sentence, subString := getGivenStrings()
	reversedSentence, reversedSubString := getRevValues(sentence, subString)
	subStrStartIndex := getStartIndex(reversedSentence, reversedSubString)
	fmt.Println(getLastSubstrIndex(subStrStartIndex, len(sentence)))
}

func getLastSubstrIndex(index, length int) int {
	return length - index - 1
}

func getStartIndex(sentence, substr string) int {
	index := strings.Index(sentence, substr)
	charsToFirstIndex := len(substr) - 1
	return index + charsToFirstIndex
}

func revStr12(val string) (revVal string) {
	for i := getLastIndex(val); i >= 0; i-- {
		revVal += string(val[i])
	}
	return
}

func getLastIndex(val string) int {
	return len(val) - 1
}

func getGivenStrings() (string, string) {
	sentence := "Emma is a data scientist who knows Python. Emma works at google."
	subString := "Emma"
	return sentence, subString
}

func getRevValues(sentence, substr string) (string, string) {
	reversedSentence := revStr12(sentence)
	reversedSubString := revStr12(substr)
	return reversedSentence, reversedSubString
}
