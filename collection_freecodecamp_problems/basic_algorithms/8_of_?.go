// Return the provided string with the first letter of each word capitalized.
// Make sure the rest of the word is in lower case.
package main

import (
	"fmt"
	"strings"
)

func Solve8() {
	sentence1 := "I'm a little tea pot"
	sentence2 := "sHoRt AnD sToUt"
	sentence3 := "HERE IS MY HANDLE HERE IS MY SPOUT"

	fmt.Println(makeTitleCase(sentence1))
	fmt.Println(makeTitleCase(sentence2))
	fmt.Println(makeTitleCase(sentence3))
}

func makeTitleCase(sentence string) string {
	arr := strings.Split(sentence, " ")
	titleCaseWords := []string{}
	
	for _, word := range arr {
		titleCaseWords = append(titleCaseWords, updateWord(word))
	}
	return strings.Join(titleCaseWords, " ")
}


func updateWord(word string) string {
	newWord := ""
	for i, char := range word {
		if i == 0 {
			newWord += strings.ToUpper(string(char)) 
		} else {
			newWord += strings.ToLower(string(char))
		}
	}
	return newWord
}
